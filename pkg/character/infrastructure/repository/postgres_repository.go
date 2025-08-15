package repository

import (
	"database/sql"
	"time"

	"github.com/alvaroeds/amaris-dragon-ball/pkg/character/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository crea un nuevo repositorio PostgreSQL
func NewPostgresRepository(db *sql.DB) domain.CharacterRepository {
	return &PostgresRepository{
		db: db,
	}
}

// GetByName obtiene un personaje por su nombre
func (r *PostgresRepository) GetByName(name string) (*domain.Character, error) {
	query := `
		SELECT id, external_id, name, race, ki, description, image, created_at, updated_at
		FROM characters
		WHERE LOWER(name) = LOWER($1)
		LIMIT 1`

	character := &domain.Character{}
	err := r.db.QueryRow(query, name).Scan(
		&character.ID,
		&character.ExternalID,
		&character.Name,
		&character.Race,
		&character.Ki,
		&character.Description,
		&character.Image,
		&character.CreatedAt,
		&character.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return character, err
}

// Create inserta un nuevo personaje en la base de datos
func (r *PostgresRepository) Create(character *domain.Character) error {
	query := `
		INSERT INTO characters (external_id, name, race, ki, description, image, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	now := time.Now()
	character.CreatedAt = now
	character.UpdatedAt = now

	err := r.db.QueryRow(
		query,
		character.ExternalID,
		character.Name,
		character.Race,
		character.Ki,
		character.Description,
		character.Image,
		character.CreatedAt,
		character.UpdatedAt,
	).Scan(&character.ID)

	return err
}
