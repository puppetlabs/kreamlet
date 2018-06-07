package main

import (
	"fmt"
	"log"

	"github.com/puppetlabs/kreamlet/bootstrap/kubelet"
)

func main() {
	run("services.linuxkit", nextExecID(), "kubelet", []string{"pwd"})
	run("services.linuxkit", nextExecID(), "kubelet", []string{"ls", "-alt"})
	run("services.linuxkit", nextExecID(), "kubelet", []string{"kubeadm-init.sh"})

}

func run(namespace string, processID string, containerID string, command []string) {
	err := kubelet.Run(namespace, processID, containerID, command)
	if err != nil {
		log.Fatal(err)
	}

}
func nextExecID() string {
	execIDCounter = execIDCounter + 1
	return fmt.Sprintf("exec_id_%v", execIDCounter)
}

var execIDCounter = 0
