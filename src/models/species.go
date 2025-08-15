package models

import (
	"strings"

	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Name    string `json:"name" gorm:"unique"`
	Breeds  []Breed
	Animals []Animal
}

func (s *Species) ValidateAndFormat() error {
	const minimal_name_length = 3

	s.Name = strings.TrimSpace(s.Name)
	if s.Name == "" || len(s.Name) < minimal_name_length {
		return gorm.ErrInvalidData
	}
	return nil
}

func (s *Species) BeforeSave(tx *gorm.DB) error {
	return s.ValidateAndFormat()
}
