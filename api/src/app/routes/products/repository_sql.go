package products

import (
	"context"
	"database/sql"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

// SQLRepository struct
type SQLRepository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

// NewSQLRepository func
func NewSQLRepository(db *sql.DB, l *zerolog.Logger) *SQLRepository {
	return &SQLRepository{db, l}
}

// GetProducts method
func (r *SQLRepository) GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`
		SELECT id, name, description, price, comment
		FROM products.products
		ORDER BY id
		LIMIT $1 OFFSET $2
		`,
		query.Limit, query.Offset)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	defer rows.Close()

	products := []*Product{}

	for rows.Next() {
		p := &Product{}

		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Comment); err != nil {
			r.logger.Err(err).Send()
			return nil, goutils.ErrInternalError
		}

		products = append(products, p)
	}

	return products, nil
}

// GetProduct method
func (r *SQLRepository) GetProduct(ctx context.Context, query *GetProductQuery) (*Product, error) {
	p := &Product{}

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT id, name, description, price, comment
		FROM products.products
		WHERE id=$1
		`,
		query.ID).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Comment)

	if err != nil {
		r.logger.Err(err).Send()
		switch err {
		case sql.ErrNoRows:
			return nil, goutils.ErrNotFound
		default:
			return nil, goutils.ErrInternalError
		}
	}

	return p, nil
}

// CreateProduct method
func (r *SQLRepository) CreateProduct(ctx context.Context, command *CreateProductCommand) error {
	p := command.Product

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO products.products(name, description, price, comment)
		VALUES($1, $2, $3, $4)
		RETURNING id
		`,
		p.Name, p.Description, p.Price, p.Comment).Scan(&p.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}

// UpdateProduct method
func (r *SQLRepository) UpdateProduct(ctx context.Context, command *UpdateProductCommand) error {
	p := command.Product

	result, err := r.db.ExecContext(
		ctx,
		`
		UPDATE products.products
		SET name=$1, description=$2, price=$3, comment=$4
		WHERE id=$5
		`,
		p.Name, p.Description, p.Price, p.Comment, p.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	count, err := result.RowsAffected()

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	if count < 1 {
		return goutils.ErrNotFound
	}

	return nil
}

// DeleteProduct method
func (r *SQLRepository) DeleteProduct(ctx context.Context, command *DeleteProductCommand) error {
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM products.products
		WHERE id=$1
		`,
		command.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	count, err := result.RowsAffected()

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	if count < 1 {
		return goutils.ErrNotFound
	}

	return nil
}

// GetPriceRange method
func (r *SQLRepository) GetPriceRange(ctx context.Context) (*PriceRange, error) {
	pr := &PriceRange{}

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT min_price, max_price
		FROM products.price_range()
		`).Scan(&pr.MinPrice, &pr.MaxPrice)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	return pr, nil
}
