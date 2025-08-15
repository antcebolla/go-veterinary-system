package models

import (
	"strings"

	"gorm.io/gorm"
)

type BehaviorNote struct {
	gorm.Model
	Note     string `json:"behavior_note" gorm:"type:text not null;max:240"`
	AnimalID uint
}

func (b *BehaviorNote) ValidateAndFormat() error {
	max_length := 240
	b.Note = strings.TrimSpace(b.Note)
	if b.Note == "" || len(b.Note) > max_length {
		return gorm.ErrInvalidData
	}
	return nil
}
