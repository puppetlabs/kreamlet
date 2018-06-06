package kubelet

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
)

//Run the passed command on the specified containerID with namespace, assigning processID for diagnostics
func Run(namespace string, processID string, containerID string, command []string) error {

	log.Printf("Entered with namespace %v, processID %v, containerID %v and command %v\n", namespace, processID, containerID, command)

	// create a new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")

	if err != nil {
		return err
	}
	defer client.Close()

	log.Printf("Setting container namespace to %v\n", namespace)
	ctx := namespaces.WithNamespace(context.Background(), namespace)

	log.Printf("Loading container %v\n", containerID)
	//connect to kubelet container
	container, err := client.LoadContainer(
		ctx,
		containerID,
	)
	if err != nil {
		return err
	}

	log.Printf("Getting OCI runtime specification\n")
	spec, err := container.Spec(ctx)
	if err != nil {
		return err
	}

	log.Printf("Getting container task\n")
	task, err := container.Task(ctx, nil)
	if err != nil {
		return err
	}

	defer task.Delete(ctx)

	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	}

	pspec := spec.Process
	pspec.Args = command

	log.Printf("Creating new process with processID %v on container task with command %v\n", processID, command)
	process, err := task.Exec(ctx, processID, pspec, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}

	log.Printf("Starting process\n")
	if err := process.Start(ctx); err != nil {
		return err
	}

	log.Printf("Collecting result\n")
	status := <-exitStatusC
	statusCode, exitedAt, err := status.Result()
	log.Printf("Exited with status %v at %v, error is %v\n", statusCode, exitedAt, err)

	if err != nil {
		return err
	}

	return nil

}
