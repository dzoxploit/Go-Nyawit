package repository

import "context"

func (r *Repository) GetTreeHeights(
	ctx context.Context,
	estateID string,
) ([]int, error) {

	rows, err := r.DB.QueryContext(
		ctx,
		`
		SELECT height
		FROM trees
		WHERE estate_id = $1
		ORDER BY height
		`,
		estateID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var heights []int

	for rows.Next() {
		var h int

		if err := rows.Scan(&h); err != nil {
			return nil, err
		}

		heights = append(heights, h)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return heights, nil
}
