package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goryuken/goryuken-core/pkg/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.RegisterAuthRoutes(r)
	http.ListenAndServe(":3000", r)
}
