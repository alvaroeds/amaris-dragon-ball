package domain

// ExternalCharacter representa la respuesta de la API externa de DragonBall
type ExternalCharacter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Character   string `json:"character"` // Nombre alternativo que puede venir en la API
	Ki          string `json:"ki"`
	MaxKi       string `json:"maxKi"`
	Race        string `json:"race"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Affiliation string `json:"affiliation"`
}
