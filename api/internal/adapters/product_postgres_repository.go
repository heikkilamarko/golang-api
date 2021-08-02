package adapters

import (
	"context"
	"database/sql"
	_ "embed"
	"product-api/internal/domain"
	"product-api/internal/ports"
)

var (
	//go:embed sql/get_products.sql
	qetProductsSQL string
	//go:embed sql/get_product.sql
	qetProductSQL string
	//go:embed sql/create_product.sql
	createProductSQL string
	//go:embed sql/update_product.sql
	updateProductSQL string
	//go:embed sql/delete_product.sql
	deleteProductSQL string
	//go:embed sql/get_price_range.sql
	getPriceRangeSQL string
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{db}
}

func (r *ProductPostgresRepository) GetProducts(ctx context.Context, query *ports.GetProductsQuery) ([]*domain.Product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetProductsSQL,
		query.Limit, query.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		p := &domain.Product{}

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

func (r *ProductPostgresRepository) GetProduct(ctx context.Context, id int) (*domain.Product, error) {
	p := &domain.Product{}

	err := r.db.QueryRowContext(ctx, qetProductSQL, id).Scan(
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
			return nil, ports.ErrNotFound
		default:
			return nil, err
		}
	}

	return p, nil
}

func (r *ProductPostgresRepository) CreateProduct(ctx context.Context, p *domain.Product) error {
	return r.db.QueryRowContext(ctx, createProductSQL,
		p.Name,
		p.Description,
		p.Price,
		p.Comment,
		p.CreatedAt).
		Scan(&p.ID)
}

func (r *ProductPostgresRepository) UpdateProduct(ctx context.Context, p *domain.Product) error {
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
		return ports.ErrNotFound
	}

	return r.db.QueryRowContext(ctx, qetProductSQL, p.ID).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Comment,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}

func (r *ProductPostgresRepository) DeleteProduct(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, deleteProductSQL, id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count < 1 {
		return ports.ErrNotFound
	}

	return nil
}

func (r *ProductPostgresRepository) GetPriceRange(ctx context.Context) (*domain.PriceRange, error) {
	pr := &domain.PriceRange{}

	err := r.db.QueryRowContext(ctx, getPriceRangeSQL).Scan(
		&pr.MinPrice,
		&pr.MaxPrice,
	)

	if err != nil {
		return nil, err
	}

	return pr, nil
}
