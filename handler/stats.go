package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

func (s *Server) GetEstateIdStats(
	ctx echo.Context,
	id types.UUID,
) error {

	heights, err := s.Repository.GetTreeHeights(
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

	count := len(heights)

	if count == 0 {
		return ctx.JSON(
			http.StatusOK,
			generated.StatsResponse{
				Count:  IntPtr(0),
				Max:    IntPtr(0),
				Min:    IntPtr(0),
				Median: IntPtr(0),
			},
		)
	}

	minHeight := heights[0]
	maxHeight := heights[count-1]

	median := 0

	if count%2 == 1 {
		median = heights[count/2]
	} else {
		median = (heights[count/2-1] + heights[count/2]) / 2
	}

	return ctx.JSON(
		http.StatusOK,
		generated.StatsResponse{
			Count:  IntPtr(count),
			Max:    IntPtr(maxHeight),
			Min:    IntPtr(minHeight),
			Median: IntPtr(median),
		},
	)
}
