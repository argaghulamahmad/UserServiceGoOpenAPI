package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) GetProfile(ctx context.Context, input GetProfileInput) (output GetProfileOutput, err error) {
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

func (r *Repository) UpdateProfile(ctx context.Context, input UpdateProfileInput) (output UpdateProfileOutput, err error) {
	query := "UPDATE users SET fullname = $1 WHERE phone = $2"
	_, err = r.Db.ExecContext(ctx, query, input.FullName, input.Phone)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (r *Repository) InsertProfile(ctx context.Context, input InsertProfileInput) (output InsertProfileOutput, err error) {
	query := "INSERT INTO users (fullname, phone, password) VALUES ($1, $2, $3)"
	_, err = r.Db.ExecContext(ctx, query, input.FullName, input.Phone)
	if err != nil {
		return output, err
	}

	return output, nil
}
