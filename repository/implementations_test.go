package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func setupMockRepository(t *testing.T) (*gomock.Controller, *MockRepositoryInterface, *Repository) {
	mockCtrl := gomock.NewController(t)
	mockDb := NewMockRepositoryInterface(mockCtrl)
	repo := &Repository{Db: nil}
	return mockCtrl, mockDb, repo
}

func TestRepository(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*MockRepositoryInterface)
		assertion func(*testing.T, error)
	}{
		{
			name: "GetUserByPhone - Success",
			setup: func(mockDb *MockRepositoryInterface) {
				expectedOutput := GetUserOutput{
					ID:       1,
					FullName: "John Doe",
					Phone:    "1234567890",
					Password: "hashed_password",
				}
				phone := "1234567890"
				mockDb.EXPECT().GetUserByPhone(gomock.Any(), phone).Return(expectedOutput, nil)
			},
			assertion: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			},
		},
		{
			name: "GetUserByPhone - ErrorNotFound",
			setup: func(mockDb *MockRepositoryInterface) {
				phone := "non_existent_phone"
				mockDb.EXPECT().GetUserByPhone(gomock.Any(), phone).Return(GetUserOutput{}, sql.ErrNoRows)
			},
			assertion: func(t *testing.T, err error) {
				if !errors.Is(err, sql.ErrNoRows) {
					t.Errorf("Expected sql.ErrNoRows error, but got %v", err)
				}
			},
		},
		{
			name: "UpdateUser - Success",
			setup: func(mockDb *MockRepositoryInterface) {
				input := UpdateUserInput{
					FullName: "Jane Doe",
					Phone:    "9876543210",
				}
				mockDb.EXPECT().UpdateUser(gomock.Any(), input).Return(UpdateUserOutput{
					FullName: input.FullName,
					Phone:    input.Phone,
				}, nil)
			},
			assertion: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			},
		},
		{
			name: "InsertUser - Success",
			setup: func(mockDb *MockRepositoryInterface) {
				input := InsertUserInput{
					FullName: "Test User",
					Phone:    "5555555555",
					Password: "test_password",
				}
				mockDb.EXPECT().InsertUser(gomock.Any(), input).Return(InsertUserOutput{
					ID: 1,
				}, nil)
			},
			assertion: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl, mockDb, _ := setupMockRepository(t)
			defer mockCtrl.Finish()

			test.setup(mockDb)

			switch test.name {
			case "GetUserByPhone - Success":

				output, err := mockDb.GetUserByPhone(context.Background(), "1234567890")

				test.assertion(t, err)

				expectedOutput := GetUserOutput{
					ID:       1,
					FullName: "John Doe",
					Phone:    "1234567890",
					Password: "hashed_password",
				}
				if output != expectedOutput {
					t.Errorf("Expected output %+v, but got %+v", expectedOutput, output)
				}

			case "GetUserByPhone - ErrorNotFound":

				_, err := mockDb.GetUserByPhone(context.Background(), "non_existent_phone")

				test.assertion(t, err)

			case "UpdateUser - Success":

				output, err := mockDb.UpdateUser(context.Background(), UpdateUserInput{
					FullName: "Jane Doe",
					Phone:    "9876543210",
				})

				test.assertion(t, err)

				expectedOutput := UpdateUserOutput{
					FullName: "Jane Doe",
					Phone:    "9876543210",
				}
				if output != expectedOutput {
					t.Errorf("Expected output %+v, but got %+v", expectedOutput, output)
				}

			case "InsertUser - Success":

				output, err := mockDb.InsertUser(context.Background(), InsertUserInput{
					FullName: "Test User",
					Phone:    "5555555555",
					Password: "test_password",
				})

				test.assertion(t, err)

				expectedOutput := InsertUserOutput{
					ID: 1,
				}
				if output != expectedOutput {
					t.Errorf("Expected output %+v, but got %+v", expectedOutput, output)
				}
			}
		})
	}
}
