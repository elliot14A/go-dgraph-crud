package author

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/actions"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Delete(uid string) error {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()
	mu := api.Mutation{
		DelNquads: []byte(fmt.Sprintf(`
			<%s> * * .
		`, uid)),
	}

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Mutations: []*api.Mutation{&mu},
		CommitNow: true,
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	fmt.Println(res.GetTxn().Preds)

	if (len(res.GetTxn().Preds)) < 2 {
		logger.Error(fmt.Sprintf("author with id '%s' not found.", uid))
		actionErr := actions.ActionErr{
			Message: fmt.Sprintf("author with id '%s' not found.", uid),
			Type:    actions.ErrNotFound,
		}
		return actionErr
	}
	return nil
}
