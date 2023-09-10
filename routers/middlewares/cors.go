package middlewares

import (
	"github.com/go-chi/cors"
	"net/http"
)

func Cors() func(handler http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		/**
		 * Allowing all origin is dangerous
		 */
		AllowedOrigins: []string{"*"},

		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},

		ExposedHeaders: []string{"Link"},

		AllowCredentials: true,

		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
	})
}
