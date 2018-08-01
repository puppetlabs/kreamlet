package cmd

import (

	"bufio"
	"fmt"
	"github.com/puppetlabs/kreamlet/client"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
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

	client, session, err := client.Ssh("root", "127.0.0.1:2222")
	if err != nil {
		fmt.Println("could not connect to the node. please make sure it is running")
	}
	defer client.Close()
	// Set redirection of stdout and stderr
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	in, _ := session.StdinPipe()
	// Set up terminal
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Println("request for pseudo terminal failed: %s", err)
		os.Exit(1)
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		fmt.Println("failed to start shell: %s", err)
		os.Exit(1)
	}
	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}
}

func scanConfig() string {
	config, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	config = strings.Trim(config, "\n")
	return config
}
