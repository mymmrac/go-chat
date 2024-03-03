package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mymmrac/go-chat/pkg/client"
)

func init() {
	rootCmd.AddCommand(joinCmd)
}

var joinCmd = &cobra.Command{
	Use:   "join address",
	Short: "Join chat server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set("address", args[0])
		return client.Run()
	},
}
