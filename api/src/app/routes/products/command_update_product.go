package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// UpdateProduct command
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p := newUpdateProductRequestParser(r).parse()

	if !p.IsValid() {
		utils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.UpdateProduct(r.Context(), p.command); err != nil {
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteOK(w, p.command.Product, nil)
}

func newUpdateProductRequestParser(r *http.Request) *updateProductRequestParser {
	return &updateProductRequestParser{utils.RequestValidator{Request: r}, nil}
}

type updateProductRequestParser struct {
	utils.RequestValidator
	command *UpdateProductCommand
}

func (p *updateProductRequestParser) parse() *updateProductRequestParser {
	validationErrors := map[string]string{}

	id, err := utils.GetRequestVarInt(p.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	product := &Product{}
	if err := json.NewDecoder(p.Request.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
		return p
	}

	if id != product.ID {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		p.ValidationErrors = validationErrors
	} else {
		p.command = &UpdateProductCommand{product}
	}

	return p
}
