package adapters

import (
	"encoding/json"
	"net/http"
	"product-api/internal/application"
	"product-api/internal/application/command"
	"product-api/internal/application/query"
	"product-api/internal/domain"
	"product-api/internal/ports"
	"strconv"

	"github.com/go-chi/chi/v5"
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

type ProductHTTPHandlers struct {
	app    *application.Application
	logger *zerolog.Logger
}

func NewProductHTTPHandlers(app *application.Application, logger *zerolog.Logger) *ProductHTTPHandlers {
	return &ProductHTTPHandlers{app, logger}
}

// Handlers

func (h *ProductHTTPHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	q, err := parseGetProductsQuery(r)

	if err != nil {
		h.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := h.app.Queries.GetProducts.Handle(r.Context(), q)

	if err != nil {
		h.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	meta := &paginationMeta{
		Offset: q.Offset,
		Limit:  q.Limit,
	}

	goutils.WriteOK(w, todos, meta)
}

func (h *ProductHTTPHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	c, err := parseCreateProductCommand(r)

	if err != nil {
		h.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := h.app.Commands.CreateProduct.Handle(r.Context(), c); err != nil {
		h.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, c.Product, nil)
}

func (h *ProductHTTPHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	q, err := parseGetProductQuery(r)

	if err != nil {
		h.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	product, err := h.app.Queries.GetProduct.Handle(r.Context(), q)

	if err != nil {
		h.logError(err)
		switch err {
		case ports.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, product, nil)
}

func (h *ProductHTTPHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	c, err := parseUpdateProductCommand(r)

	if err != nil {
		h.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := h.app.Commands.UpdateProduct.Handle(r.Context(), c); err != nil {
		h.logError(err)
		switch err {
		case ports.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, c.Product, nil)
}

func (h *ProductHTTPHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	c, err := parseDeleteProductCommand(r)

	if err != nil {
		h.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := h.app.Commands.DeleteProduct.Handle(r.Context(), c); err != nil {
		h.logError(err)
		switch err {
		case ports.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteNoContent(w)
}

func (h *ProductHTTPHandlers) GetPriceRange(w http.ResponseWriter, r *http.Request) {
	pr, err := h.app.Queries.GetPriceRange.Handle(r.Context())

	if err != nil {
		h.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, pr, nil)
}

// Utils

func (h *ProductHTTPHandlers) logError(err error) {
	h.logger.Error().Err(err).Send()
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

	id, err := strconv.Atoi(chi.URLParam(r, fieldID))
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

	id, err := strconv.Atoi(chi.URLParam(r, fieldID))
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

	id, err := strconv.Atoi(chi.URLParam(r, fieldID))
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.DeleteProduct{ID: id}, nil
}
