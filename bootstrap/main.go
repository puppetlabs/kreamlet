package main

import (
	"fmt"
	"github.com/puppetlabs/bootstrap/kubelet"
	"log"
	"time"
)

func main() {

	// Time to wait for the kubelet container to start
	time.Sleep(5 * time.Second)
	// Initalise the cluster
	run("services.linuxkit", nextExecID(), "kubelet", []string{"kubeadm-init.sh"})
	// Get the admin creds
	run("services.linuxkit", nextExecID(), "kubelet", []string{"cat", "/etc/kubernetes/admin.conf"})

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
