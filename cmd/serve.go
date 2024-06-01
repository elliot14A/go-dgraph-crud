package cmd

import (
	"fmt"
	"log"

	"github.com/elliot14A/go-dgraph-crud/api"
	"github.com/elliot14A/go-dgraph-crud/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts server for the application",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

type JsonValidator struct {
	validator *validator.Validate
}

func (v *JsonValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func serve() {
	app := echo.New()

	app.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				fmt.Printf("REQUEST: uri: %v, status: %v, latency: %v ns\n", v.URI, v.Status, v.Latency.Nanoseconds())
				return nil
			},
		}),
	)

	app.Validator = &JsonValidator{validator: validator.New()}

	pkg.SetupLogger()
	logger := pkg.GetLogger()
	err := pkg.SetupDgraphClient(logger)
	if err != nil {
		logger.Error("error initializing dgraph client")
		return
	}

	app.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello There!")
	})

	api.InitRoutes(app)

	log.Fatal(app.Start(":8000"))
}
