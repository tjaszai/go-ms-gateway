package repository

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/model"
)

type UserRepository struct {
	DatabaseManager *db.DatabaseManager
}

func NewUserRepository(m *db.DatabaseManager) *UserRepository {
	return &UserRepository{DatabaseManager: m}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var m []model.User
	err := r.DatabaseManager.GetDB().Find(&m).Error
	return m, err
}

func (r *UserRepository) Find(id string) (*model.User, error) {
	var m model.User
	err := r.DatabaseManager.GetDB().Where("id = ?", id).First(&m).Error
	return &m, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var m model.User
	err := r.DatabaseManager.GetDB().Where("email = ?", email).First(&m).Error
	return &m, err
}

func (r *UserRepository) CreateFromReqDto(d *dto.CreateUserReqDto) (*model.User, error) {
	m := d.UserReqDtoToModel()
	if err := m.HashPassword(); err != nil {
		return nil, err
	}
	m.ID = uuid.New()
	err := r.DatabaseManager.GetDB().Create(&m).Error
	return m, err
}

func (r *UserRepository) Update(d *dto.UpdateUserReqDto, m *model.User) error {
	var err error
	m = d.UserReqDtoToModel(m)
	if d.Password != nil {
		err = m.HashPassword()
		if err != nil {
			return err
		}
	}
	return r.DatabaseManager.GetDB().Save(m).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.DatabaseManager.GetDB().Delete(&model.User{}, "id = ?", id).Error
}
