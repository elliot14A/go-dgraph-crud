package genre

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	genreActions "github.com/elliot14A/go-dgraph-crud/actions/genre"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/labstack/echo/v4"
)

func update(c echo.Context) error {
	updateGenreReq := new(models.Genere)

	uid := c.Param("uid")

	if err := c.Bind(updateGenreReq); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid body")
	}

	if err := c.Validate(updateGenreReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := genreActions.Details(uid)
	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrNotFound {
				return c.JSON(http.StatusNotFound, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = genreActions.Update(uid, updateGenreReq.Name)

	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrConflict {
				return c.JSON(http.StatusConflict, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "update genre successfully")
}
