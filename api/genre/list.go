package genre

import (
	"net/http"

	genreActions "github.com/elliot14A/go-dgraph-crud/actions/genre"
	"github.com/labstack/echo/v4"
)

func list(c echo.Context) error {
	genreList, err := genreActions.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, genreList)
}
