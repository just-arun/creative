package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/just-arun/creative/api/v1"
	"github.com/just-arun/creative/middleware"
)

func Api(port, env string) {
	r := chi.NewRouter()
	r.Use(middleware.MetaData())
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", v1.ApiV1(env))
		r.Route("/v2", v1.ApiV1(env))
	})
	log.Fatal(http.ListenAndServe(port, r))
}
