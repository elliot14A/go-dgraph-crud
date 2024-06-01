package author

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func List() (*models.ListAuthorResponse, error) {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      q(func: type(Author)) {
        uid
    		email
    		name
     }
  	}
  `)

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Query: query,
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var authorList models.ListAuthorResponse

	err = json.Unmarshal(res.GetJson(), &authorList)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &authorList, nil
}
