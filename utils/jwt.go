package utils

import (
	"time"

	configs "com.lopster-pos/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(userID string, phone string) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"sub":   userID,
		"phone": phone,
		"type":  "access",
		"exp":   time.Now().Add(time.Minute * time.Duration(configs.AccessTokenTTL)).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(configs.JWTSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"sub":   userID,
		"phone": phone,
		"type":  "refresh",
		"exp":   time.Now().Add(time.Hour * 24 * time.Duration(configs.RefreshTokenTTL)).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(configs.JWTSecret)

	return accessString, refreshString, err
}

func ValidateAccessToken(tokenStr string) (string, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return configs.JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["type"] != "access" {
		return "", "", jwt.ErrInvalidType
	}

	userID, _ := claims["sub"].(string)
	phone, _ := claims["phone"].(string)

	return userID, phone, nil
}

func ValidateRefreshToken(tokenStr string) (string, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return configs.JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["type"] != "refresh" {
		return "", "", jwt.ErrHashUnavailable
	}

	userID, _ := claims["sub"].(string)
	phone, _ := claims["phone"].(string)

	return userID, phone, nil
}
