package route

import (
	"github.com/BelyaevEI/auth-server/internal/handlers"
	"github.com/go-chi/chi"
)

func NewRoute(hanlder handlers.Handler) *chi.Mux {

	// New router
	route := chi.NewRouter()

	// Handlers
	route.Post("/api/v1/create", hanlder.CreateNewToken) // Create new tokens

	return route
}
