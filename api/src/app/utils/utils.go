package utils

import (
	"errors"
	"net/http"

	"github.com/heikkilamarko/goutils"
)

// HandleParseRequestError func
func HandleParseRequestError(err error, w http.ResponseWriter) {
	var verr *goutils.ValidationError
	if errors.As(err, &verr) {
		goutils.WriteBadRequest(w, verr.ValidationErrors)
	} else {
		goutils.WriteInternalError(w, nil)
	}
}
