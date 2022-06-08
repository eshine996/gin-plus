package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-plus",
	Short: "Engineering example based on gin framework",
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
