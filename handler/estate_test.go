package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestPostEstate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	mockRepo.
		EXPECT().
		CreateEstate(gomock.Any(), 5, 10).
		Return("5fc3572c-2847-48b0-80e6-66cdc91d5a7f", nil)

	server := &Server{
		Repository: mockRepo,
	}

	body := []byte(`{"width":5,"length":10}`)

	e := echo.New()

	req := httptest.NewRequest(
		http.MethodPost,
		"/estate",
		bytes.NewReader(body),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := server.PostEstate(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusCreated,
			rec.Code,
		)
	}
}
