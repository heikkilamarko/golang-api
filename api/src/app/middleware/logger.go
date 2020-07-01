package middleware

import (
	"fmt"
	"net/http"
	"products-api/app/utils"
	"time"

	"github.com/rs/zerolog/log"
)

// Logger middleware.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		lw := utils.NewLoggingResponseWriter(w)
		next.ServeHTTP(lw, r)

		log.Info().
			Str("request", fmt.Sprintf("%s %s", r.Method, r.RequestURI)).
			Str("status", fmt.Sprint(lw.StatusCode)).
			Str("duration", fmt.Sprint(time.Since(t))).
			Send()
	})
}
