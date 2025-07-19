package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	secret string
	passwd string
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		secret: cfg.JWTSecret,
		passwd: cfg.Password,
	}
}

func (as *AuthService) Authenticate(p *models.Password) (string, error) {
	if p.Password == "" || p.Password != as.passwd {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"hash": hashPassword(as.passwd),
		"exp":  time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(as.secret))
}

func (as *AuthService) ValidateToken(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.secret), nil
	})
	if err != nil || !token.Valid {
		return false, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("invalid claims")
	}

	return claims["hash"] == hashPassword(as.passwd), nil
}

func hashPassword(p string) string {
	hash := sha256.Sum256([]byte(p))
	return hex.EncodeToString(hash[:])
}
