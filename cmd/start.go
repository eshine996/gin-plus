package cmd

import (
	"fmt"
	r "gin-plus/internal/cmd"
	"github.com/spf13/cobra"
)

var (
	config string
)

func init() {
	flags := startCmd.PersistentFlags()
	flags.StringVarP(&config, "file", "f", "./configs/config.yaml", "config file path")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start project",
	Long:  "start project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config)
		r.TestStart()
	},
}
