package main

import (
	"log"

	"github.com/puppetlabs/bootstrap/kubelet"
)

func main() {

	err := kubelet.Run("services.linuxkit", "process_1", "kubelet", []string{"kubeadm-init.sh"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("let's have a look")
	err = kubelet.Run("services.linuxkit", "process_2", "kubelet", []string{"ls", "-alt"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("we reached this far")
}
