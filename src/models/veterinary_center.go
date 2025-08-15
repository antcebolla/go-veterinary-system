package models

import (
	"strings"

	"github.com/antcebolla/web-server/src/utils"
	"gorm.io/gorm"
)

type VeterinaryCenter struct {
	gorm.Model
	Name          string         `json:"name" gorm:"unique, not null"`
	Address       string         `json:"address" gorm:"not null"`
	Phone         string         `json:"phone"`
	Owners        []Owner        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Veterinarians []Veterinarian `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Animals       []Animal       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Appointments  []Appointment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (vc *VeterinaryCenter) ValidateAndFormat() error {
	const minimal_name_length = 3
	const minimal_address_length = 7

	vc.Name = strings.TrimSpace(vc.Name)
	vc.Address = strings.TrimSpace(vc.Address)
	vc.Phone = strings.TrimSpace(vc.Phone)
	if vc.Name == "" || vc.Address == "" {
		return gorm.ErrInvalidData
	}
	if len(vc.Name) < minimal_name_length || len(vc.Address) < minimal_address_length {
		return gorm.ErrInvalidData
	}
	if !(utils.IsValidPhoneNumber(vc.Phone)) {
		return gorm.ErrInvalidData
	}
	return nil
}

func (vc *VeterinaryCenter) BeforeSave(tx *gorm.DB) error {
	return vc.ValidateAndFormat()
}

func (vc *VeterinaryCenter) BeforeUpdate(tx *gorm.DB) error {
	return vc.ValidateAndFormat()
}
