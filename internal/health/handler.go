package health

import (
	"net/http"

	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/server/response"
)

// Checker interfaz simple para verificar si un servicio responde
type Checker interface {
	Ping() error
}

// Handler maneja el endpoint de health check
type Handler struct {
	database Checker
	cache    Checker
}

// NewHandler crea un nuevo handler de health
func NewHandler(database, cache Checker) *Handler {
	return &Handler{
		database: database,
		cache:    cache,
	}
}

// CheckHealth verifica que PostgreSQL y Redis estén funcionando
func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	status := "healthy"
	services := response.Map{}

	// Pruebo PostgreSQL
	if err := h.database.Ping(); err != nil {
		status = "unhealthy"
		services["database"] = "error: " + err.Error()
	} else {
		services["database"] = "ok"
	}

	// Pruebo Redis
	if err := h.cache.Ping(); err != nil {
		status = "unhealthy"
		services["cache"] = "error: " + err.Error()
	} else {
		services["cache"] = "ok"
	}

	// Armo la respuesta
	result := response.Map{
		"status":   status,
		"services": services,
	}

	// Si algo falla devuelvo 503, sino 200
	if status == "unhealthy" {
		response.Send(w, http.StatusServiceUnavailable, result)
	} else {
		response.OK(w, result)
	}
}
