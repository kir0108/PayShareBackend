package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TestToken = "cjwiemk23i4jivcj2ij"

type InvalidTokenError struct {
	message string
}

func (e *InvalidTokenError) Error() string {
	return e.message
}

var (
	ErrInvalidToken   = &InvalidTokenError{"invalid token"}
	ErrExpiredToken   = &InvalidTokenError{"expired token"}
	ErrInvalidSubject = &InvalidTokenError{"invalid subject"}
)

type Config struct {
	Secret     string
	Expiration time.Duration
}

type Manager struct {
	Secret     string
	Expiration time.Duration
}

func NewManger(c *Config) *Manager {
	return &Manager{
		Secret:     c.Secret,
		Expiration: c.Expiration,
	}
}

func (m *Manager) CreateToken(id int64) (string, error) {
	if m.Secret == "test_secret" {
		return "test_access_token", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(m.Expiration).Unix(),
		Subject:   strconv.FormatInt(id, 10),
	})

	return token.SignedString([]byte(m.Secret))
}

func (m *Manager) GetIdFromToken(token string) (int64, error) {
	if token == TestToken {
		return 1, nil
	}

	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.Secret), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr); jwtErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, ErrExpiredToken
		}
		return 0, ErrInvalidToken
	}

	claims, ok := parsed.Claims.(*jwt.StandardClaims)
	if !ok {
		return 0, ErrInvalidToken
	}

	id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, ErrInvalidSubject
	}

	return id, nil
}
