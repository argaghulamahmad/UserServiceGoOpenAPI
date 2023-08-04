package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestNewRepository(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(NewRepositoryOptions{
		Dsn: "test-dsn",
	})
	if repo == nil {
		t.Fatalf("Repository is nil")
	}

	if repo.Db == nil {
		t.Fatalf("Db field is not set")
	}
}
