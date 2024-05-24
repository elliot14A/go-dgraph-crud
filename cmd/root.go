package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Simple echo server",
	Long:  `Simple curd api with dgraph and echo`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
