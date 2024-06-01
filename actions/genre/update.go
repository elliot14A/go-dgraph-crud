package genre

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/actions"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Update(uid, name string) error {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      genre as var(func: eq(name, "%s"))
  	}
  `, name)

	mu := api.Mutation{
		Cond: `@if(eq(len(genre), 0))`,
		SetNquads: []byte(fmt.Sprintf(
			`
     	  <%s> <name> "%s" .
      `,
			uid,
			name,
		)),
	}

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{&mu},
		CommitNow: true,
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if len(res.GetTxn().GetPreds()) == 0 {
		logger.Error(fmt.Sprintf("Genre with same name '%s' already exists", name))
		actionErr := actions.ActionErr{
			Message: fmt.Sprintf("Genre with same name '%s' already exists", name),
			Type:    actions.ErrConflict,
		}
		return actionErr
	}

	return nil
}
