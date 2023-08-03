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
