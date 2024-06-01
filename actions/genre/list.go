package genre

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func List() ([]*models.Genere, error) {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      q(func: type(Genre)) {
        uid
    		name
     }
  	}
  `)

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Query: query,
	})

	var genreList struct {
		Q []*models.Genere `json:"q"`
	}

	err = json.Unmarshal(res.GetJson(), &genreList)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return genreList.Q, nil
}
