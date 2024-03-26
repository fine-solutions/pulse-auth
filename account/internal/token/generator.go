package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"pulse-auth/cmd/pulse/config"
	"pulse-auth/internal/model"
	"time"
)

const ExpirationDuration = time.Hour * 24

type Generator struct {
	saltValue  []byte
	serverName string
}

func NewGenerator(config config.ApplicationConfig) *Generator {
	return &Generator{
		saltValue:  []byte(config.SaltValue),
		serverName: config.App,
	}
}

func (g *Generator) GenerateToken(user *model.User) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if user == nil {
		return "", fmt.Errorf("user cannot be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": g.serverName,
		"sub": user.Username,
	})

	s, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("signed string: %w", err)
	}

	return s, nil
}

func (g *Generator) GetExpirationDate() time.Time {
	return time.Now().Add(ExpirationDuration)
}
