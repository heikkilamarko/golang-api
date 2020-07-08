package app

import (
	"fmt"
	"net/http"
	"products-api/app/config"
	"products-api/app/middleware"
	"products-api/app/routes/products"
	"products-api/app/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

// App struct
type App struct {
	ProductsRepository products.Repository
	ProductsController *products.Controller
	Router             *mux.Router
}

// New func
func New() *App {
	return &App{}
}

// Run method
func (a *App) Run() {
	a.configureControllers()
	a.configureRouter()

	addr := fmt.Sprintf(":%s", config.Config.Port)

	log.Info().Msgf("Application running at %s", addr)

	var handler http.Handler = a.Router

	if config.Config.CorsEnabled {
		handler = cors.AllowAll().Handler(a.Router)
	}

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func (a *App) configureControllers() {
	a.ProductsRepository = products.NewRepository()
	a.ProductsController = products.NewController(a.ProductsRepository)
}

func (a *App) configureRouter() {
	r := mux.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)
	r.Use(middleware.Auth)

	c := a.ProductsController

	r.HandleFunc("/products", c.GetProducts).Methods("GET")
	r.HandleFunc("/products", c.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", c.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", c.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", c.DeleteProduct).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			utils.WriteNotFound(w, nil)
		})

	a.Router = r
}
