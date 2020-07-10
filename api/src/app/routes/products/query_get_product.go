package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	p := newGetProductRequestParser(r)
	p.parse()

	if !p.IsValid() {
		utils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	product, err := c.Repository.GetProduct(r.Context(), p.query)

	if err != nil {
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteOK(w, product, nil)
}

func newGetProductRequestParser(r *http.Request) *getProductRequestParser {
	return &getProductRequestParser{utils.RequestValidator{Request: r}, nil}
}

type getProductRequestParser struct {
	utils.RequestValidator
	query *GetProductQuery
}

func (v *getProductRequestParser) parse() {
	validationErrors := map[string]string{}

	id, err := utils.GetRequestVarInt(v.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.query = &GetProductQuery{id}
	}
}
