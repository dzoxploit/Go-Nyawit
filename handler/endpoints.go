package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, server *Server) {
	generated.RegisterHandlers(e, server)
}
