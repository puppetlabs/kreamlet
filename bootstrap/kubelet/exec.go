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
//returns the command output only if the captureOutput bool parameter is set to true
func Run(namespace string, processID string, containerID string, command []string, captureOutput bool) (string, error) {

	log.Printf("Run::Entered with namespace %v, processID %v, containerID %v, captureOutput %v and command %v\n", namespace, processID, containerID, captureOutput, command)

	stdErr = ""
	output = ""

	// create a new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")

	if err != nil {
		return "", err
	}
	defer client.Close()

	log.Printf("Run::Setting container namespace to %v\n", namespace)
	ctx := namespaces.WithNamespace(context.Background(), namespace)

	log.Printf("Run::Loading container %v\n", containerID)
	//connect to kubelet container
	container, err := client.LoadContainer(
		ctx,
		containerID,
	)
	if err != nil {
		return "", err
	}

	log.Printf("Run::Getting OCI runtime specification\n")
	spec, err := container.Spec(ctx)
	if err != nil {
		return "", err
	}

	log.Printf("Run::Getting container task\n")
	task, err := container.Task(ctx, nil)
	if err != nil {
		return "", err
	}

	defer cleanup(ctx, task)

	pspec := spec.Process
	pspec.Args = command

	log.Printf("Run::Creating new process with processID %v on container task with command %v\n", processID, command)

	var creator cio.Creator
	if captureOutput {
		creator = cio.NewCreator(withIO)
	} else {
		creator = cio.NewCreator(cio.WithStdio)
	}
	process, err := task.Exec(ctx, processID, pspec, creator)
	if err != nil {
		return "", err
	}

	exitStatusC, err := process.Wait(ctx)
	if err != nil {
		fmt.Printf("Run::Error waiting for process %v", err)
		return "", err
	}

	log.Printf("Run::Starting process\n")
	if err := process.Start(ctx); err != nil {
		fmt.Printf("Run::Error starting process %v", err)
		return "", err
	}

	log.Printf("Run::Collecting result\n\n\n\n")
	status := <-exitStatusC
	statusCode, exitedAt, err := status.Result()
	log.Printf("Run::Exited from command %v with status [%v] at [%v], error is [%v], output is \n\t[%v]\n and stdErr is \n\t[%v]\n", command, statusCode, exitedAt, err, output, stdErr)

	log.Printf("Run::the output is [%v]\n\n\n\n", output)

	s := output
	return s, err
}

func cleanup(ctx context.Context, task containerd.Task) {
	log.Printf("cleanup::cleaning up task with ID %v and PID %v \n", task.ID(), task.Pid())
	task.Delete(ctx)
}

func withIO(opt *cio.Streams) {
	fmt.Println("withIO::setting up our io")
	withOurStreams(os.Stdin, os.Stdout, os.Stderr)(opt)
}

func withOurStreams(stdin io.Reader, stdout, stderr io.Writer) cio.Opt {
	return func(opt *cio.Streams) {
		opt.Stdin = stdin
		opt.Stdout = outWriter{}
		opt.Stderr = errWriter{}
	}
}

type outWriter struct{}

func (outWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	fmt.Printf("outWriter::Write::writing output (len %v):[%v]\n", len(p), s)
	output = output + s
	return len(p), nil
}

type errWriter struct{}

func (errWriter) Write(p []byte) (n int, err error) {
	s := string(p[:])
	fmt.Printf("errWriter::Write::writing output (len %v):[%v]\n", len(p), s)
	stdErr = stdErr + s
	return len(p), nil
}

var output string
var stdErr string
