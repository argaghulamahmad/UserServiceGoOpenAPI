package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginUser(t *testing.T) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := repository.NewMockRepositoryInterface(ctrl)

	testInput := repository.IsPhonePasswordUserExistInput{
		Phone:    "1234567890",
		Password: "password123",
	}

	expectedOutput := repository.IsPhonePasswordUserExistOutput{
		FullName: "John Doe",
		Phone:    testInput.Phone,
	}

	mockDB.EXPECT().IsPhonePasswordUserExist(context.Background(), testInput).Return(expectedOutput, nil)

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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := repository.NewMockRepositoryInterface(ctrl)

	testPhone := "1234567890"
	expectedOutput := repository.GetUserOutput{
		FullName: "John Doe",
		Phone:    testPhone,
	}

	mockDB.EXPECT().GetUser(context.Background(), repository.GetUserInput{Phone: testPhone}).Return(expectedOutput, nil)

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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := repository.NewMockRepositoryInterface(ctrl)

	testInput := repository.UpdateUserInput{
		FullName: "Updated Name",
		Phone:    "1234567890",
	}

	mockDB.EXPECT().UpdateUser(context.Background(), testInput).Return(repository.UpdateUserOutput{}, nil)

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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := repository.NewMockRepositoryInterface(ctrl)

	testInput := repository.InsertUserInput{
		FullName: "Jane Smith",
		Phone:    "9876543210",
		Password: "password123",
	}

	mockDB.EXPECT().InsertUser(context.Background(), testInput).Return(repository.InsertUserOutput{}, nil)

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
