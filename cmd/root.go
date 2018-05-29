package cmd

import (
	"github.com/puppetlabs/kreamlet/logging"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kream",
	Short: "kream",
	Long:  `Kream is a tool for spinning up Kubernetes clusters on Linuxkit locally.`,
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&logging.DebugEnabled, "debug", "d", false, "Print debug logging")
}
