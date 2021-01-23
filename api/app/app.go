// Package app provides application level functionality.
package app

import (
	"net/http"
	"products-api/app/config"
	"products-api/app/routes"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/ory/graceful"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

// App struct
type App struct {
	Config *config.Config
	Logger *zerolog.Logger
}

// New func
func New(c *config.Config, l *zerolog.Logger) *App {
	return &App{c, l}
}

// Run method
func (a *App) Run() {
	router := mux.NewRouter()

	router.Use(
		middleware.Logger(a.Logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.APIKey(a.Config.APIKey, a.Config.APIKeyHeader),
		middleware.Timeout(a.Config.RequestTimeout),
	)

	if err := routes.RegisterRoutes(router, a.Config, a.Logger); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	var handler http.Handler = router

	if a.Config.CORSEnabled {
		handler = cors.AllowAll().Handler(router)
	}

	server := graceful.WithDefaults(&http.Server{
		Addr:    a.Config.Address,
		Handler: handler})

	a.Logger.Info().Msgf("Application running at %s", a.Config.Address)

	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	a.Logger.Info().Msg("Application shutdown gracefully")
}
