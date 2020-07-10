package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	v := newDeleteProductCommandParser(r)
	v.parse()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	if err := c.Repository.DeleteProduct(r.Context(), v.command); err != nil {
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteNoContent(w)
}

func newDeleteProductCommandParser(r *http.Request) *deleteProductCommandParser {
	return &deleteProductCommandParser{utils.RequestValidator{Request: r}, nil}
}

type deleteProductCommandParser struct {
	utils.RequestValidator
	command *DeleteProductCommand
}

func (v *deleteProductCommandParser) parse() {
	validationErrors := map[string]string{}

	id, err := utils.GetRequestVarInt(v.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.command = &DeleteProductCommand{id}
	}
}
