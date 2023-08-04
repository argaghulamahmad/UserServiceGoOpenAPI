package handler

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"regexp"
	"strings"
	"time"
)

var JWTSecretKey = []byte("argaghulamahmad-secretkey")

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

func validatePassword(password string) bool {
	if len(password) < 6 || len(password) > 64 {
		return false
	}

	hasUppercase := false
	hasLowercase := false
	hasNumber := false
	hasSpecial := false
	specialChars := "!@#$%^&*()_-+=[]{}|;:,.<>?"

	for _, char := range password {
		charStr := string(char)
		if strings.ContainsAny(charStr, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			hasUppercase = true
		} else if strings.ContainsAny(charStr, "abcdefghijklmnopqrstuvwxyz") {
			hasLowercase = true
		} else if strings.ContainsAny(charStr, "0123456789") {
			hasNumber = true
		} else if strings.ContainsAny(charStr, specialChars) {
			hasSpecial = true
		}
	}

	return hasLowercase && hasUppercase && hasNumber && hasSpecial
}

func validatePhoneNumber(phoneNumber string) bool {
	phonePattern := `^\+62\d{10,13}$`
	matched, _ := regexp.MatchString(phonePattern, phoneNumber)
	return matched
}

func validateFullName(fullName string) bool {
	return len(fullName) >= 3 && len(fullName) <= 60
}
