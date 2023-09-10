package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetBoilerplate(writer http.ResponseWriter, request *http.Request) {
	boilerplateId := chi.URLParam(request, "id")

	_, err := writer.Write([]byte("get" + boilerplateId))
	if err != nil {
		return
	}
}
