package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var JWTSecretKey = []byte("argaghulamahmad-secretkey")

func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) LoginUser(ctx echo.Context) error {
	var params generated.LoginRequest
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if params.Phone == "" || params.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone number and password are required")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": "userId",
		"token":  "token",
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

func (s *Server) GetProfile(ctx echo.Context) error {
	_ = ctx.Get("user_id").(int)

	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) UpdateProfile(ctx echo.Context) error {
	var params = generated.UpdateProfileRequest{
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
