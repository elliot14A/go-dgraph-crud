package author

import (
	"net/http"

	authorActions "github.com/elliot14A/go-dgraph-crud/actions/author"
	"github.com/labstack/echo/v4"
)

func list(c echo.Context) error {
	authorList, err := authorActions.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, authorList.Q)
}
