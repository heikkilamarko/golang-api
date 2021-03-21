package products

import (
	"context"
	"database/sql"
	_ "embed"

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

//go:embed sql/get_products.sql
var qetProductsSQL string

// GetProducts method
func (r *SQLRepository) GetProducts(ctx context.Context, query *GetProductsQuery) ([]*Product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetProductsSQL,
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

//go:embed sql/get_product.sql
var qetProductSQL string

// GetProduct method
func (r *SQLRepository) GetProduct(ctx context.Context, query *GetProductQuery) (*Product, error) {
	p := &Product{}

	err := r.db.QueryRowContext(
		ctx,
		qetProductSQL,
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

//go:embed sql/create_product.sql
var createProductSQL string

// CreateProduct method
func (r *SQLRepository) CreateProduct(ctx context.Context, command *CreateProductCommand) error {
	p := command.Product

	err := r.db.QueryRowContext(
		ctx,
		createProductSQL,
		p.Name, p.Description, p.Price, p.Comment).Scan(&p.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}

//go:embed sql/update_product.sql
var updateProductSQL string

// UpdateProduct method
func (r *SQLRepository) UpdateProduct(ctx context.Context, command *UpdateProductCommand) error {
	p := command.Product

	result, err := r.db.ExecContext(
		ctx,
		updateProductSQL,
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

//go:embed sql/delete_product.sql
var deleteProductSQL string

// DeleteProduct method
func (r *SQLRepository) DeleteProduct(ctx context.Context, command *DeleteProductCommand) error {
	result, err := r.db.ExecContext(
		ctx,
		deleteProductSQL,
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

//go:embed sql/get_price_range.sql
var getPriceRangeSQL string

// GetPriceRange method
func (r *SQLRepository) GetPriceRange(ctx context.Context) (*PriceRange, error) {
	pr := &PriceRange{}

	err := r.db.QueryRowContext(
		ctx,
		getPriceRangeSQL).Scan(&pr.MinPrice, &pr.MaxPrice)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	return pr, nil
}
