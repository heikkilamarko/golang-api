package middleware

import (
	"net/http"
	"products-api/app/utils"

	"github.com/rs/zerolog/log"
)

// Recovery middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("%s", err)
				utils.WriteInternalError(w, nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
