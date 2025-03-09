package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Microservice struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string    `gorm:"uniqueIndex:ms_name_unique_idx"`
	Description *string   `gorm:"type:text"`
}
