package products

import (
	"context"
	"database/sql"
	_ "embed"
	"products-api/app/utils"
	"time"
)

type repository struct {
	db *sql.DB
}

func (r *repository) getProducts(ctx context.Context, query *getProductsQuery) ([]*product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetProductsSQL,
		query.Limit, query.Offset)

	if err != nil {
		return nil, err
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
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

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
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, err
		}
	}

	return p, nil
}

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
		return err
	}

	return nil
}

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
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count < 1 {
		return utils.ErrNotFound
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
		return err
	}

	return nil
}

func (r *repository) deleteProduct(ctx context.Context, command *deleteProductCommand) error {
	result, err := r.db.ExecContext(ctx, deleteProductSQL, command.ID)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (r *repository) getPriceRange(ctx context.Context) (*priceRange, error) {
	pr := &priceRange{}

	err := r.db.QueryRowContext(ctx, getPriceRangeSQL).Scan(
		&pr.MinPrice,
		&pr.MaxPrice,
	)

	if err != nil {
		return nil, err
	}

	return pr, nil
}
