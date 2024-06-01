package author

import (
	"net/http"

	"github.com/elliot14A/go-dgraph-crud/actions"
	authorActions "github.com/elliot14A/go-dgraph-crud/actions/author"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/labstack/echo/v4"
)

func update(c echo.Context) error {
	updateAuthorInput := new(models.UpdateAuthorReq)

	uid := c.Param("uid")

	if err := c.Bind(updateAuthorInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid body"})
	}

	if err := c.Validate(updateAuthorInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	_, err := authorActions.Details(uid)
	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrNotFound {
				return c.JSON(http.StatusNotFound, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = authorActions.Update(uid, *updateAuthorInput)

	if err != nil {
		if actionErr, ok := err.(actions.ActionErr); ok {
			if actionErr.Type == actions.ErrConflict {
				return c.JSON(http.StatusConflict, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "updated author successfully")
}
