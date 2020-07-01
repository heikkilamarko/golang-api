package middleware

import (
	"net/http"
	"products-api/app/config"
	"products-api/app/constants"
	"products-api/app/utils"
)

// Auth middleware.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expected := config.Config.APIKey
		actual := r.Header.Get(constants.HeaderAPIKey)

		if actual != expected {
			utils.WriteUnauthorized(w, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
