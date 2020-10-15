package middlewares

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/config"
)

func GenerateToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte(config.Cfg.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func RefreshToken(token string) (string, error) {
	claims, isValid := extractClaims(token)

	if !isValid {
		return "", fmt.Errorf("Wasn't able to refresh token")
	}

	idInterface := claims["id"]
	id, ok := idInterface.(uint)

	if !ok {
		return "", fmt.Errorf("Wasn't able to refresh token")
	}

	return GenerateToken(id)
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	secret := []byte(config.Cfg.Secret)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
