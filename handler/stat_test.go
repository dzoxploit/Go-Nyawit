package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestGetEstateStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	mockRepo.
		EXPECT().
		GetEstateByID(gomock.Any(), gomock.Any()).
		Return(true, 10, 20, nil)

	mockRepo.
		EXPECT().
		GetTreeHeights(gomock.Any(), gomock.Any()).
		Return([]int{5, 10, 15}, nil)

	server := &Server{
		Repository: mockRepo,
	}

	e := echo.New()

	req := httptest.NewRequest(
		http.MethodGet,
		"/stats",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	id := openapi_types.UUID(uuid.New())

	err := server.GetEstateIdStats(ctx, id)

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}
