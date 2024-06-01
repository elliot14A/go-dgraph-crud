package author

import "github.com/labstack/echo/v4"

func InitRoutes(group *echo.Group) {
	api := group.Group("/author")
	api.POST("/", create)
	api.GET("/", list)
	api.GET("/:uid", details)
	api.PATCH("/:uid", update)
	api.DELETE("/:uid", Delete)
}
