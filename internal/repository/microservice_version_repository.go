package repository

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/model"
)

type MicroserviceVersionRepository struct {
	DatabaseManager *db.DatabaseManager
}

func NewMicroserviceVersionRepository(m *db.DatabaseManager) *MicroserviceVersionRepository {
	return &MicroserviceVersionRepository{DatabaseManager: m}
}

func (r *MicroserviceVersionRepository) FindAll() ([]model.MicroserviceVersion, error) {
	var m []model.MicroserviceVersion
	err := r.DatabaseManager.GetDB().Find(&m).Error
	return m, err
}

func (r *MicroserviceVersionRepository) Find(id string) (*model.MicroserviceVersion, error) {
	var m model.MicroserviceVersion
	err := r.DatabaseManager.GetDB().Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *MicroserviceVersionRepository) FindByName(name string) (*model.MicroserviceVersion, error) {
	var m model.MicroserviceVersion
	err := r.DatabaseManager.GetDB().Where("name = ?", name).First(&m).Error
	return &m, err
}

func (r *MicroserviceVersionRepository) CreateFrom(msID string, d *dto.MsVersionInputDto) (*model.MicroserviceVersion, error) {
	m := d.ToModel(nil)
	m.MicroserviceID = uuid.MustParse(msID)
	err := r.DatabaseManager.GetDB().Create(&m).Error
	return m, err
}

func (r *MicroserviceVersionRepository) UpdateFrom(m *model.MicroserviceVersion, d *dto.MsVersionInputDto) (*model.MicroserviceVersion, error) {
	m = d.ToModel(m)
	err := r.DatabaseManager.GetDB().Save(m).Error
	return m, err
}

func (r *MicroserviceVersionRepository) Delete(id string) error {
	return r.DatabaseManager.GetDB().Delete(&model.MicroserviceVersion{}, "id = ?", id).Error
}
