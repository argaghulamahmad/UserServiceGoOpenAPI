package handler

import (
	"testing"
)

func TestGenerateAndParseJWTToken(t *testing.T) {
	generateToken := func(phoneNumber string) string {
		tokenString, err := generateJWTToken(phoneNumber)
		if err != nil {
			t.Fatalf("Error generating JWT token: %v", err)
		}
		return tokenString
	}

	parseTokenAndGetPhoneNumber := func(tokenString string) string {
		phoneNumber, err := getPhoneNumberFromJWTToken(tokenString)
		if err != nil {
			t.Fatalf("Error getting phone number from JWT token: %v", err)
		}
		return phoneNumber
	}

	t.Run("ValidToken", func(t *testing.T) {
		phoneNumber := "1234567890"
		tokenString := generateToken(phoneNumber)
		retrievedPhoneNumber := parseTokenAndGetPhoneNumber(tokenString)

		if retrievedPhoneNumber != phoneNumber {
			t.Errorf("Expected phone number %s, but got %s", phoneNumber, retrievedPhoneNumber)
		}
	})

	t.Run("InvalidToken", func(t *testing.T) {
		invalidToken := "invalid_token"
		_, err := getPhoneNumberFromJWTToken(invalidToken)
		if err == nil {
			t.Errorf("Expected an error for an invalid token, but got none")
		} else if err.Error() != "token is malformed: token contains an invalid number of segments" {
			t.Errorf("Expected error message 'Invalid token', but got: %v", err)
		}
	})
}
