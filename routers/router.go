package routers

import (
	"follwit/routers/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func SetupRoute() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middlewares.Cors())

	router.Use(render.SetContentType(render.ContentTypeJSON))

	RegisterRoutes(router)

	return router
}
