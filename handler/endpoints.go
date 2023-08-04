package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func (s *Server) LoginUser(ctx echo.Context) error {
	var params generated.LoginRequest
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if params.Phone == "" || params.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone number and password are required")
	}

	user, err := s.Repository.GetUserByPhone(context.Background(), params.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid phone number or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid phone number or password")
	}

	token, err := generateJWTToken(params.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": user.ID,
		"token":  token,
	})
}

func (s *Server) GetUser(ctx echo.Context) error {
	authorizationValue := ctx.Request().Header.Get("Authorization")
	authorizationValue = strings.Replace(authorizationValue, "Bearer ", "", 1)

	phoneNumber, err := getPhoneNumberFromJWTToken(authorizationValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid authorization token")
	}

	user, err := s.Repository.GetUserByPhone(context.Background(), phoneNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	return ctx.JSON(http.StatusOK, generated.ProfileResponse{
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

func (s *Server) UpdateUser(ctx echo.Context) error {
	authorizationValue := ctx.Request().Header.Get("Authorization")
	authorizationValue = strings.Replace(authorizationValue, "Bearer ", "", 1)

	phoneNumber, err := getPhoneNumberFromJWTToken(authorizationValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid authorization token")
	}

	var params generated.UpdateUserRequest
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	updatedFields := make(map[string]interface{})
	if params.FullName != nil {
		updatedFields["FullName"] = *params.FullName
	}
	if params.Phone != nil {
		updatedFields["Phone"] = *params.Phone
	}

	updateInput := repository.UpdateUserInput{
		FullName: *params.FullName,
		Phone:    phoneNumber,
	}

	user, err := s.Repository.UpdateUser(context.Background(), updateInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	return ctx.JSON(http.StatusOK, generated.ProfileResponse{
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

func (s *Server) RegisterUser(ctx echo.Context) error {
	var params = generated.RegisterRequest{}
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	var validationErrors []string

	if !validatePhoneNumber(params.Phone) {
		validationErrors = append(validationErrors, "Invalid phone number format")
	}

	if !validateFullName(params.FullName) {
		validationErrors = append(validationErrors, "Full name must be between 3 and 60 characters")
	}

	if !validatePassword(params.Password) {
		validationErrors = append(validationErrors, "Password must be between 6 and 64 characters and contain at least 1 uppercase letter, 1 number, and 1 special character")
	}

	if len(validationErrors) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, validationErrors)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	user, err := s.Repository.InsertUser(context.Background(), repository.InsertUserInput{
		FullName: params.FullName,
		Phone:    params.Phone,
		Password: string(hashedPassword),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register user")
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"userId": user.ID,
	})
}
