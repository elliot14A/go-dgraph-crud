package author

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/actions"
	"github.com/elliot14A/go-dgraph-crud/models"
	"github.com/elliot14A/go-dgraph-crud/pkg"
)

func Create(input models.CreateAuthorReq) error {
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
		SetNquads: []byte(fmt.Sprintf(`
		 	_:newUser <email> "%s" . 
			_:newUser <name> "%s" .
			_:newUser <dgraph.type> "Author" .
		`,
			input.Email,
			input.Name,
		)),
	}

	res, err := dg.NewTxn().Do(ctx, &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{&mu},
		CommitNow: true,
	})

	if err != nil {
		logger.Error("error creating new author", err)
		return err
	}

	if len(res.GetUids()) == 0 {
		logger.Error(fmt.Sprintf("Author with same email '%s' already exists", input.Email))
		err := actions.ActionErr{
			Message: fmt.Sprintf("Author with same email '%s' already exists", input.Email),
			Type:    actions.ErrConflict,
		}
		return err
	}

	return nil
}
