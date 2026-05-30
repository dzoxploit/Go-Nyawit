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

func (r *Repository) GetEstateTrees(
	ctx context.Context,
	estateID string,
) ([]Tree, error) {

	query := `
		SELECT x, y, height
		FROM trees
		WHERE estate_id = $1
	`

	rows, err := r.DB.QueryContext(
		ctx,
		query,
		estateID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trees []Tree

	for rows.Next() {

		var tree Tree

		err := rows.Scan(
			&tree.X,
			&tree.Y,
			&tree.Height,
		)

		if err != nil {
			return nil, err
		}

		trees = append(trees, tree)
	}

	return trees, nil
}
