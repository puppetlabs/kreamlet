package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
)

func main() {

	if err := kubelet(); err != nil {
		log.Fatal(err)
	}
}

func kubelet() error {

	namespace := "services.linuxkit"
	kube := "kubelet"
	command := "kubeadm-init.sh"

	// create a new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), namespace)
	fmt.Println("loading kube container")
	//connect to kubelet container
	container, err := client.LoadContainer(
		ctx,
		kube,
	)
	if err != nil {
		return err
	}

	fmt.Println("getting OCI runtime spec")
	spec, err := container.Spec(ctx)
	if err != nil {
		return err
	}

	fmt.Println("creating container task")
	task, err := container.Task(ctx, nil)
	if err != nil {
		return err
	}

	defer task.Delete(ctx)

	fmt.Println("collecting result")
	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	}

	pspec := spec.Process
	pspec.Args = []string{command}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	creation := cio.NewIO(reader, os.Stdout, writer)
	process, err := task.Exec(ctx, command, pspec, creation)
	if err != nil {
		return err
	}

	// start the task
	if err := process.Start(ctx); err != nil {
		task.Delete(ctx)
		return err
	}

	status := <-exitStatusC
	_, _, err = status.Result()
	if err != nil {
		return err
	}

	return nil

}
