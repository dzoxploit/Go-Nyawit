// repository/estate.go

package repository

import (
	"context"
)

func (r *Repository) CreateEstate(
	ctx context.Context,
	width int,
	length int,
) (string, error) {

	var id string

	query := `
		INSERT INTO estates(width, length)
		VALUES($1, $2)
		RETURNING id
	`

	err := r.DB.QueryRowContext(
		ctx,
		query,
		width,
		length,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *Repository) GetEstateByID(
	ctx context.Context,
	id string,
) (bool, int, int, error) {

	var width int
	var length int

	query := `
		SELECT width, length
		FROM estates
		WHERE id = $1
	`

	err := r.DB.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(&width, &length)

	if err != nil {
		return false, 0, 0, err
	}

	return true, width, length, nil
}
