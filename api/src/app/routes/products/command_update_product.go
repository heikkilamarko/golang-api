package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// UpdateProduct command
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	v := updateProductCommandValidator{utils.RequestValidator{Request: r}, nil}
	v.parseAndValidate()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	if err := c.Repository.updateProduct(r.Context(), v.command.product); err != nil {
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteOK(w, v.command.product, nil)
}

type updateProductCommandValidator struct {
	utils.RequestValidator
	command *updateProductCommand
}

func (v *updateProductCommandValidator) parseAndValidate() {
	validationErrors := map[string]string{}

	id, err := utils.GetRequestVarInt(v.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	product := &product{}
	if err := json.NewDecoder(v.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
		return
	}

	if id != product.ID {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.command = &updateProductCommand{product}
	}
}
