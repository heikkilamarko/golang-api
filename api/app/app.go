// Package app provides application level functionality.
package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"products-api/app/config"
	"products-api/app/middleware"
	"products-api/app/products"
	"syscall"
	"time"

	"github.com/gorilla/mux"
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

// Run method
func (a *App) Run() {

	a.loadConfig()
	a.initLogger()

	a.logInfo("application is starting up...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := a.initDB(ctx); err != nil {
		a.logFatal(err)
	}

	a.initRouter()

	a.registerRoutes()

	a.initServer()

	if err := a.serve(ctx); err != nil {
		a.logFatal(err)
	}

	a.logInfo("application is shut down")
}

func (a *App) loadConfig() {
	a.config = config.Load()
}

func (a *App) initLogger() {

	level, err := zerolog.ParseLevel(a.config.LogLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("app", a.config.App).
		Logger()

	a.logger = &logger
}

func (a *App) initDB(ctx context.Context) error {

	db, err := sql.Open("pgx", a.config.DBConnectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
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
	)

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	a.router = router
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

func (a *App) serve(ctx context.Context) error {

	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		a.logInfo("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = a.server.Shutdown(ctx)
		_ = a.db.Close()

		errChan <- nil
	}()

	a.logInfo("application is running at %s", a.config.Address)

	if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}

func (a *App) logInfo(msg string, v ...interface{}) {
	a.logger.Info().Msgf(msg, v...)
}

func (a *App) logFatal(err error) {
	a.logger.Fatal().Err(err).Send()
}
