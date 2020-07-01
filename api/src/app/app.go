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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// App struct
type App struct {
	Repository products.Repository
	Products   *products.Controller
	Router     *mux.Router
}

// NewApp func
func NewApp() *App {
	return &App{}
}

// Initialize method
func (a *App) Initialize() {
	a.initializeConfig()
	a.initializeLogger()
	a.initializeRepository()
	a.initializeRouter()
}

// Run method
func (a *App) Run() {

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

func (a *App) initializeConfig() {
	config.Load()
}

func (a *App) initializeLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
}

func (a *App) initializeRepository() {
	a.Repository = products.NewRepository()
	a.Repository.Initialize()
}

func (a *App) initializeRouter() {
	a.Router = mux.NewRouter()
	a.addMiddleware()
	a.addProductsRoutes()
	a.addNotFoundRoute()
}

func (a *App) addMiddleware() {
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recovery)
	a.Router.Use(middleware.Auth)
}

func (a *App) addProductsRoutes() {
	c := products.NewController(a.Repository)
	r := a.Router
	r.HandleFunc("/products", c.GetProducts).Methods("GET")
	r.HandleFunc("/products", c.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", c.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", c.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", c.DeleteProduct).Methods("DELETE")
	a.Products = c
}

func (a *App) addNotFoundRoute() {
	a.Router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			utils.WriteNotFound(w, nil)
		})
}
