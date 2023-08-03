package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginUser(t *testing.T) {
	e := echo.New()

	loginData := generated.LoginRequest{
		Phone:    "testuser",
		Password: "testpassword",
	}
	body, _ := json.Marshal(loginData)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	server := &Server{}

	err := server.LoginUser(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Contains(t, response, "userId")
		assert.Contains(t, response, "token")
	}
}

func TestGetUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctx.Set("user_id", 123)

	server := &Server{}

	err := server.GetUser(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	e := echo.New()

	updateData := generated.UpdateUserRequest{
		FullName: nil,
		Phone:    nil,
	}
	body, _ := json.Marshal(updateData)

	req := httptest.NewRequest(http.MethodPut, "/profile", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	server := &Server{}

	err := server.UpdateUser(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRegisterUser(t *testing.T) {
	e := echo.New()

	registerData := generated.RegisterRequest{
		FullName: "John Doe",
		Phone:    "testuser",
		Password: "testpassword",
	}
	body, _ := json.Marshal(registerData)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	server := &Server{}

	err := server.RegisterUser(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
