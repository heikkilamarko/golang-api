package utils

import (
	"encoding/json"
	"net/http"
)

// WriteOK writes 200 response
func WriteOK(w http.ResponseWriter, data, meta interface{}) {
	writeResponse(w, http.StatusOK, newDataResponse(data, meta))
}

// WriteCreated writes 201 response
func WriteCreated(w http.ResponseWriter, data, meta interface{}) {
	writeResponse(w, http.StatusCreated, newDataResponse(data, meta))
}

// WriteNoContent writes 204 response
func WriteNoContent(w http.ResponseWriter) {
	writeResponse(w, http.StatusNoContent, nil)
}

// WriteBadRequest writes 400 response
func WriteBadRequest(w http.ResponseWriter, details map[string]string) {
	writeResponse(w, http.StatusBadRequest, newBadRequestResponse(details))
}

// WriteUnauthorized writes 401 response
func WriteUnauthorized(w http.ResponseWriter, details map[string]string) {
	writeResponse(w, http.StatusUnauthorized, newUnauthorizedResponse(details))
}

// WriteNotFound writes 404 response
func WriteNotFound(w http.ResponseWriter, details map[string]string) {
	writeResponse(w, http.StatusNotFound, newNotFoundResponse(details))
}

// WriteInternalError writes 500 response
func WriteInternalError(w http.ResponseWriter, details map[string]string) {
	writeResponse(w, http.StatusInternalServerError, newInternalErrorResponse(details))
}

func writeResponse(w http.ResponseWriter, code int, body interface{}) {
	if body != nil {
		content, err := json.Marshal(body)

		if err != nil {
			WriteInternalError(w, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(content)
	} else {
		w.WriteHeader(code)
	}
}
