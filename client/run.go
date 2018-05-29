package client

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(sshPort string, kubePort string, cpus string, memory string, disk string) error {
	// Checking for linuxkit binary
	binary, lookErr := exec.LookPath("linuxkit")
	if lookErr != nil {
		panic(lookErr)
	}

	// Getting users home dir to use later
	var homedir string = os.Getenv("HOME")

	// Removing old state if the run function is called
	if _, err := os.Stat(homedir + "/.kream/kube-master-state"); err != nil {

		fmt.Println("Creating a new cluster state directory")
		os.Mkdir(homedir+"/.kream/kube-master-state", 0700)
		os.OpenFile(homedir+"/.kream/kube-master-state/metadata.json", os.O_RDONLY|os.O_CREATE, 0700)
	} else {
		fmt.Println("Removing the old cluster state")
		// We need to recreate the state folder to add the metadata.json file
		os.RemoveAll(homedir + "/.kream/kube-master-state")
		os.Mkdir(homedir+"/.kream/kube-master-state", 0700)
		os.OpenFile(homedir+"/.kream/kube-master-state/metadata.json", os.O_RDONLY|os.O_CREATE, 0700)
	}
	//TODO vendor in Linuxkit
	args := []string{
		"linuxkit",
		"run",
		"qemu",
		"-containerized",
		"-detached",
		"-publish", sshPort + ":22",
		"-publish", kubePort + ":6443",
		"-networking", "default",
		"-cpus", cpus,
		"-mem", memory,
		"-state", homedir + "/.kream/kube-master-state",
		"-disk", "size=" + disk,
		"-data-file", homedir + "/.kream/kube-master-state/metadata.json",
		"-iso", homedir + "/.kream/kube-master.iso",
	}

	env := os.Environ()

	// This uses syscall to pass the above args to run linuxkit
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	return execErr

	return nil
}
