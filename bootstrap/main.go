package main

import (
	"fmt"
	"github.com/puppetlabs/bootstrap/kubelet"
	"log"
	"time"
)

func main() {

	for true {
		// Time to wait for the kubelet container to start
		time.Sleep(30 * time.Second)
		run("services.linuxkit", nextExecID(), "kubelet", []string{"kubeadm-init.sh"})

	}
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
