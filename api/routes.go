package api

import (
	"github.com/elliot14A/go-dgraph-crud/api/author"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	author.InitRoutes(api)
}
