package v1

import (
	"follwit/controllers"
	"github.com/go-chi/chi/v5"
)

/*
RegisterBoilerplateRoutes
临时使用注册的方案，后期更换为配置文件
*/
func RegisterBoilerplateRoutes(router chi.Router) {
	router.Route("/boilerplate", func(r chi.Router) {
		r.Get("/{id}", controllers.GetBoilerplate)
	})
}
