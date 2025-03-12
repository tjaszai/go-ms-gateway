package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Microservice struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	// TODO: fix unique index -> soft delete
	Name        string `gorm:"uniqueIndex:ms_name_unique_idx"`
	Description string `gorm:"type:text,omitempty"`
}
