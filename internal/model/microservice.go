package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: fix unique index -> soft delete

type Microservice struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid"`
	Name        string         `gorm:"uniqueIndex:ms_name_unique_idx"`
	Description sql.NullString `gorm:"type:text"`
}

type MicroserviceVersion struct {
	gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid"`
	MicroserviceID uuid.UUID      `gorm:"type:uuid;uniqueIndex:ms_version_unique_idx"`
	Microservice   Microservice   `gorm:"foreignKey:MicroserviceID"`
	Name           string         `gorm:"uniqueIndex:ms_version_unique_idx"`
	Description    sql.NullString `gorm:"type:text"`
	Url            string
	OpenAPIUrl     string
}
