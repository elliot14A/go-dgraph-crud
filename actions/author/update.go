package author

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/actions"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Update(uid string, input models.UpdateAuthorReq) error {

	logger := pkg.GetLogger()
	dg := pkg.GetDgraphClient()
	ctx := context.Background()

	query := fmt.Sprintf(`
  	query {
      user as var(func: eq(email, "%s"))
  	}
  `, input.Email)

	mu := api.Mutation{
		Cond: `@if(eq(len(user), 0))`,
		SetNquads: []byte(fmt.Sprintf(
			`
     		<%s> <email> "%s" .
     	  <%s> <name> "%s" .
      `,
			uid,
			input.Email,
			uid,
			input.Name,
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
		logger.Error(fmt.Sprintf("Author with same email '%s' already exists", input.Email))
		actionErr := actions.ActionErr{
			Message: fmt.Sprintf("Author with same email '%s' already exists", input.Email),
			Type:    actions.ErrConflict,
		}
		return actionErr
	}

	return nil
}
