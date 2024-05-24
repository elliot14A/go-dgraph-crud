package cmd

import (
	"context"

	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/elliot14A/go-dgraph-crud/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run dgraph schema migrations",
	Run: func(cmd *cobra.Command, args []string) {
		buildSchema()
	},
}

func buildSchema() error {
	pkg.SetupLogger()
	logger := pkg.GetLogger()
	err := pkg.SetupDgraphClient(logger)

	if err != nil {
		return err
	}

	// set schema
	op := &api.Operation{
		Schema: `
 		type Customer {
      id: ID!
      name: String!
      email: String!
      accounts: [Account] @hasInverse(field: customer)
    }

    type Account {
      id: ID!
      accountNumber: String!
      balance: Float!
      customer: Customer
      transactions: [Transaction] @hasInverse(field: account)
    }
		`,
	}

	ctx := context.Background()

	dgraphClient := pkg.GetDgraphClient()

	err = dgraphClient.Alter(ctx, op)
	if err != nil {
		logger.Error("Error altering schema", "error", err.Error())
		return err
	}

	return nil
}
