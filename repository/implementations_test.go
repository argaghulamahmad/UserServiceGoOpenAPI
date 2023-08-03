package repository

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockRepositoryInterface(ctrl)

	testPhone := "1234567890"
	expectedOutput := GetUserOutput{
		FullName: "John Doe",
		Phone:    testPhone,
	}

	mockDB.EXPECT().GetUser(context.Background(), GetUserInput{Phone: testPhone}).Return(expectedOutput, nil)

	output, err := mockDB.GetUser(context.Background(), GetUserInput{Phone: testPhone})

	// Check the results
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if output.FullName != expectedOutput.FullName || output.Phone != expectedOutput.Phone {
		t.Errorf("Expected output %v, but got %v", expectedOutput, output)
	}
}

// Add similar test functions for other repository functions

func TestInsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockRepositoryInterface(ctrl)

	testInput := InsertUserInput{
		FullName: "Jane Smith",
		Phone:    "9876543210",
		Password: "password123",
	}

	mockDB.EXPECT().InsertUser(context.Background(), testInput).Return(InsertUserOutput{}, nil)

	_, err := mockDB.InsertUser(context.Background(), testInput)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockRepositoryInterface(ctrl)

	testInput := UpdateUserInput{
		FullName: "Updated Name",
		Phone:    "1234567890",
	}

	mockDB.EXPECT().UpdateUser(context.Background(), testInput).Return(UpdateUserOutput{}, nil)

	_, err := mockDB.UpdateUser(context.Background(), testInput)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestIsPhonePasswordUserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockRepositoryInterface(ctrl)

	testInput := IsPhonePasswordUserExistInput{
		Phone:    "1234567890",
		Password: "password123",
	}

	expectedOutput := IsPhonePasswordUserExistOutput{
		FullName: "John Doe",
		Phone:    testInput.Phone,
	}

	mockDB.EXPECT().IsPhonePasswordUserExist(context.Background(), testInput).Return(expectedOutput, nil)

	output, err := mockDB.IsPhonePasswordUserExist(context.Background(), testInput)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if output.FullName != expectedOutput.FullName || output.Phone != expectedOutput.Phone {
		t.Errorf("Expected output %v, but got %v", expectedOutput, output)
	}
}
