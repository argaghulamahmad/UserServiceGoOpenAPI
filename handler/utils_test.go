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

func TestValidatePassword(t *testing.T) {

	testCases := []struct {
		password    string
		expectedRes bool
	}{
		{"Abcdef123!", true},
		{"abcdef123!", false},
		{"ABCDEF123!", false},
		{"Abcdefghi", false},
		{"Abcdef1234567890!", true},
	}

	for _, tc := range testCases {
		actualRes := validatePassword(tc.password)
		if actualRes != tc.expectedRes {
			t.Errorf("Expected validatePassword(%q) to be %v, but got %v", tc.password, tc.expectedRes, actualRes)
		}
	}
}

func TestValidatePhoneNumber(t *testing.T) {

	testCases := []struct {
		phoneNumber string
		expectedRes bool
	}{
		{"+6281234567890", true},
		{"+62812345678", false},
		{"+62812345678901234", false},
		{"081234567890", false},
		{"+621234567890", true},
		{"+628abcdefghij", false},
	}

	for _, tc := range testCases {
		actualRes := validatePhoneNumber(tc.phoneNumber)
		if actualRes != tc.expectedRes {
			t.Errorf("Expected validatePhoneNumber(%q) to be %v, but got %v", tc.phoneNumber, tc.expectedRes, actualRes)
		}
	}
}

func TestValidateFullName(t *testing.T) {

	testCases := []struct {
		fullName    string
		expectedRes bool
	}{
		{"John Doe", true},
		{"A", false},
		{"Lorem Ipsum Dolor Sit Amet Consectetur Adipiscing Elit Sed Do Eiusmod Tempor Incididunt", false},
	}

	for _, tc := range testCases {
		actualRes := validateFullName(tc.fullName)
		if actualRes != tc.expectedRes {
			t.Errorf("Expected validateFullName(%q) to be %v, but got %v", tc.fullName, tc.expectedRes, actualRes)
		}
	}
}
