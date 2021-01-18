// Package app provides application level functionality.
package app

import (
	"database/sql"
	"net/http"
	"products-api/app/config"
	"products-api/app/routes/products"
	"time"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/ory/graceful"
	"github.com/rs/cors"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", a.Config.PostgresConnectionString())
	if err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	pr := products.NewSQLRepository(db, a.Logger)
	pc := products.NewController(pr)

	router := mux.NewRouter()

	router.Use(
		middleware.Logger(a.Logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.APIKey(a.Config.APIKey, a.Config.APIKeyHeader),
		middleware.Timeout(a.Config.RequestTimeout),
	)

	router.HandleFunc("/products", pc.GetProducts).
		Methods("GET")

	router.HandleFunc("/products", pc.CreateProduct).
		Methods("POST")

	router.HandleFunc("/products/{id:[0-9]+}", pc.GetProduct).
		Methods("GET")

	router.HandleFunc("/products/{id:[0-9]+}", pc.UpdateProduct).
		Methods("PUT")

	router.HandleFunc("/products/{id:[0-9]+}", pc.DeleteProduct).
		Methods("DELETE")

	router.HandleFunc("/products/pricerange", pc.GetPriceRange).
		Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	var handler http.Handler = router

	if a.Config.CORSEnabled {
		handler = cors.AllowAll().Handler(router)
	}

	addr := a.Config.ServerAddr()

	server := graceful.WithDefaults(&http.Server{
		Addr:    addr,
		Handler: handler})

	a.Logger.Info().Msgf("Application running at %s", addr)

	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	a.Logger.Info().Msg("Application shutdown gracefully")
}
