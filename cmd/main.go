package main

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)

	e.GET("/hello", func(c echo.Context) error {
		id := c.QueryParam("id")

		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "id is required",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello User " + id,
		})
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
