package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RouterInitializer(cont Controllers) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))

	Routes(cont, router)
	return router
}