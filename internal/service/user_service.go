package service

import (
	"errors"
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo       *repository.UserRepository
	jwtSecret  string
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: "your-secret-key-change-in-production",
	}
}

func (s *UserService) Login(username, password string) (*models.LoginResponse, error) {
	userDTO, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDTO.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(userDTO.ID, userDTO.Username, userDTO.Role)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		Role:  userDTO.Role,
	}, nil
}

func (s *UserService) generateToken(userID int64, username, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *UserService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
