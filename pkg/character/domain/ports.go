package domain

// CharacterRepository define los métodos que debe implementar el repositorio de personajes
type CharacterRepository interface {
	GetByName(name string) (*Character, error)
	Create(character *Character) error
}

// ExternalCharacterAPI define los métodos para consultar APIs externas de personajes
type ExternalCharacterAPI interface {
	SearchCharacterByName(name string) (*SearchResult, error)
}

// CharacterService define los métodos del servicio de aplicación de personajes
type CharacterService interface {
	CreateCharacter(name string) (*CreateCharacterResult, error)
}

// CacheRepository define operaciones de cache para caracteres
type CacheRepository interface {
	GetCharacter(name string) (*Character, error)
	SaveCharacter(character *Character) error
	GetSearchResult(name string) (*SearchResult, error)
	SaveSearchResult(name string, result *SearchResult) error
}
