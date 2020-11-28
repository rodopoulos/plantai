package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the binary name, intended to be used in the CLI.
var RootCmd = &cobra.Command{
	Use:   "plantai",
	Short: "server for plantai core business logic",
	Long:  "server for plantai core business logic",
}

// Execute executes a Cobra command.
func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
