package handler

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"Phone": "1234567890", "Password": "Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Password123@"), bcrypt.DefaultCost)

	mockUser := repository.GetUserOutput{
		ID:       1,
		FullName: "Test User",
		Phone:    "1234567890",
		Password: string(hashedPassword),
	}

	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), "1234567890").Return(mockUser, nil)

	err = server.LoginUser(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	jwtToken, _ := generateJWTToken("1234567890")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", jwtToken))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockUser := repository.GetUserOutput{
		ID:       1,
		FullName: "Test User",
		Phone:    "1234567890",
		Password: "$2a$10$abcdefgh...",
	}

	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), "1234567890").Return(mockUser, nil)

	err := server.GetUser(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"fullName":"Test User","phone":"1234567890"}` + "\n"
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	jwtToken, _ := generateJWTToken("1234567890")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/user", strings.NewReader(`{"FullName": "Updated User", "Phone": "+6212345678901"}`))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", jwtToken))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(repository.UpdateUserOutput{
		FullName: "Updated User",
		Phone:    "1234567890",
	}, nil)

	err := server.UpdateUser(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"fullName":"Updated User","phone":"1234567890"}` + "\n"
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"fullName": "Test User", "phone": "+6212345678901", "password": "Passw0rd!"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(repository.InsertUserOutput{
		ID: 1,
	}, nil)

	err := server.RegisterUser(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedResponse := `{"userId":1}` + "\n"
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestInvalidLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"Phone": "1234567890", "Password": "Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), "1234567890").Return(repository.GetUserOutput{}, errors.New("user not found"))

	err := server.LoginUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestInvalidUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/user", strings.NewReader(`{"FullName": "Updated User", "Phone": "1234567890"}`))
	req.Header.Set(echo.HeaderAuthorization, "Bearer test_token")
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(repository.UpdateUserOutput{}, errors.New("user update failed"))

	err := server.UpdateUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestInvalidRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"FullName": "Test User", "Phone": "1234567890", "Password": "Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(repository.InsertUserOutput{}, errors.New("user registration failed"))

	err := server.RegisterUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestInvalidGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer test_token")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), "1234567890").Return(repository.GetUserOutput{}, errors.New("user not found"))

	err := server.GetUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestInvalidPhoneNumberFormatRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"FullName": "Test User", "Phone": "1234567", "Password": "Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := server.RegisterUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid phone number format")
}

func TestInvalidFullNameLengthRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"FullName": "A", "Phone": "1234567890", "Password": "Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := server.RegisterUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Full name must be between 3 and 60 characters")
}

func TestInvalidPasswordFormatRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := NewServer(NewServerOptions{
		Repository: mockRepo,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"FullName": "Test User", "Phone": "1234567890", "Password": "password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := server.RegisterUser(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Password must be between 6 and 64 characters and contain at least 1 uppercase letter, 1 number, and 1 special character")
}
