package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tjaszai/go-ms-gateway/internal/contract/enum"
	"gorm.io/gorm"
	"slices"
)

type User struct {
	gorm.Model
	ID       uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Name     string          `gorm:"type:varchar(50);uniqueIndex:user_name_uniq_idx;not null"`
	Email    string          `gorm:"type:varchar(80);uniqueIndex:user_email_uniq_idx;not null"`
	Password string          `gorm:"type:varchar(100);not null"`
	Roles    pq.Int64Array   `gorm:"type:integer[];not null"`
	Status   enum.UserStatus `gorm:"type:integer;not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.Status = enum.UserStatusInactive
	u.Roles = pq.Int64Array{int64(enum.UserRoleUser)}
	return nil
}

func (u *User) StrRoles() []string {
	var roles = make([]string, len(u.Roles))
	for i, role := range u.Roles {
		roles[i] = enum.UserRole(role).String()
	}
	return roles
}

func (u *User) IsAdmin() bool {
	return slices.ContainsFunc(u.Roles, func(role int64) bool {
		return enum.UserRole(role).IsAdmin()
	})
}

func (u *User) StrStatus() string {
	return u.Status.String()
}
