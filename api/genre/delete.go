package genre

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	genreActions "github.com/elliot14A/go-dgraph-crud/actions/genre"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	uid := c.Param("uid")
	err := genreActions.Delete(uid)
	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrNotFound {
				return c.JSON(http.StatusNotFound, actionErr.Message)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted genre successfully")
}
