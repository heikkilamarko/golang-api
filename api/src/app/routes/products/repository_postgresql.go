// +build data_postgresql

package products

import (
	"context"
	"database/sql"
	"fmt"
	"products-api/app/config"
	"products-api/app/constants"
	"products-api/app/utils"

	"github.com/rs/zerolog/log"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func (repo *repository) initialize() {
	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=require",
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBName,
		config.Config.DBUsername,
		config.Config.DBPassword)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	repo.db = db
}

func (repo *repository) GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DBQueryTimeout)
	defer cancel()

	rows, err := repo.db.QueryContext(
		ctx,
		`
		SELECT id, name, description, price, comment
		FROM products.products
		ORDER BY id
		LIMIT $1 OFFSET $2
		`,
		query.Limit, query.Offset)

	if err != nil {
		log.Err(err).Send()
		return nil, utils.ErrInternalError
	}

	defer rows.Close()

	products := []*Product{}

	for rows.Next() {
		p := &Product{}

		var description sql.NullString
		var comment sql.NullString

		if err := rows.Scan(&p.ID, &p.Name, &description, &p.Price, &comment); err != nil {
			log.Err(err).Send()
			return nil, utils.ErrInternalError
		}

		p.Description = utils.GetNullStringValue(description)
		p.Comment = utils.GetNullStringValue(comment)

		products = append(products, p)
	}

	return products, nil
}

func (repo *repository) GetProduct(ctx context.Context, query *GetProductQuery) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DBQueryTimeout)
	defer cancel()

	var description sql.NullString
	var comment sql.NullString

	p := &Product{}

	err := repo.db.QueryRowContext(
		ctx,
		`
		SELECT id, name, description, price, comment
		FROM products.products
		WHERE id=$1
		`,
		query.ID).Scan(&p.ID, &p.Name, &description, &p.Price, &comment)

	if err != nil {
		log.Err(err).Send()
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, utils.ErrInternalError
		}
	}

	p.Description = utils.GetNullStringValue(description)
	p.Comment = utils.GetNullStringValue(comment)

	return p, nil
}

func (repo *repository) CreateProduct(ctx context.Context, command *CreateProductCommand) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DBQueryTimeout)
	defer cancel()

	p := command.Product

	err := repo.db.QueryRowContext(
		ctx,
		`
		INSERT INTO products.products(name, description, price, comment)
		VALUES($1, $2, $3, $4)
		RETURNING id
		`,
		p.Name, p.Description, p.Price, p.Comment).Scan(&p.ID)

	if err != nil {
		log.Err(err).Send()
		return utils.ErrInternalError
	}

	return nil
}

func (repo *repository) UpdateProduct(ctx context.Context, command *UpdateProductCommand) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DBQueryTimeout)
	defer cancel()

	p := command.Product

	result, err := repo.db.ExecContext(
		ctx,
		`
		UPDATE products.products
		SET name=$1, description=$2, price=$3, comment=$4
		WHERE id=$5
		`,
		p.Name, p.Description, p.Price, p.Comment, p.ID)

	if err != nil {
		log.Err(err).Send()
		return utils.ErrInternalError
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Err(err).Send()
		return utils.ErrInternalError
	}

	if count < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (repo *repository) DeleteProduct(ctx context.Context, command *DeleteProductCommand) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DBQueryTimeout)
	defer cancel()

	result, err := repo.db.ExecContext(
		ctx,
		`
		DELETE FROM products.products
		WHERE id=$1
		`,
		command.ID)

	if err != nil {
		log.Err(err).Send()
		return utils.ErrInternalError
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Err(err).Send()
		return utils.ErrInternalError
	}

	if count < 1 {
		return utils.ErrNotFound
	}

	return nil
}
