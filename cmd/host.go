package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mymmrac/go-chat/pkg/server"
)

func init() {
	rootCmd.AddCommand(hostCmd)
}

var hostCmd = &cobra.Command{
	Use:   "host port",
	Short: "Host chat server on port",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set("port", args[0])
		return server.Run()
	},
}
