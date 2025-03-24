package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Microservice struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid"`
	Name        string         `gorm:"uniqueIndex:ms_name_uniq_idx"`
	Description sql.NullString `gorm:"type:text"`
}

func (m *Microservice) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}

type MicroserviceVersion struct {
	gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid"`
	MicroserviceID uuid.UUID      `gorm:"type:uuid;uniqueIndex:ms_version_uniq_idx"`
	Microservice   Microservice   `gorm:"foreignKey:MicroserviceID"`
	Name           string         `gorm:"uniqueIndex:ms_version_uniq_idx"`
	Description    sql.NullString `gorm:"type:text"`
	Url            string
	OpenAPIUrl     string
}

func (mv *MicroserviceVersion) BeforeCreate(tx *gorm.DB) error {
	mv.ID = uuid.New()
	return nil
}
