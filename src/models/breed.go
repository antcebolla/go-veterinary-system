package models

import (
	"strings"

	"gorm.io/gorm"
)

type Breed struct {
	gorm.Model
	Name      string `json:"name"`
	SpeciesID uint
	Animals   []Animal
}

func (b *Breed) ValidateAndFormat() error {
	const minimal_name_length = 3
	b.Name = strings.TrimSpace(b.Name)
	if b.Name == "" || len(b.Name) < minimal_name_length {
		return gorm.ErrInvalidData
	}
	return nil
}
