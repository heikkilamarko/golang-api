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
	"github.com/go-chi/cors"
	"golang.org/x/exp/slog"

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
	logger *slog.Logger
	db     *sql.DB
	app    *application.Application
	server *http.Server
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s.loadConfig()
	s.initLogger()

	s.logger.Info("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.logger.Error(err.Error())
		os.Exit(1)
	}

	s.initApplication()
	s.initHTTPServer()

	if err := s.serve(ctx); err != nil {
		s.logger.Error(err.Error())
		os.Exit(1)
	}

	s.logger.Info("application is shut down")
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
	level := slog.LevelInfo

	level.UnmarshalText([]byte(s.config.LogLevel))

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stderr, opts).
		WithAttrs([]slog.Attr{
			slog.String("app", s.config.App),
		})

	logger := slog.New(handler)

	slog.SetDefault(logger)

	s.logger = logger
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

	router.Use(middleware.Recoverer)

	if s.config.CORSEnabled {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodOptions,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
	}

	router.Use(adapters.APIKey(s.config.APIKey, s.config.APIKeyHeader))

	h := adapters.NewProductHTTPHandlers(s.app, s.logger)

	router.Get("/products", h.GetProducts)
	router.Post("/products", h.CreateProduct)
	router.Get("/products/{id:[0-9]+}", h.GetProduct)
	router.Put("/products/{id:[0-9]+}", h.UpdateProduct)
	router.Delete("/products/{id:[0-9]+}", h.DeleteProduct)
	router.Get("/products/pricerange", h.GetPriceRange)

	router.NotFound(adapters.NotFound)

	s.server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         s.config.Address,
		Handler:      router,
	}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logger.Info("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.server.Shutdown(ctx)
		_ = s.db.Close()

		errChan <- nil
	}()

	s.logger.Info("application is running at " + s.server.Addr)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}
