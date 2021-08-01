package app

import (
	"product-api/internal/app/command"
	"product-api/internal/app/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProduct *command.CreateProductHandler
	UpdateProduct *command.UpdateProductHandler
	DeleteProduct *command.DeleteProductHandler
}

type Queries struct {
	GetProducts   *query.GetProductsHandler
	GetProduct    *query.GetProductHandler
	GetPriceRange *query.GetPriceRangeHandler
}
