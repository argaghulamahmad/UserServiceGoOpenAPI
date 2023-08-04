package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (output GetUserOutput, err error) {
	query := "SELECT id, fullname, phone, password FROM users WHERE phone = $1"
	err = r.Db.QueryRowContext(ctx, query, phone).Scan(&output.ID, &output.FullName, &output.Phone, &output.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, fmt.Errorf("no profile found for phone number: %s", phone)
		}
		return output, err
	}
	return output, nil
}
func (r *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (output UpdateUserOutput, err error) {
	query := "UPDATE users SET fullname = $1 WHERE phone = $2"
	_, err = r.Db.ExecContext(ctx, query, input.FullName, input.Phone)
	if err != nil {
		return output, fmt.Errorf("failed to update user: %w", err)
	}

	output.FullName = input.FullName
	output.Phone = input.Phone
	return output, nil
}

func (r *Repository) InsertUser(ctx context.Context, input InsertUserInput) (output InsertUserOutput, err error) {
	query := "INSERT INTO users (fullname, phone, password) VALUES ($1, $2, $3) RETURNING id"
	err = r.Db.QueryRowContext(ctx, query, input.FullName, input.Phone, input.Password).Scan(&output.ID)
	if err != nil {
		return output, err
	}

	return output, nil
}
