package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) GetProfile(ctx context.Context, input GetProfileByPhoneNumberInput) (output GetProfileByPhoneNumberOutput, err error) {
	query := "SELECT * FROM users WHERE phone = $1"
	err = r.Db.QueryRowContext(ctx, query, input.PhoneNumber).Scan(&output.FullName, &output.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, fmt.Errorf("no profile found for phone number: %s", input.PhoneNumber)
		}
		return output, err
	}
	return output, nil
}
