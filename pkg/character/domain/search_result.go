package domain

// SearchResult representa el resultado de una búsqueda en la API externa
type SearchResult struct {
	ExactMatch          *ExternalCharacter `json:"exact_match,omitempty"`
	AvailableCharacters []string           `json:"available_characters,omitempty"`
	IsExactMatch        bool               `json:"is_exact_match"`
}

// CreateCharacterResult representa el resultado de crear un personaje
type CreateCharacterResult struct {
	Character           *Character `json:"character,omitempty"`
	IsNew               bool       `json:"is_new"`
	Source              string     `json:"source"` // "cache", "database", "external_api"
	AvailableCharacters []string   `json:"available_characters,omitempty"`
	Error               string     `json:"error,omitempty"`
}

// DataSource constantes para indicar fuente de datos
const (
	SourceCache       = "cache"
	SourceDatabase    = "database"
	SourceExternalAPI = "external_api"
)
