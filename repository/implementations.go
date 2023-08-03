package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error) {
	query := "SELECT *, phone FROM users WHERE phone = $1"
	err = r.Db.QueryRowContext(ctx, query, input.Phone).Scan(&output.FullName, &output.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, fmt.Errorf("no profile found for phone number: %s", input.Phone)
		}
		return output, err
	}
	return output, nil
}

func (r *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (output UpdateUserOutput, err error) {
	query := "UPDATE users SET fullname = $1 WHERE phone = $2"
	_, err = r.Db.ExecContext(ctx, query, input.FullName, input.Phone)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (r *Repository) InsertUser(ctx context.Context, input InsertUserInput) (output InsertUserOutput, err error) {
	query := "INSERT INTO users (fullname, phone, password) VALUES ($1, $2, $3)"
	_, err = r.Db.ExecContext(ctx, query, input.FullName, input.Phone)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (r *Repository) IsPhonePasswordUserExist(ctx context.Context, input IsPhonePasswordUserExistInput) (output IsPhonePasswordUserExistOutput, err error) {
	query := "SELECT id, fullname, phone FROM users WHERE phone = $1 AND password = $2"
	err = r.Db.QueryRowContext(ctx, query, input.Phone, input.Password).Scan(&output.Id, &output.FullName, &output.Phone)
	if err != nil {
		output.IsExist = false
		if err == sql.ErrNoRows {
			return output, fmt.Errorf("no profile found for phone number: %s", input.Phone)
		}
		return output, err
	}
	output.IsExist = true
	return output, nil
}
