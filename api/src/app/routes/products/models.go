package products

type product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Comment     string  `json:"comment"`
}

// Queries

type getProductsQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type getProductQuery struct {
	id int
}

// Commands

type createProductCommand struct {
	product *product
}

type updateProductCommand struct {
	product *product
}

type deleteProductCommand struct {
	id int
}
