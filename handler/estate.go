package handler

import (
	"context"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

func (s *Server) PostEstate(ctx echo.Context) error {

	var req generated.CreateEstateRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid request"),
			},
		)
	}

	if req.Width < 1 || req.Width > 50000 {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid width"),
			},
		)
	}

	if req.Length < 1 || req.Length > 50000 {
		return ctx.JSON(
			http.StatusBadRequest,
			generated.ErrorResponse{
				Message: StringPtr("invalid length"),
			},
		)
	}

	id, err := s.Repository.CreateEstate(
		ctx.Request().Context(),
		req.Width,
		req.Length,
	)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			generated.ErrorResponse{
				Message: StringPtr(err.Error()),
			},
		)
	}

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			generated.ErrorResponse{
				Message: StringPtr("failed to parse uuid"),
			},
		)
	}

	responseUUID := types.UUID(parsedUUID)

	return ctx.JSON(
		http.StatusCreated,
		generated.EstateResponse{
			Id: &responseUUID,
		},
	)
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
