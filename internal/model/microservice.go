package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Microservice struct {
	gorm.Model
	ID             uuid.UUID             `gorm:"type:uuid"`
	Name           string                `gorm:"type:varchar(100);not null;uniqueIndex:ms_name_uniq_idx"`
	Description    sql.NullString        `gorm:"type:text"`
	Versions       []MicroserviceVersion `gorm:"foreignKey:MicroserviceID"`
	CurrentVersion *MicroserviceVersion  `gorm:"-"`
}

func (m *Microservice) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}

type MicroserviceVersion struct {
	gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid"`
	MicroserviceID uuid.UUID      `gorm:"type:uuid;not null;index"`
	Microservice   Microservice   `gorm:"foreignKey:MicroserviceID;constraint:OnDelete:CASCADE"`
	Name           string         `gorm:"type:varchar(30);not null;uniqueIndex:ms_version_uniq_idx"`
	Description    sql.NullString `gorm:"type:text"`
	Url            string         `gorm:"type:varchar;not null"`
	OpenAPIUrl     string         `gorm:"type:varchar;not null"`
}

func (mv *MicroserviceVersion) BeforeCreate(tx *gorm.DB) error {
	mv.ID = uuid.New()
	return nil
}
