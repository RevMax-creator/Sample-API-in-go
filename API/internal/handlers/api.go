package handlers

import (
	middleware "API/internal/middleware" // Your custom authorization middleware package

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capitalise initial tell compiler that it can be imported into other packages
// Lowercase will be a private function which can only be used in this package itself
// Middleware is a function that's get called before the primary function which handles the endpoint
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes) // Adding middleware to a point

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})

}
