package pkg

import (
	"log/slog"

	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var dgraphClient *dgo.Dgraph

func SetupDgraphClient(logger *slog.Logger) error {
	conn, err := grpc.Dial("alpha:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Error("Error dialing grpc server: ", "error", err.Error())
		return err
	}

	dgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return nil
}

func GetDgraphClient() *dgo.Dgraph {
	return dgraphClient
}
