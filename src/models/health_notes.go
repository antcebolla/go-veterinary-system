package models

import (
	"strings"

	"gorm.io/gorm"
)

type HealthNote struct {
	gorm.Model
	Note     string `json:"health_note" gorm:"type:text not null; max:240"`
	AnimalID uint
}

func (h *HealthNote) ValidateAndFormat() error {
	max_length := 240
	h.Note = strings.TrimSpace(h.Note)
	if h.Note == "" || len(h.Note) > max_length {
		return gorm.ErrInvalidData
	}
	return nil
}
