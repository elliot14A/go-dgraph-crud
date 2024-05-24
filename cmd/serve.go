package cmd

import (
	"fmt"
	"log"

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

func serve() {
	app := echo.New()
	app.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				fmt.Printf("REQUEST: uri: %v, status: %v, latency: %v\n", v.URI, v.Status, v.Latency.Milliseconds())
				return nil
			},
		}),
	)

	app.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello There!")
	})

	log.Fatal(app.Start(":8000"))
}