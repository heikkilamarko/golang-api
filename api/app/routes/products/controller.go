// Package products provides product functionality.
package products

import (
	"database/sql"
	"products-api/app/config"

	"github.com/rs/zerolog"
)

// Controller struct
type Controller struct {
	config     *config.Config
	logger     *zerolog.Logger
	db         *sql.DB
	repository *repository
}

// NewController func
func NewController(config *config.Config, logger *zerolog.Logger, db *sql.DB) *Controller {
	repository := newRepository(db, logger)
	return &Controller{config, logger, db, repository}
}
