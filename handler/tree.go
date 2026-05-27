package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

func (s *Server) PostEstateIdTree(
	ctx echo.Context,
	id types.UUID,
) error {

	var req generated.CreateTreeRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid request"),
			},
		)
	}

	if req.Height < 1 || req.Height > 30 {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid height"),
			},
		)
	}

	exists, width, length, err := s.Repository.GetEstateByID(
		ctx.Request().Context(),
		id.String(),
	)

	if err != nil || !exists {
		return ctx.JSON(
			http.StatusNotFound,
			generated.ErrorResponse{
				Message: StringPtr("estate not found"),
			},
		)
	}

	if req.X < 1 || req.X > length {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid x"),
			},
		)
	}

	if req.Y < 1 || req.Y > width {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid y"),
			},
		)
	}

	treeID, err := s.Repository.CreateTree(
		ctx.Request().Context(),
		id.String(),
		req.X,
		req.Y,
		req.Height,
	)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			generated.ErrorResponse{
				Message: StringPtr(err.Error()),
			},
		)
	}

	parsedUUID, err := uuid.Parse(treeID)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			generated.ErrorResponse{
				Message: StringPtr("failed parse uuid"),
			},
		)
	}

	responseUUID := types.UUID(parsedUUID)

	return ctx.JSON(
		http.StatusCreated,
		generated.TreeResponse{
			Id: &responseUUID,
		},
	)
}
