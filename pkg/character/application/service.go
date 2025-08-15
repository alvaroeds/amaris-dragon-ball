package application

import (
	"fmt"

	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/domain"
)

type Service struct {
	repository  domain.CharacterRepository
	externalAPI domain.ExternalCharacterAPI
	cache       domain.CacheRepository
}

// NewService crea un nuevo servicio de caracteres
func NewService(repository domain.CharacterRepository, externalAPI domain.ExternalCharacterAPI, cache domain.CacheRepository) domain.CharacterService {
	return &Service{
		repository:  repository,
		externalAPI: externalAPI,
		cache:       cache,
	}
}

// CreateCharacter busca o crea un personaje siguiendo esta lógica:
// Cache -> Base datos -> API externa
func (s *Service) CreateCharacter(name string) (*domain.CreateCharacterResult, error) {
	if name == "" {
		return nil, fmt.Errorf("el nombre es requerido")
	}

	// Primero busco en cache para respuesta rápida
	if character, err := s.cache.GetCharacter(name); err == nil {
		return &domain.CreateCharacterResult{
			Character: character,
			IsNew:     false,
		}, nil
	}

	// Si no está en cache, busco en base de datos
	character, err := s.repository.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("error al buscar personaje en BD: %w", err)
	}

	// Si lo encontré en BD, lo guardo en cache para la próxima
	if character != nil {
		s.cache.SaveCharacter(character)
		return &domain.CreateCharacterResult{
			Character: character,
			IsNew:     false,
		}, nil
	}

	// Verifico si ya busqué este nombre en la API antes (sin resultado exacto)
	if searchResult, err := s.cache.GetSearchResult(name); err == nil {
		return s.handleNoExactMatch(searchResult, name)
	}

	// Como último recurso, consulto la API externa
	searchResult, err := s.externalAPI.SearchCharacterByName(name)
	if err != nil {
		return nil, fmt.Errorf("error al consultar API externa: %w", err)
	}

	// Si encontré coincidencia exacta, creo el personaje
	if searchResult.IsExactMatch {
		return s.createFromExactMatch(searchResult.ExactMatch)
	}

	// Si no hay coincidencia exacta, guardo el resultado para no volver a consultar
	s.cache.SaveSearchResult(name, searchResult)

	return s.handleNoExactMatch(searchResult, name)
}

// createFromExactMatch cuando encuentro el personaje exacto en la API, lo creo en mi BD
func (s *Service) createFromExactMatch(exactMatch *domain.ExternalCharacter) (*domain.CreateCharacterResult, error) {
	// Convierto el personaje externo a mi formato interno
	character := domain.ExternalToCharacter(exactMatch)

	// Lo guardo en mi base de datos
	if err := s.repository.Create(character); err != nil {
		return nil, fmt.Errorf("error al guardar personaje en BD: %w", err)
	}

	// También lo dejo en cache para consultas futuras
	s.cache.SaveCharacter(character)

	return &domain.CreateCharacterResult{
		Character: character,
		IsNew:     true,
	}, nil
}

// handleNoExactMatch cuando no encuentro el nombre exacto, devuelvo sugerencias
func (s *Service) handleNoExactMatch(searchResult *domain.SearchResult, searchName string) (*domain.CreateCharacterResult, error) {
	return &domain.CreateCharacterResult{
		IsNew:               false,
		AvailableCharacters: searchResult.AvailableCharacters,
		Error:               fmt.Sprintf("No exact match found for '%s'", searchName),
	}, nil
}
