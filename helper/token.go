package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateTokens(username string) (string, string, error) {
	accessTokenExpiration := time.Now().Add(5 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiration),
		},
	}

	refreshTokenExpiration := time.Now().Add(5 * time.Minute)
	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiration),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	claims := Claims{}

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("refresh token invalid")
	}

	return generateNewAccessToken(claims.Username)
}

func generateNewAccessToken(username string) (string, error) {
	accessTokenExpiration := time.Now().Add(5 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiration),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
