package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestPostEstateIdTreeDuplicateCoordinate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	estateID := uuid.New()

	mockRepo.
		EXPECT().
		GetEstateByID(gomock.Any(), estateID.String()).
		Return(true, 10, 10, nil)

	mockRepo.
		EXPECT().
		CreateTree(gomock.Any(), estateID.String(), 2, 3, 5).
		Return("", repository.ErrTreeAlreadyExists)

	server := &Server{
		Repository: mockRepo,
	}

	body := []byte(`{"x":2,"y":3,"height":5}`)

	e := echo.New()

	req := httptest.NewRequest(
		http.MethodPost,
		"/estate/"+estateID.String()+"/tree",
		bytes.NewReader(body),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := server.PostEstateIdTree(ctx, openapi_types.UUID(estateID))

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
