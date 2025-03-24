package repository

import (
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
	"strings"
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
	err := r.DatabaseManager.GetDB().Where("email = ?", strings.ToLower(email)).First(&m).Error
	return &m, err
}

func (r *UserRepository) CreateFrom(d *dto.UserInputDto) (*model.User, error) {
	m := d.ToModel(nil)
	pwd, err := util.GenerateUserPwdHash(m.Password)
	if err != nil {
		return nil, err
	}
	m.Password = *pwd
	err = r.DatabaseManager.GetDB().Create(&m).Error
	return m, err
}

func (r *UserRepository) UpdateFrom(m *model.User, d *dto.UserInputDto) (*model.User, error) {
	m = d.ToModel(m)
	pwd, err := util.GenerateUserPwdHash(m.Password)
	if err != nil {
		return nil, err
	}
	m.Password = *pwd
	err = r.DatabaseManager.GetDB().Save(m).Error
	return m, err
}

func (r *UserRepository) Delete(id string) error {
	return r.DatabaseManager.GetDB().Delete(&model.User{}, "id = ?", id).Error
}
