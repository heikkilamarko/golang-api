// Package app provides application level functionality.
package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"products-api/app/config"
	"products-api/app/routes/products"

	"github.com/gorilla/mux"
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

	pr := products.NewSQLRepository(db, a.Logger)
	pc := products.NewController(pr)

	router := mux.NewRouter()

	router.Use(
		a.loggerMiddleware,
		a.recoveryMiddleware,
		a.authMiddleware,
		a.timeoutMiddleware,
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

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	var handler http.Handler = router

	if a.Config.CorsEnabled {
		handler = cors.AllowAll().Handler(router)
	}

	addr := fmt.Sprintf(":%s", a.Config.Port)

	a.Logger.Info().Msgf("Application running at %s", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}
}
