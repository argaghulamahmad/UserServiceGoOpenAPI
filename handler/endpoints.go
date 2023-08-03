package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
)

func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

type RegistrationParams struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
}

func (s *Server) Register(ctx echo.Context) error {
	var params RegistrationParams
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if !isValidPhoneNumber(params.PhoneNumber) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid phone number")
	}

	if !isValidFullName(params.FullName) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid full name")
	}

	if !isValidPassword(params.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid password")
	}

	_ = hashAndSaltPassword(params.Password)

	userID := 123

	return ctx.JSON(http.StatusOK, map[string]int{
		"userId": userID,
	})
}

func isValidPhoneNumber(phoneNumber string) bool {
	return len(phoneNumber) >= 10 && len(phoneNumber) <= 13 && regexp.MustCompile(`^\+62`).MatchString(phoneNumber)
}

func isValidFullName(fullName string) bool {
	return len(fullName) >= 3 && len(fullName) <= 60
}

func isValidPassword(password string) bool {
	hasCapitalLetter, _ := regexp.MatchString(`[A-Z]`, password)
	hasNumber, _ := regexp.MatchString(`[0-9]`, password)
	hasSpecialChar, _ := regexp.MatchString(`[^a-zA-Z0-9]`, password)

	return len(password) >= 6 && len(password) <= 64 && hasCapitalLetter && hasNumber && hasSpecialChar
}

func hashAndSaltPassword(password string) string {
	return password
}

type LoginParams struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (s *Server) Login(ctx echo.Context) error {
	var params LoginParams
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	userID := 123

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": userID,
		"token":  "sample-jwt-token",
	})
}

func (s *Server) GetProfile(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	if !isValidJWTToken(token) {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid or expired JWT token")
	}

	userName := "John Doe"
	phoneNumber := "+628123456789"

	return ctx.JSON(http.StatusOK, map[string]string{
		"name":  userName,
		"phone": phoneNumber,
	})
}

func isValidJWTToken(token string) bool {
	return true
}

type UpdateProfileParams struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
}

func (s *Server) UpdateProfile(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	if !isValidJWTToken(token) {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid or unauthorized JWT token")
	}

	var params UpdateProfileParams
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if params.PhoneNumber == "" && params.FullName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No fields to update")
	}

	if params.PhoneNumber != "" {
		if isPhoneNumberTaken(params.PhoneNumber) {
			return echo.NewHTTPError(http.StatusConflict, "Phone number already exists")
		}
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Profile updated successfully",
	})
}

func isPhoneNumberTaken(phoneNumber string) bool {
	return false
}
