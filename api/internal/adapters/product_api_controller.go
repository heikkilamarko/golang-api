package adapters

import (
	"encoding/json"
	"net/http"
	"product-api/internal/app"
	"product-api/internal/app/command"
	"product-api/internal/app/query"
	"product-api/internal/domain"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

const (
	errCodeInvalidID          = "invalid_id"
	errCodeInvalidOffset      = "invalid_offset"
	errCodeInvalidLimit       = "invalid_limit"
	errCodeInvalidRequestBody = "invalid_request_body"
)

const (
	fieldID               = "id"
	fieldPaginationOffset = "offset"
	fieldPaginationLimit  = "limit"
	fieldRequestBody      = "request_body"
)

const (
	limitMaxPageSize = 100
)

type ProductAPIController struct {
	app    *app.App
	logger *zerolog.Logger
}

func NewProductAPIController(app *app.App, logger *zerolog.Logger) *ProductAPIController {
	return &ProductAPIController{app, logger}
}

func (c *ProductAPIController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", c.getProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", c.createProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", c.getProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/{id:[0-9]+}", c.updateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{id:[0-9]+}", c.deleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/products/pricerange", c.getPriceRange).Methods(http.MethodGet)
}

// Handlers

func (c *ProductAPIController) getProducts(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductsQuery(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := c.app.Queries.GetProducts.Handle(r.Context(), query)

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func (c *ProductAPIController) createProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateProductCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.CreateProduct.Handle(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, command.Product, nil)
}

func (c *ProductAPIController) getProduct(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductQuery(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	product, err := c.app.Queries.GetProduct.Handle(r.Context(), query)

	if err != nil {
		c.logError(err)
		switch err {
		case domain.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, product, nil)
}

func (c *ProductAPIController) updateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseUpdateProductCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.UpdateProduct.Handle(r.Context(), command); err != nil {
		c.logError(err)
		switch err {
		case domain.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, command.Product, nil)
}

func (c *ProductAPIController) deleteProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseDeleteProductCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.DeleteProduct.Handle(r.Context(), command); err != nil {
		c.logError(err)
		switch err {
		case domain.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteNoContent(w)
}

func (c *ProductAPIController) getPriceRange(w http.ResponseWriter, r *http.Request) {
	pr, err := c.app.Queries.GetPriceRange.Handle(r.Context())

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, pr, nil)
}

// Input parsers

func parseGetProductsQuery(r *http.Request) (*query.GetProducts, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := limitMaxPageSize

	var err error

	if value := r.FormValue(fieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[fieldPaginationOffset] = errCodeInvalidOffset
		}
	}

	if value := r.FormValue(fieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || limitMaxPageSize < limit {
			errorMap[fieldPaginationLimit] = errCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &query.GetProducts{
		Offset: offset,
		Limit:  limit,
	}, nil
}

func parseCreateProductCommand(r *http.Request) (*command.CreateProduct, error) {
	errorMap := map[string]string{}

	product := &domain.Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorMap[fieldRequestBody] = errCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.CreateProduct{Product: product}, nil
}

func parseGetProductQuery(r *http.Request) (*query.GetProduct, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[fieldID])
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &query.GetProduct{ID: id}, nil
}

func parseUpdateProductCommand(r *http.Request) (*command.UpdateProduct, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[fieldID])
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	product := &domain.Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorMap[fieldRequestBody] = errCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	if id != product.ID {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.UpdateProduct{Product: product}, nil
}

func parseDeleteProductCommand(r *http.Request) (*command.DeleteProduct, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[fieldID])
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.DeleteProduct{ID: id}, nil
}

// Utils

func (c *ProductAPIController) logError(err error) {
	c.logger.Error().Err(err).Send()
}
