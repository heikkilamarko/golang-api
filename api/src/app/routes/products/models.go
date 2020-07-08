package products

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Comment     string  `json:"comment"`
}

// GetProductsQuery struct
type GetProductsQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetProductQuery struct
type GetProductQuery struct {
	ID int
}

// CreateProductCommand struct
type CreateProductCommand struct {
	Product *Product
}

// UpdateProductCommand struct
type UpdateProductCommand struct {
	Product *Product
}

// DeleteProductCommand struct
type DeleteProductCommand struct {
	ID int
}
