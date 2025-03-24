package repository

import (
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/model"
)

type MicroserviceRepository struct {
	DatabaseManager *db.DatabaseManager
}

func NewMicroserviceRepository(m *db.DatabaseManager) *MicroserviceRepository {
	return &MicroserviceRepository{DatabaseManager: m}
}

func (r *MicroserviceRepository) FindAll() ([]model.Microservice, error) {
	var m []model.Microservice
	err := r.DatabaseManager.GetDB().Find(&m).Error
	return m, err
}

func (r *MicroserviceRepository) Find(id string) (*model.Microservice, error) {
	var m model.Microservice
	err := r.DatabaseManager.GetDB().Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *MicroserviceRepository) FindByName(name string) (*model.Microservice, error) {
	var m model.Microservice
	err := r.DatabaseManager.GetDB().Where("name = ?", name).First(&m).Error
	return &m, err
}

func (r *MicroserviceRepository) CreateFrom(d *dto.MsInputDto) (*model.Microservice, error) {
	m := d.ToModel(nil)
	err := r.DatabaseManager.GetDB().Create(&m).Error
	return m, err
}

func (r *MicroserviceRepository) UpdateFrom(m *model.Microservice, d *dto.MsInputDto) (*model.Microservice, error) {
	m = d.ToModel(m)
	err := r.DatabaseManager.GetDB().Save(m).Error
	return m, err
}

func (r *MicroserviceRepository) Delete(id string) error {
	return r.DatabaseManager.GetDB().Delete(&model.Microservice{}, "id = ?", id).Error
}
