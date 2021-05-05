package products

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

type repository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func newRepository(db *sql.DB, l *zerolog.Logger) *repository {
	return &repository{db, l}
}

//go:embed sql/get_products.sql
var qetProductsSQL string

func (r *repository) getProducts(ctx context.Context, query *getProductsQuery) ([]*product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetProductsSQL,
		query.Limit, query.Offset)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	defer rows.Close()

	products := []*product{}

	for rows.Next() {
		p := &product{}

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Comment,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			r.logger.Err(err).Send()
			return nil, goutils.ErrInternalError
		}

		products = append(products, p)
	}

	return products, nil
}

//go:embed sql/get_product.sql
var qetProductSQL string

func (r *repository) getProduct(ctx context.Context, query *getProductQuery) (*product, error) {
	p := &product{}

	err := r.db.QueryRowContext(ctx, qetProductSQL, query.ID).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Comment,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

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

func (r *repository) createProduct(ctx context.Context, command *createProductCommand) error {
	p := command.Product

	p.CreatedAt = time.Now()

	err := r.db.QueryRowContext(ctx, createProductSQL,
		p.Name,
		p.Description,
		p.Price,
		p.Comment,
		p.CreatedAt).
		Scan(&p.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}

//go:embed sql/update_product.sql
var updateProductSQL string

func (r *repository) updateProduct(ctx context.Context, command *updateProductCommand) error {
	p := command.Product

	now := time.Now()

	p.UpdatedAt = &now

	result, err := r.db.ExecContext(ctx, updateProductSQL,
		p.Name,
		p.Description,
		p.Price,
		p.Comment,
		p.UpdatedAt,
		p.ID,
	)

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

	err = r.db.QueryRowContext(ctx, qetProductSQL, p.ID).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Comment,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}

//go:embed sql/delete_product.sql
var deleteProductSQL string

func (r *repository) deleteProduct(ctx context.Context, command *deleteProductCommand) error {
	result, err := r.db.ExecContext(ctx, deleteProductSQL, command.ID)

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

func (r *repository) getPriceRange(ctx context.Context) (*priceRange, error) {
	pr := &priceRange{}

	err := r.db.QueryRowContext(ctx, getPriceRangeSQL).Scan(
		&pr.MinPrice,
		&pr.MaxPrice,
	)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	return pr, nil
}
