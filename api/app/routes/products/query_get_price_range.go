package products

import (
	"net/http"

	"github.com/heikkilamarko/goutils"
)

// GetPriceRange query
func (c *Controller) GetPriceRange(w http.ResponseWriter, r *http.Request) {
	pr, err := c.repository.getPriceRange(r.Context())

	if err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, pr, nil)
}
