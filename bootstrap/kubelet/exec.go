package kubelet

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
)

//Run the passed command on the specified containerID with namespace, assigning processID for diagnostics
func Run(namespace string, processID string, containerID string, command []string) (string, error) {

	log.Printf("Entered with namespace %v, processID %v, containerID %v and command %v\n", namespace, processID, containerID, command)

	// create a new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")

	if err != nil {
		return "", err
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
		return "", err
	}

	log.Printf("Getting OCI runtime specification\n")
	spec, err := container.Spec(ctx)
	if err != nil {
		return "", err
	}

	log.Printf("Getting container task\n")
	task, err := container.Task(ctx, nil)
	if err != nil {
		return "", err
	}

	defer cleanup(ctx, task)

	pspec := spec.Process
	pspec.Args = command

	log.Printf("Creating new process with processID %v on container task with command %v\n", processID, command)
	process, err := task.Exec(ctx, processID, pspec, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return "", err
	}

	exitStatusC, err := process.Wait(ctx)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Starting process\n")
	if err := process.Start(ctx); err != nil {
		return "", err
	}

	log.Printf("Collecting result\n")
	status := <-exitStatusC
	statusCode, exitedAt, err := status.Result()
	log.Printf("Exited with status %v at %v, error is %v\n", statusCode, exitedAt, err)

	if err != nil {
		return "", err
	}

	if statusCode != 0 {
		return "", fmt.Errorf("Status code of %v recevied when trying to execute command %v", statusCode, command)
	}
	return output, nil

}

func cleanup(ctx context.Context, task containerd.Task) {
	log.Printf("cleaning up task with ID %v and PID %v \n", task.ID(), task.Pid())
	task.Delete(ctx)
}

func withIO(opt *cio.Streams) {
	withOurStreams(os.Stdin, os.Stdout, os.Stderr)(opt)
}

func withOurStreams(stdin io.Reader, stdout, stderr io.Writer) cio.Opt {
	return func(opt *cio.Streams) {
		opt.Stdin = stdin
		opt.Stdout = ourWriter{}
		opt.Stderr = stderr
	}
}

type ourWriter struct{}

func (ourWriter) Write(p []byte) (n int, err error) {
	output = string(p[:])
	return 2, nil
}

var output string
