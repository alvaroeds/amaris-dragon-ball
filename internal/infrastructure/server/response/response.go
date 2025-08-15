package response

import (
	"encoding/json"
	"net/http"
)

// Map para respuestas flexibles
type Map map[string]interface{}

// Send envía una respuesta JSON con el código de estado especificado
func Send(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(Map{"result": data})
}

// OK envía una respuesta 200 con datos
func OK(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusOK, data)
}

// Created envía una respuesta 201 para recursos creados
func Created(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusCreated, data)
}

// BadRequest envía un 400 con mensaje de error y sugerencias
func BadRequest(w http.ResponseWriter, message string, suggestions []string) error {
	data := Map{
		"error":       message,
		"suggestions": suggestions,
	}
	return Send(w, http.StatusBadRequest, data)
}

// Conflict envía un 409 cuando el recurso ya existe
func Conflict(w http.ResponseWriter, message string, data interface{}) error {
	response := Map{
		"error": message,
		"data":  data,
	}
	return Send(w, http.StatusConflict, response)
}
