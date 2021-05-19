package products

import (
	"net/http"
	"products-api/app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductRequest(r)

	if err != nil {
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	product, err := c.repository.getProduct(r.Context(), query)

	if err != nil {
		c.logError(err)
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

func parseGetProductRequest(r *http.Request) (*getProductQuery, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, utils.NewValidationError(errorMap)
	}

	return &getProductQuery{id}, nil
}
