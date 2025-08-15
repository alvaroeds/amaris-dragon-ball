package http

import (
	"net/http"

	"github.com/alvaroeds/amaris-dragon-ball/internal/health"
	character_handler "github.com/alvaroeds/amaris-dragon-ball/pkg/character/infrastructure/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Routes function sets route handlers for Dragon Ball API.
func Routes(characterHandler *character_handler.Handler, healthHandler *health.Handler) http.Handler {
	r := chi.NewRouter()

	// CORS configuration
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"Cache-Control",
			"X-Requested-With",
			"Origin",
		},
		ExposedHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check endpoint
	r.Get("/health", healthHandler.CheckHealth)

	// API routes
	r.Mount("/api/v1", getRoutesV1(characterHandler))

	return r
}

// getRoutesV1 defines all v1 API routes
func getRoutesV1(characterHandler *character_handler.Handler) http.Handler {
	r := chi.NewMux()

	// Character routes
	r.Route("/characters", func(r chi.Router) {
		r.Post("/", characterHandler.CreateCharacter)
	})

	return r
}
