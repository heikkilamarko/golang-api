package service

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"product-api/internal/adapters"
	"product-api/internal/application"
	"product-api/internal/application/command"
	"product-api/internal/application/query"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type config struct {
	App                string
	Address            string
	APIKey             string
	APIKeyHeader       string
	DBConnectionString string
	CORSEnabled        bool
	LogLevel           string
}

type Service struct {
	config *config
	logger *zerolog.Logger
	db     *sql.DB
	app    *application.Application
	server *http.Server
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s.loadConfig()
	s.initLogger()

	s.logInfo("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.logFatal(err)
	}

	s.initApplication()
	s.initHTTPServer()

	if err := s.serve(ctx); err != nil {
		s.logFatal(err)
	}

	s.logInfo("application is shut down")
}

func (s *Service) loadConfig() {
	s.config = &config{
		App:                env("APP_NAME", ""),
		Address:            env("APP_ADDRESS", ":8080"),
		APIKey:             env("APP_API_KEY", ""),
		APIKeyHeader:       env("APP_API_KEY_HEADER", "X-Api-Key"),
		DBConnectionString: env("APP_DB_CONNECTION_STRING", ""),
		CORSEnabled:        env("APP_CORS_ENABLED", "") == "true",
		LogLevel:           env("APP_LOG_LEVEL", "warn"),
	}
}

func (s *Service) initLogger() {
	level, err := zerolog.ParseLevel(s.config.LogLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("app", s.config.App).
		Logger()

	s.logger = &logger
}

func (s *Service) initDB(ctx context.Context) error {
	db, err := sql.Open("pgx", s.config.DBConnectionString)
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

	s.db = db

	return nil
}

func (s *Service) initApplication() {
	productRepository := adapters.NewProductPostgresRepository(s.db)

	s.app = &application.Application{
		Commands: application.Commands{
			CreateProduct: command.NewCreateProductHandler(productRepository),
			UpdateProduct: command.NewUpdateProductHandler(productRepository),
			DeleteProduct: command.NewDeleteProductHandler(productRepository),
		},
		Queries: application.Queries{
			GetProducts:   query.NewGetProductsHandler(productRepository),
			GetProduct:    query.NewGetProductHandler(productRepository),
			GetPriceRange: query.NewGetPriceRangeHandler(productRepository),
		},
	}
}

func (s *Service) initHTTPServer() {
	router := chi.NewRouter()

	router.Use(
		middleware.Recoverer,
		adapters.APIKey(s.config.APIKey, s.config.APIKeyHeader),
	)

	productHandlers := adapters.NewProductHTTPHandlers(s.app, s.logger)

	router.MethodFunc(http.MethodGet, "/products", productHandlers.GetProducts)
	router.MethodFunc(http.MethodPost, "/products", productHandlers.CreateProduct)
	router.MethodFunc(http.MethodGet, "/products/{id:[0-9]+}", productHandlers.GetProduct)
	router.MethodFunc(http.MethodPut, "/products/{id:[0-9]+}", productHandlers.UpdateProduct)
	router.MethodFunc(http.MethodDelete, "/products/{id:[0-9]+}", productHandlers.DeleteProduct)
	router.MethodFunc(http.MethodGet, "/products/pricerange", productHandlers.GetPriceRange)

	router.NotFound(adapters.NotFound)

	var handler http.Handler = router

	if s.config.CORSEnabled {
		handler = cors.AllowAll().Handler(router)
	}

	s.server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         s.config.Address,
		Handler:      handler,
	}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logInfo("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.server.Shutdown(ctx)
		_ = s.db.Close()

		errChan <- nil
	}()

	s.logInfo("application is running at %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}
