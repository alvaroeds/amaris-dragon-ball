package domain

// ExternalToCharacter convierte un ExternalCharacter a Character
func ExternalToCharacter(ext *ExternalCharacter) *Character {
	return &Character{
		ExternalID:  ext.ID,
		Name:        ext.Name,
		Race:        ext.Race,        // Dato básico
		Ki:          ext.Ki,          // Campo adicional 1
		Description: ext.Description, // Campo adicional 2
		Image:       ext.Image,       // Campo adicional 3
	}
}
