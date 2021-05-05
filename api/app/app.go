// Package app provides application level functionality.
package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"products-api/app/config"
	"products-api/app/routes/products"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// App struct
type App struct {
	config *config.Config
	logger *zerolog.Logger
	db     *sql.DB
	router *mux.Router
	server *http.Server
}

// New func
func New(c *config.Config, l *zerolog.Logger) *App {
	return &App{config: c, logger: l}
}

// Run method
func (a *App) Run() {
	a.logInfo("application is starting up...")

	if err := a.initDB(); err != nil {
		a.logFatal(err)
	}

	a.initRouter()

	a.registerRoutes()

	a.initServer()

	if err := a.serve(); err != nil {
		a.logFatal(err)
	}

	a.logInfo("application is shut down")
}

func (a *App) initDB() error {

	db, err := sql.Open("pgx", a.config.DBConnectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *App) initRouter() {

	router := mux.NewRouter()

	router.Use(
		middleware.Logger(a.logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.APIKey(a.config.APIKey, a.config.APIKeyHeader),
		middleware.Timeout(a.config.RequestTimeout),
	)

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	a.router = router
}

func (a *App) initServer() {

	var handler http.Handler = a.router

	if a.config.CORSEnabled {
		handler = cors.AllowAll().Handler(a.router)
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         a.config.Address,
		Handler:      handler,
	}

	a.server = server
}

func (a *App) registerRoutes() {

	c := products.NewController(a.config, a.logger, a.db)

	a.router.HandleFunc("/products", c.GetProducts).Methods(http.MethodGet)
	a.router.HandleFunc("/products", c.CreateProduct).Methods(http.MethodPost)
	a.router.HandleFunc("/products/{id:[0-9]+}", c.GetProduct).Methods(http.MethodGet)
	a.router.HandleFunc("/products/{id:[0-9]+}", c.UpdateProduct).Methods(http.MethodPut)
	a.router.HandleFunc("/products/{id:[0-9]+}", c.DeleteProduct).Methods(http.MethodDelete)
	a.router.HandleFunc("/products/pricerange", c.GetPriceRange).Methods(http.MethodGet)
}

func (a *App) serve() error {

	var (
		s = make(chan os.Signal)
		e = make(chan error)
	)

	go func() {
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)

		<-s

		a.logInfo("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = a.server.Shutdown(ctx)
		_ = a.db.Close()

		e <- nil
	}()

	a.logInfo("application is running at %s", a.config.Address)

	if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-e
}

func (a *App) logInfo(msg string, v ...interface{}) {
	a.logger.Info().Msgf(msg, v...)
}

func (a *App) logFatal(err error) {
	a.logger.Fatal().Err(err).Send()
}
