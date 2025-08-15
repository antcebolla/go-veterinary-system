package models

import (
	"strings"

	"github.com/antcebolla/web-server/src/utils"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Name               string `json:"name" gorm:"not null"`
	Email              string `json:"email" gorm:"unique, not null"`
	Phone              string `json:"phone" gorm:"not null"`
	VeterinaryCenterID uint
	Animals            []Animal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (o *Owner) ValidateAndFormat() error {
	const minimal_name_length = 3
	const minimal_email_length = 7
	const minimal_phone_length = 7

	o.Name = strings.TrimSpace(o.Name)
	o.Email = strings.TrimSpace(o.Email)
	o.Phone = strings.TrimSpace(o.Phone)
	if o.Name == "" || o.Email == "" || o.Phone == "" {
		return gorm.ErrInvalidData
	}
	if len(o.Name) < minimal_name_length || len(o.Email) < minimal_email_length || len(o.Phone) < minimal_phone_length {
		return gorm.ErrInvalidData
	}
	if !utils.IsValidEmail(o.Email) {
		return gorm.ErrInvalidData
	}
	if !utils.IsValidPhoneNumber(o.Phone) {
		return gorm.ErrInvalidData
	}

	return nil
}

func (o *Owner) BeforeSave(tx *gorm.DB) error {
	return o.ValidateAndFormat()
}
