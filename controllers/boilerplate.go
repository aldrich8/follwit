package controllers

import (
	"follwit/models"
	"follwit/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func GetBoilerplate(writer http.ResponseWriter, request *http.Request) {
	boilerplateId := chi.URLParam(request, "id")

	_, err := writer.Write([]byte("get" + boilerplateId))
	if err != nil {
		return
	}
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

type BoilerplateResponse struct {
	Data       *models.Boilerplate `json:"data"`
	AppCode    int64               `json:"code,omitempty"`
	StatusText string              `json:"status"`
}

func (b BoilerplateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateBoilerplate(writer http.ResponseWriter, request *http.Request) {
	boilerplate := &models.Boilerplate{}

	if err := render.Bind(request, boilerplate); err != nil {
		_ = render.Render(writer, request, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Invalid request data",
			ErrorText:      err.Error(),
			AppCode:        400,
		})
	}

	if err := repository.Save(&boilerplate); err != nil {
		println("Error: ", err.Error())

		_ = render.Render(writer, request, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 500,
			StatusText:     "Internal server error",
			ErrorText:      err.Error(),
			AppCode:        500,
		})

	}

	_ = render.Render(writer, request, &BoilerplateResponse{
		Data:       boilerplate,
		AppCode:    200,
		StatusText: "Success",
	})
}

func FindBoilerplate(w http.ResponseWriter, r *http.Request) {
	boilerplateId := chi.URLParam(r, "id")

	boilerplate := &models.Boilerplate{}

	if err := repository.Find(&boilerplate, boilerplateId); err != nil {
		println("Error: ", err.Error())

		_ = render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 500,
			StatusText:     "Internal server error",
			ErrorText:      err.Error(),
			AppCode:        500,
		})
	}
}
