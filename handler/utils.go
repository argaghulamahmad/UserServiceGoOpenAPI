package handler

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func generateJWTToken(phoneNumber string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = phoneNumber
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getPhoneNumberFromJWTToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if phoneNumber, ok := claims["sub"].(string); ok {
			return phoneNumber, nil
		}
		return "", errors.New("Phone number not found in token claims")
	}

	return "", errors.New("Invalid token")
}
