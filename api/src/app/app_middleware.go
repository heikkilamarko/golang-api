package app

import (
	"context"
	"fmt"
	"net/http"
	"products-api/app/constants"
	"time"

	"github.com/heikkilamarko/goutils"
)

func (a *App) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		lw := goutils.NewExtendedResponseWriter(w)
		next.ServeHTTP(lw, r)

		a.Logger.Info().
			Str("request", fmt.Sprintf("%s %s", r.Method, r.RequestURI)).
			Str("status", fmt.Sprint(lw.StatusCode)).
			Str("duration", fmt.Sprint(time.Since(t))).
			Send()
	})
}

func (a *App) recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				a.Logger.Error().Msgf("%s", err)
				goutils.WriteInternalError(w, nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expected := a.Config.APIKey
		actual := r.Header.Get(constants.HeaderAPIKey)

		if actual != expected {
			goutils.WriteUnauthorized(w, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *App) timeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), constants.RequestTimeout)
		defer cancel()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	goutils.WriteNotFound(w, nil)
}
