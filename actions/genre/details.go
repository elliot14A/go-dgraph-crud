package genre

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Details(uid string) (*models.Genere, error) {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      q(func: uid(%s)) {
        uid
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
		Q []*models.Genere `json:"q"`
	}

	err = json.Unmarshal(res.GetJson(), &response)

	if err != nil {
		return nil, err
	}

	return response.Q[0], nil

}
