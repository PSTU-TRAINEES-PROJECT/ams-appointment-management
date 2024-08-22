package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
