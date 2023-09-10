package routers

import (
	v1 "follwit/routers/v1"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(router *chi.Mux) {
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Running"))
		if err != nil {
			return
		}
	})

	v1.RegisterV1Routes(router)
}
