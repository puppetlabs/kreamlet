package cmd

import (
	"fmt"
	"github.com/puppetlabs/kreamlet/client"
	"github.com/spf13/cobra"
	"os"
)

var PowCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "Powers off nodes.",
	Long:  `Powers off Kubernetes controllers and nodes.`,
	Run: func(ccmd *cobra.Command, args []string) {
		PowerOff()
	},
}

func init() {

	RootCmd.AddCommand(PowCmd)

}

func PowerOff() error {

	var (
		command = "poweroff"
		user    = "root"
		host    = "127.0.0.1:2222"
	)

	client, session, err := client.Ssh(user, host)
	if err != nil {
		fmt.Println("could not connect to the node. please make sure it is running")
		os.Exit(1)
	}

	out, err := session.CombinedOutput(command)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	client.Close()

	fmt.Println("the nodes have been powered off")
	return err
}
