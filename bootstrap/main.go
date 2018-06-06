package main

import (
	"github.com/puppetlabs/bootstrap/kubelet"
	//"log"
)

func main() {

	i := ContainerComand{"service.linuxkit", "kube", "kubelet", "kubeadm-init.sh"}
	initialise := i.containerExec.Run
	//if err != nil {
	//	log.Fatal(err)
	//	}
	return initialise
}
