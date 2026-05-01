package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Get allowed origins from environment or use default
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173"
	}

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(allowedOrigins, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health endpoint
	r.Get("/api/health", healthHandler)

	return r
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		log.Printf("Failed to encode health response: %v", err)
	}
}
