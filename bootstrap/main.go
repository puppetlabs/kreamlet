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

	//connect to kubelet container
	container, err := client.LoadContainer(
		ctx,
		kube,
	)
	if err != nil {
		return err
	}

	spec, err := container.Spec(ctx)
	if err != nil {
		return err
	}

	task, err := container.Task(ctx, nil)
	if err != nil {
		return err
	}

	defer task.Delete(ctx)

	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	creation := cio.NewIO(reader, os.Stdout, writer)
	if _, err := task.Exec(ctx, command, spec.Process, creation); err != nil {
		return err
	}

	// start the task
	if err := task.Start(ctx); err != nil {
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
