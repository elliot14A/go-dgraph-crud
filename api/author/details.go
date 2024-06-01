package author

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	authorActions "github.com/elliot14A/go-dgraph-crud/actions/author"
	"github.com/labstack/echo/v4"
)

func details(c echo.Context) error {
	uid := c.Param("uid")
	author, err := authorActions.Details(uid)
	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrNotFound {
				return c.JSON(http.StatusNotFound, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, author)
}
