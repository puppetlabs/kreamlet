package cmd

import (
	"fmt"
	"github.com/puppetlabs/kreamlet/client"
	//	"github.com/scotty-c/kream-v2/logging"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Creates a kubernetes controller.",
	Long:  `Creates a kubernetes controller.`,
	Run: func(ccmd *cobra.Command, args []string) {
		StartController()
	},
}

var sshPort, kubePort, cpus, memory, disk string

func init() {

	RunCmd.Flags().StringVarP(&sshPort, "ssh-port", "s", "2222", "The port ssh will listen to locally")
	RunCmd.Flags().StringVarP(&kubePort, "kube-port", "k", "6443", "The port kubectl will connect to locally")
	RunCmd.Flags().StringVarP(&cpus, "cpus", "c", "2", "The number of cpus to give the controller")
	RunCmd.Flags().StringVarP(&memory, "memory", "m", "2048", "The amount of memory to give the master")
	RunCmd.Flags().StringVarP(&disk, "disk-space", "g", "4G", "The amount of disk the controller os has")

	RunCmd.MarkFlagRequired("sshPort")
	RunCmd.MarkFlagRequired("KubePort")
	RunCmd.MarkFlagRequired("cpus")
	RunCmd.MarkFlagRequired("memory")
	RunCmd.MarkFlagRequired("disk")

	RootCmd.AddCommand(RunCmd)

}

func StartController() error {

	err := client.Run(sshPort, kubePort, cpus, memory, disk)
	if err != nil {
		fmt.Println(err)
	}
	return err

}
