package routers

import (
	"follwit/routers/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoute() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(middlewares.Cors())

	RegisterRoutes(router)

	return router
}
