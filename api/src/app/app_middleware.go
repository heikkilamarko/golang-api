package app

import (
	"fmt"
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
	"time"

	"github.com/rs/zerolog/log"
)

func (a *App) loggerMiddleware(next http.Handler) http.Handler {
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

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expected := a.Config.APIKey
		actual := r.Header.Get(constants.HeaderAPIKey)

		if actual != expected {
			utils.WriteUnauthorized(w, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *App) recoveryMiddleware(next http.Handler) http.Handler {
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteNotFound(w, nil)
}
