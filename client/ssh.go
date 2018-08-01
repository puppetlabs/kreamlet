package client

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

func Ssh(user, host string) (*ssh.Client, *ssh.Session, error) {

	homedir := os.Getenv("HOME")

	file := homedir + "/.kream/ssh/id_rsa"

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("there was an error reading the ssh key")
		os.Exit(1)
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		fmt.Println("there was an error reading the ssh key")
		os.Exit(1)
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(key)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
