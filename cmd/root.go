package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "Simple chat server & client",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "Debug mode")
	must(viper.BindPFlags(rootCmd.PersistentFlags()))
}

func must(err error) {
	if err != nil {
		log.Fatal("Must", "error", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Debug("Run", "error", err)
		os.Exit(1)
	}
}
