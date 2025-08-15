package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/server/response"
	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/domain"
)

type Handler struct {
	service domain.CharacterService
}

// NewHandler crea un nuevo handler HTTP
func NewHandler(service domain.CharacterService) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateCharacter maneja las peticiones POST para crear/buscar personajes
func (h *Handler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	// Parseo el JSON del request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.InvalidJSON(w)
		return
	}

	// Valido que me hayan enviado un nombre
	if req.Name == "" {
		response.ValidationError(w, "Character name is required")
		return
	}

	// Llamo al servicio para buscar/crear el personaje
	result, err := h.service.CreateCharacter(req.Name)
	if err != nil {
		response.InternalError(w)
		return
	}

	// Si no encontré coincidencia exacta, devuelvo sugerencias
	if result.Error != "" {
		response.BadRequest(w, result.Error, result.AvailableCharacters)
		return
	}

	// Si el personaje ya existía, devuelvo conflicto
	if !result.IsNew {
		response.Conflict(w, "Character already exists", result.Character)
		return
	}

	// Si lo creé exitosamente, devuelvo 201
	response.Created(w, result.Character)
}
