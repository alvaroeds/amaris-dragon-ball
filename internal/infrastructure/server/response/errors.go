package response

import (
	"encoding/json"
	"net/http"
)

// Error envía una respuesta de error con tipo y mensaje
func Error(w http.ResponseWriter, statusCode int, errorType, message string) error {
	data := Map{
		"error":   errorType,
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// InternalError envía un error 500 del servidor
func InternalError(w http.ResponseWriter) error {
	return Error(w, http.StatusInternalServerError, "internal_error", "Internal server error")
}

// InvalidJSON envía un error 400 por JSON malformado
func InvalidJSON(w http.ResponseWriter) error {
	return Error(w, http.StatusBadRequest, "invalid_json", "Invalid or malformed JSON")
}

// ValidationError envía un error 400 de validación con mensaje personalizado
func ValidationError(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusBadRequest, "validation_error", message)
}
