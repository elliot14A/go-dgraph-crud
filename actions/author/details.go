package author

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/actions"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Details(uid string) (*models.Author, error) {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      q(func: uid(%s)) {
    		email
    		name
     }
  	}
  `, uid)

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Query: query,
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var response struct {
		Q []*models.Author `json:"q"`
	}

	err = json.Unmarshal(res.GetJson(), &response)

	if err != nil {
		return nil, err
	}

	if len(response.Q) == 0 {
		logger.Error(fmt.Sprintf("Cannot find author with uid: %s", uid))
		actionErr := actions.ActionErr{
			Message: fmt.Sprintf("Cannot find author with uid: %s", uid),
			Type:    actions.ErrNotFound,
		}
		return nil, actionErr
	}

	return response.Q[0], nil
}
