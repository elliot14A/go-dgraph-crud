package author

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	authorActions "github.com/elliot14A/go-dgraph-crud/actions/author"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/labstack/echo/v4"
)

func create(ctx echo.Context) error {
	createAuthorInput := new(models.CreateAuthorReq)

	if err := ctx.Bind(createAuthorInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid body"})
	}

	if err := ctx.Validate(createAuthorInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := authorActions.Create(*createAuthorInput)

	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrConflict {
				return ctx.JSON(http.StatusConflict, actionErr.Message)
			}
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "created author successfully")
}
