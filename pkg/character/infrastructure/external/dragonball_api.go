package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/domain"
)

type DragonBallAPIClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewDragonBallAPIClient crea un nuevo cliente para la API de DragonBall
func NewDragonBallAPIClient(baseURL string, timeoutSeconds int) domain.ExternalCharacterAPI {
	return &DragonBallAPIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
	}
}

// SearchCharacterByName busca personajes por nombre en la API externa
func (c *DragonBallAPIClient) SearchCharacterByName(name string) (*domain.SearchResult, error) {
	// Construir URL con parámetro de búsqueda
	apiURL := fmt.Sprintf("%s/characters?name=%s", c.baseURL, url.QueryEscape(name))

	// Realizar petición HTTP
	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error al llamar API externa: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API externa retornó status %d", resp.StatusCode)
	}

	// Decodificar respuesta - la API devuelve un array
	var characters []domain.ExternalCharacter
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, fmt.Errorf("error al decodificar respuesta de API externa: %w", err)
	}

	// Si no se encontraron personajes
	if len(characters) == 0 {
		return &domain.SearchResult{
			IsExactMatch:        false,
			AvailableCharacters: []string{},
		}, nil
	}

	// Normalizar nombres de todos los personajes
	for i := range characters {
		c.normalizeCharacterName(&characters[i])
	}

	// Buscar coincidencia exacta (case insensitive)
	searchNameLower := strings.ToLower(name)
	var exactMatch *domain.ExternalCharacter
	var availableNames []string

	for i := range characters {
		charName := strings.ToLower(characters[i].Name)
		availableNames = append(availableNames, characters[i].Name)

		if charName == searchNameLower {
			exactMatch = &characters[i]
		}
	}

	return &domain.SearchResult{
		ExactMatch:          exactMatch,
		AvailableCharacters: availableNames,
		IsExactMatch:        exactMatch != nil,
	}, nil
}

// normalizeCharacterName arregla inconsistencias en los nombres de la API externa
func (c *DragonBallAPIClient) normalizeCharacterName(character *domain.ExternalCharacter) {
	// La API a veces usa el campo 'character' en vez de 'name'
	if character.Name == "" && character.Character != "" {
		character.Name = character.Character
	}
}
