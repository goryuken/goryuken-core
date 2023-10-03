package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/goryuken/goryuken-core/pkg/controllers"
)

var RegisterAuthRoutes = func(router *chi.Mux) {
	router.Post("/auth/signup", controllers.Signup)
	router.Post("/auth/login", controllers.Login)
	router.Post("/auth/logout", controllers.Logout)

}
