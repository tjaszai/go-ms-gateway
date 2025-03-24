package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tjaszai/go-ms-gateway/config"
	"github.com/tjaszai/go-ms-gateway/internal/contract"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"strings"
	"time"
)

type SecurityService struct {
	Repository *repository.UserRepository
	JwtSecret  string
}

func NewSecurityService(r *repository.UserRepository) *SecurityService {
	return &SecurityService{
		Repository: r,
		JwtSecret:  config.Config("JWT_SECRET"),
	}
}

func (s *SecurityService) Auth(c *fiber.Ctx) *contract.Error {
	err := contract.Error{Code: fiber.StatusUnauthorized, Message: "Unauthorized"}
	h := c.Get("Authorization")
	if h == "" {
		err.AddDetail("Authorization Header", "authorization header is required")
		return &err
	}
	var t string
	if strings.HasPrefix(h, "Bearer ") {
		t = strings.TrimPrefix(h, "Bearer ")
	}
	if t == "" {
		err.AddDetail("Authorization Header", "invalid Authorization Header")
		return &err
	}
	claims, tErr := s.DecodeToken(t)
	if tErr != nil {
		err.AddDetail("JWT Token", tErr)
		return &err
	}
	uID, _ := claims.GetSubject()
	u, uErr := s.fetchUser(uID)
	if uErr != nil {
		err.AddDetail("Invalid User", uErr)
		return &err
	}
	c.Locals("user", u)
	return nil
}

func (s *SecurityService) DecodeToken(token string) (*jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.JwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("invalid Token")
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("failed to verify token")
	}
	return &claims, nil
}

func (s *SecurityService) fetchUser(uID string) (*model.User, error) {
	u, err := s.Repository.Find(uID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if u.ID.String() != uID {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (s *SecurityService) GenerateToken(u *model.User) (*string, error) {
	if s.JwtSecret == "" {
		return nil, errors.New("JWT_SECRET is empty")
	}
	jT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud":   "ms-gateway",
		"sub":   u.ID,
		"roles": u.StrRoles(),
		"exp":   time.Now().Add(time.Hour * 12).Unix(),
	})
	t, err := jT.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *SecurityService) AdminGuard(c *fiber.Ctx) *contract.Error {
	u := c.Locals("user").(*model.User)
	if u == nil {
		err := contract.Error{Code: fiber.StatusUnauthorized, Message: "Unauthorized"}
		err.AddDetail("Unauthorized", "unauthorized")
		return &err
	}
	if !u.IsAdmin() {
		err := contract.Error{Code: fiber.StatusForbidden, Message: "Access Denied"}
		err.AddDetail("Permission Error", "you do not have the necessary permissions to access this resource")
		return &err
	}
	return nil
}
