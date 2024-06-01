package genre

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	genreActions "github.com/elliot14A/go-dgraph-crud/actions/genre"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/labstack/echo/v4"
)

func create(ctx echo.Context) error {
	createGenreBody := new(models.Genere)

	if err := ctx.Bind(createGenreBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid body"})
	}

	if err := ctx.Validate(createGenreBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := genreActions.Create(createGenreBody.Name)

	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrConflict {
				return ctx.JSON(http.StatusConflict, actionErr.Message)
			}
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, "create genre successfully")
}
