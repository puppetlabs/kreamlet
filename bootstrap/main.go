package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/puppetlabs/kreamlet/logging"
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

	logging.Info("getting new containerd client")

	// create a new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), namespace)
	logging.Info("loading kube container")

	//connect to kubelet container
	container, err := client.LoadContainer(
		ctx,
		kube,
	)
	if err != nil {
		return err
	}

	logging.Info("getting OCI runtime spec")
	spec, err := container.Spec(ctx)
	if err != nil {
		return err
	}

	logging.Info("creating container task")
	task, err := container.Task(ctx, nil)
	if err != nil {
		return err
	}

	defer task.Delete(ctx)

	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	}

	processSpec := spec.Process
	processSpec.Args = []string{"cat"}

	logging.Info("execing container task")
	if _, err := task.Exec(ctx, command, processSpec, cio.NullIO); err != nil {
		return err
	}

	logging.Info("collecting result")
	status := <-exitStatusC
	_, _, err = status.Result()
	if err != nil {
		return err
	}

	return nil
}
