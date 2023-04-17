package services

import (
	repository "ValREST/internal/database"
	"ValREST/internal/models"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "LsdmsdKDje984HFe83d3dhsd4858fjOSdj"
	tokenTTL  = 24 * time.Hour
	signInKey = "ijASDJIXhiuxspCSIDjxuswSAXJ"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(name, password string) (string, error) {
	user, err := s.repo.GetUser(name, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signInKey))
}
