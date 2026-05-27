package repository

import "context"

func (r *Repository) CreateTree(
	ctx context.Context,
	estateID string,
	x int,
	y int,
	height int,
) (string, error) {

	var id string

	query := `
		INSERT INTO trees(
			estate_id,
			x,
			y,
			height
		)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`

	err := r.DB.QueryRowContext(
		ctx,
		query,
		estateID,
		x,
		y,
		height,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
