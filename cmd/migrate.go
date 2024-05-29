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
		buildDgraphSchema()
	},
}

func buildDgraphSchema() error {
	pkg.SetupLogger()
	logger := pkg.GetLogger()
	err := pkg.SetupDgraphClient(logger)

	if err != nil {
		return err
	}

	// set schema
	// I was not able to find any source
	// which helps define facets in the schema itself
	op := &api.Operation{
		Schema: `
			type Author {
				name  
				email
				books
			}

			type Book {
				title
				genre
				published_year
			}

			type Borrower {
				name
				email
				books_borrowed
			}

			type Genre {
				name
			}


			name: string @index(exact) .
			title: string @index(exact) .
			books: [uid] @reverse .
			email: string @upsert @index(exact) .
			books_borrowed: [uid] @reverse .
			genre: [uid] @reverse .
			published_year: datetime @index(year) .

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
