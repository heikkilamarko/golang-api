// Package routes provides routing functionality.
package routes

import (
	"database/sql"
	"net/http"
	"products-api/app/config"
	"products-api/app/routes/products"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// RegisterRoutes func
func RegisterRoutes(router *mux.Router, config *config.Config, logger *zerolog.Logger) error {
	db, err := sql.Open("pgx", config.DBConnectionString)
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

	pc := products.NewController(
		products.NewSQLRepository(db, logger),
	)

	router.HandleFunc("/products", pc.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", pc.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id:[0-9]+}", pc.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/{id:[0-9]+}", pc.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products/{id:[0-9]+}", pc.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/products/pricerange", pc.GetPriceRange).Methods(http.MethodGet)

	return nil
}
