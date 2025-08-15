package models

import (
	"strings"

	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	Name               string `json:"name" gorm:"not null"`
	SpeciesID          uint
	BreedID            uint
	OwnerID            uint
	VeterinaryCenterID uint
	BehaviorNotes      []BehaviorNote `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HealthNotes        []HealthNote `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Appointments       []Appointment `gorm:"many2many:appointment_animals;"`
}

func (a *Animal) ValidateAndFormat() error {
	const minimal_name_length = 3
	a.Name = strings.TrimSpace(a.Name)
	if a.Name == "" || len(a.Name) < minimal_name_length {
		return gorm.ErrInvalidData
	}
	return nil
}

func (a *Animal) BeforeSave(tx *gorm.DB) error {
	return a.ValidateAndFormat()
}
