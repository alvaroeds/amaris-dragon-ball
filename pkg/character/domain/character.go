package domain

import "time"

// Character representa un personaje de Dragon Ball en nuestro sistema
// Solo guardamos campos esenciales: ID, Nombre, Raza + 3 campos adicionales de la API
type Character struct {
	ID          int       `json:"id" db:"id"`
	ExternalID  int       `json:"external_id" db:"external_id"`
	Name        string    `json:"name" db:"name"`
	Race        string    `json:"race" db:"race"`               // Dato básico
	Ki          string    `json:"ki" db:"ki"`                   // Campo adicional 1
	Description string    `json:"description" db:"description"` // Campo adicional 2
	Image       string    `json:"image" db:"image"`             // Campo adicional 3
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
