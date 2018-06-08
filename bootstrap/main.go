package main

import (
	"fmt"
	"log"
	"time"

	"crypto/rand"

	"github.com/puppetlabs/kreamlet/bootstrap/kubelet"
)

func main() {
	taskRoot := random()
	var output string
	var err error

	// Time to wait for the kubelet container to start
	time.Sleep(5 * time.Second)

	// Initalise the cluster
	output, err = kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", []string{"kubeadm-init.sh"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is \n%v\n", output)

	// Get the admin creds
	output, err = kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", []string{"cat", "/etc/kubernetes/admin.conf"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is \n%v\n", output)
}

func random() string {
	n := 3
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

func nextExecID(taskRoot string) string {
	execIDCounter = execIDCounter + 1
	return fmt.Sprintf("%v_%v", taskRoot, execIDCounter)
}

var execIDCounter = 0
