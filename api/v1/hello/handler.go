package hello

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/just-arun/creative/helper"
	"github.com/just-arun/creative/models"
)

// api endpoint for hello will be handled here

func Route(ctx *models.HandlerCtx) func(c chi.Router) {
	return func(r chi.Router) {
		r.Get("/", hello(ctx))
		r.Post("/", create(ctx))
	}
}

func hello(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		name := helper.Req(r).Q("name")

		helper.Res(w).Success(http.StatusOK, "hello " + name )
	}
}

func create(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body createHelloBody

		err := helper.Req(r).B(&body)
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}

		helper.Res(w).Success(http.StatusCreated, body)
	}
}

