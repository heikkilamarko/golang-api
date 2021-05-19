package products

import (
	"net/http"
	"products-api/app/utils"
)

// GetPriceRange query
func (c *Controller) GetPriceRange(w http.ResponseWriter, r *http.Request) {
	pr, err := c.repository.getPriceRange(r.Context())

	if err != nil {
		c.logError(err)
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, pr, nil)
}
