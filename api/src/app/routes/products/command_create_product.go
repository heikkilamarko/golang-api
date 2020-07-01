package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	v := createProductCommandValidator{utils.RequestValidator{Request: r}, nil}
	v.parseAndValidate()

	if !v.IsValid() {
		utils.WriteBadRequest(w, v.ValidationErrors)
		return
	}

	if err := c.Repository.createProduct(r.Context(), v.command.product); err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteCreated(w, v.command.product, nil)
}

type createProductCommandValidator struct {
	utils.RequestValidator
	command *createProductCommand
}

func (v *createProductCommandValidator) parseAndValidate() {
	validationErrors := map[string]string{}

	product := &product{}
	if err := json.NewDecoder(v.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.command = &createProductCommand{product}
	}
}
