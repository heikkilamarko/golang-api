package application

import (
	"product-api/internal/application/command"
	"product-api/internal/application/query"
)

type Application struct {
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
