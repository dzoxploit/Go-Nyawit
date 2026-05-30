package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

func (s *Server) GetEstateIdDronePlan(
	ctx echo.Context,
	id types.UUID,
) error {

	_, width, length, err := s.Repository.GetEstateByID(
		ctx.Request().Context(),
		id.String(),
	)

	if err != nil {
		return ctx.JSON(
			http.StatusNotFound,
			generated.ErrorResponse{
				Message: StringPtr("estate not found"),
			},
		)
	}

	trees, err := s.Repository.GetEstateTrees(
		ctx.Request().Context(),
		id.String(),
	)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			generated.ErrorResponse{
				Message: StringPtr(err.Error()),
			},
		)
	}

	grid := make(map[[2]int]int)

	for _, tree := range trees {
		grid[[2]int{tree.X, tree.Y}] = tree.Height
	}

	var heights []int

	for y := 1; y <= width; y++ {

		if y%2 == 1 {

			for x := 1; x <= length; x++ {
				heights = append(
					heights,
					1+grid[[2]int{x, y}],
				)
			}

		} else {

			for x := length; x >= 1; x-- {
				heights = append(
					heights,
					1+grid[[2]int{x, y}],
				)
			}
		}
	}

	distance := 0

	if len(heights) > 0 {

		// naik dari tanah ke plot pertama
		distance += heights[0]

		// perpindahan antar plot
		for i := 1; i < len(heights); i++ {

			distance += 10

			diff := heights[i] - heights[i-1]
			if diff < 0 {
				diff = -diff
			}

			distance += diff
		}

		// turun ke tanah dari plot terakhir
		distance += heights[len(heights)-1]
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]any{
			"distance": distance,
		},
	)
}
