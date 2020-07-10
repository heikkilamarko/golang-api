package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	v := newGetProductQueryParser(r)
	v.parse()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	product, err := c.Repository.GetProduct(r.Context(), v.query)

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

func newGetProductQueryParser(r *http.Request) *getProductQueryParser {
	return &getProductQueryParser{utils.RequestValidator{Request: r}, nil}
}

type getProductQueryParser struct {
	utils.RequestValidator
	query *GetProductQuery
}

func (v *getProductQueryParser) parse() {
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
