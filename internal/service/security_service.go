package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/tjaszai/go-ms-gateway/config"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"time"
)

type SecurityService struct {
	JWTSecret string
}

func NewSecurityService() *SecurityService {
	return &SecurityService{
		JWTSecret: config.Config("JWT_SECRET", "NotAValidJwtSecret"),
	}
}

func (s *SecurityService) GenerateToken(u *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
