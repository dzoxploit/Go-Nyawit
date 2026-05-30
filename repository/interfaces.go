package repository

import "context"

type Tree struct {
	X      int
	Y      int
	Height int
}

type RepositoryInterface interface {
	CreateEstate(
		ctx context.Context,
		width int,
		length int,
	) (string, error)

	GetEstateByID(
		ctx context.Context,
		id string,
	) (bool, int, int, error)

	CreateTree(
		ctx context.Context,
		estateID string,
		x int,
		y int,
		height int,
	) (string, error)

	GetTreeHeights(
		ctx context.Context,
		estateID string,
	) ([]int, error)

	GetEstateTrees(
		ctx context.Context,
		estateID string,
	) ([]Tree, error)
}
