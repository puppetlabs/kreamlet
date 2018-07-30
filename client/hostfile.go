package client

import (
	"fmt"
	"github.com/lextoumbourou/goodhosts"
	"os"
)

func Hostfile() error {

	hosts, _ := goodhosts.NewHosts()
	if hosts.Has("127.0.0.1", "kubernetes.default") {
		return nil
	}

	fmt.Println("please add 127.0.0.1 kubernetes.default to your hostfile")
	os.Exit(1)
	return nil
}
