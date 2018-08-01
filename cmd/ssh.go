package cmd

import (
	"fmt"
	//"github.com/puppetlabs/kreamlet/client"
	"github.com/spf13/cobra"
)

var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Creates a ssh conection to the Kubernetes controller.",
	Long:  `Creates a ssh conection to the Kubernetes controller.`,
	Run: func(ccmd *cobra.Command, args []string) {
		SshController()
	},
}

func init() {

	RootCmd.AddCommand(SshCmd)

}

func SshController() {
	fmt.Println("comming soon")
	return
}
