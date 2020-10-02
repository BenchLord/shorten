package cmd

import (
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the shorten service",
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
