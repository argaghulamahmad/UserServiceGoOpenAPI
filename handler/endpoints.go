package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/repository"
	"net/http"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var JWTSecretKey = []byte("argaghulamahmad-secretkey")

func (s *Server) LoginUser(ctx echo.Context) error {
	var params generated.LoginRequest
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if params.Phone == "" || params.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone number and password are required")
	}

	output, err := s.Repository.IsPhonePasswordUserExist(context.Background(), repository.IsPhonePasswordUserExistInput{
		Phone:    params.Phone,
		Password: params.Password,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to check user")
	}

	token, err := generateJWTToken(params.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": output.Id,
		"token":  token,
	})
}

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

func (s *Server) GetUser(ctx echo.Context) error {
	authorizationValue := ctx.Request().Header.Get("Authorization")
	authorizationValue = strings.Replace(authorizationValue, "Bearer ", "", 1)

	phoneNumber, err := getPhoneNumberFromJWTToken(authorizationValue)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid authorization token")
	}

	output, err := s.Repository.GetUser(context.Background(), repository.GetUserInput{
		Phone: phoneNumber,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	return ctx.JSON(http.StatusOK, generated.ProfileResponse{
		FullName: output.FullName,
		Phone:    phoneNumber,
	})
}

func (s *Server) UpdateUser(ctx echo.Context) error {
	var params = generated.UpdateUserRequest{
		FullName: nil,
		Phone:    nil,
	}
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) RegisterUser(ctx echo.Context) error {
	var params = generated.RegisterRequest{
		FullName: "",
		Phone:    "",
		Password: "",
	}
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	return ctx.JSON(http.StatusCreated, nil)
}
