package client

import (
	"fmt"
	"os"
	"os/exec"
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

	// Check if iso is already downloaded
	if _, err := os.Stat(homedir + "/.kream/kube-master-efi.iso"); os.IsNotExist(err) {
		fileUrl := "https://s3.amazonaws.com/puppet-cloud-and-containers/kream/kube-master-efi.iso"
		err := DownloadFile(homedir+"/.kream/kube-master-efi.iso", fileUrl)
		if err != nil {
			panic(err)
		}

	}

	//TODO vendor in Linuxkit
	args := []string{
		"run",
		"hyperkit",
		"-publish", sshPort + ":22",
		"-publish", kubePort + ":6443",
		"-publish", "50091:50091",
		"-networking", "default",
		"-cpus", cpus,
		"-mem", memory,
		"-state", homedir + "/.kream/kube-master-state",
		"-disk", "size=" + disk,
		"-data-file", homedir + "/.kream/kube-master-state/metadata.json",
		"--uefi",
		"-iso", homedir + "/.kream/kube-master-efi.iso"}

	// This uses exec.Command func to pass the above args to run linuxkit
	execErr := exec.Command(binary, args...).Start()
	if execErr != nil {
		fmt.Println("failed here")
		fmt.Println(execErr)
	}
	return nil

}
