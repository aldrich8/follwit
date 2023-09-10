package v1

import "github.com/go-chi/chi/v5"

func RegisterV1Routes(router chi.Router) {
	RegisterBoilerplateRoutes(router)
}
