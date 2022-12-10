package adapters

import (
	"net/http"

	"github.com/heikkilamarko/goutils"
)

func APIKey(apiKey, headerKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ak := r.Header.Get(headerKey); ak != apiKey {
				goutils.WriteUnauthorized(w, nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
