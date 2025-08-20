package models

import (
	"strings"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/utils"
	"gorm.io/gorm"
)

type Veterinarian struct {
	gorm.Model
	Name               string        `json:"name" gorm:"not null"`
	Email              string        `json:"email" gorm:"not null"`
	Phone              string        `json:"phone" gorm:"not null"`
	VeterinaryCenterID uint          `json:"veterinary_center_id" gorm:"not null"`
	Appointments       []Appointment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (v *Veterinarian) ValidateAndFormat() error {
	const minimal_name_length = 3
	const minimal_email_length = 7
	const minimal_phone_length = 7

	v.Name = strings.TrimSpace(v.Name)
	v.Email = strings.TrimSpace(v.Email)
	v.Phone = strings.TrimSpace(v.Phone)
	if v.Name == "" || v.Email == "" || v.Phone == "" {
		return gorm.ErrInvalidData
	}
	if len(v.Name) < minimal_name_length || len(v.Email) < minimal_email_length || len(v.Phone) < minimal_phone_length {
		return gorm.ErrInvalidData
	}
	if !utils.IsValidEmail(v.Email) {
		return gorm.ErrInvalidData
	}
	if !utils.IsValidPhoneNumber(v.Phone) {
		return gorm.ErrInvalidData
	}

	var center VeterinaryCenter
	err := database.DB.First(&center, v.VeterinaryCenterID).Error
	if err != nil {
		return gorm.ErrInvalidData
	}

	return nil
}

func (v *Veterinarian) BeforeSave(tx *gorm.DB) error {
	return v.ValidateAndFormat()
}

func (v *Veterinarian) BeforeUpdate(tx *gorm.DB) error {
	return v.ValidateAndFormat()
}
