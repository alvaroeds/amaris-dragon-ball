package repository

import (
	"encoding/json"
	"time"

	"github.com/alvaroeds/amaris-dragon-ball/internal/infrastructure/db/cache/redis"
	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/domain"
)

type CacheRepository struct {
	client *redis.Client
}

// NewCacheRepository crea un repositorio para manejar cache de personajes
func NewCacheRepository(client *redis.Client) domain.CacheRepository {
	return &CacheRepository{
		client: client,
	}
}

// GetCharacter busca un personaje en cache por nombre
func (r *CacheRepository) GetCharacter(name string) (*domain.Character, error) {
	cacheKey := r.characterCacheKey(name)
	cached, err := r.client.Get(cacheKey)
	if err != nil {
		return nil, err
	}

	var character domain.Character
	if err := json.Unmarshal([]byte(cached), &character); err != nil {
		return nil, err
	}

	return &character, nil
}

// SaveCharacter guarda un personaje en cache por 10 minutos
func (r *CacheRepository) SaveCharacter(character *domain.Character) error {
	cacheKey := r.characterCacheKey(character.Name)
	data, err := json.Marshal(character)
	if err != nil {
		return err
	}

	return r.client.Set(cacheKey, data, 10*time.Minute)
}

// GetSearchResult busca un resultado de búsqueda previo en cache
func (r *CacheRepository) GetSearchResult(name string) (*domain.SearchResult, error) {
	cacheKey := r.searchCacheKey(name)
	cached, err := r.client.Get(cacheKey)
	if err != nil {
		return nil, err
	}

	var searchResult domain.SearchResult
	if err := json.Unmarshal([]byte(cached), &searchResult); err != nil {
		return nil, err
	}

	return &searchResult, nil
}

// SaveSearchResult guarda resultado de búsqueda en cache por 30 minutos
// Solo para búsquedas sin coincidencia exacta (para evitar consultas repetidas a la API)
func (r *CacheRepository) SaveSearchResult(name string, result *domain.SearchResult) error {
	cacheKey := r.searchCacheKey(name)
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return r.client.Set(cacheKey, data, 30*time.Minute)
}

// Genero las keys para organizar el cache
func (r *CacheRepository) characterCacheKey(name string) string {
	return "character:name:" + name
}

func (r *CacheRepository) searchCacheKey(name string) string {
	return "api:search:" + name
}
