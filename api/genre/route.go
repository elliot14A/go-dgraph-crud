package genre

import "github.com/labstack/echo/v4"

func InitRoutes(group *echo.Group) {
	api := group.Group("/genre")
	api.POST("/", create)
	api.GET("/", list)
	api.DELETE("/:uid", Delete)
	api.GET("/:uid", details)
	api.PATCH("/:uid", update)
}
