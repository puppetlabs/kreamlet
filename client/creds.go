package client

import (
	"fmt"
	pb "github.com/puppetlabs/kreamlet/bootstrap/messaging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	address = "localhost:50091"
)

func Creds() error {

	fmt.Printf("waiting for the OS to boot\n")

	//wait for grpc server to start
	LoadingBar()

	// Getting users home dir to use later
	homedir := os.Getenv("HOME")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAdminCredsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAdminCreds(ctx, &pb.AdminCredsRequest{})
	if err != nil {
		log.Fatalf("could not invoke admin creds server: %v", err)
	}
	//log.Printf("Response: %s", r)

	err = ioutil.WriteFile(homedir+"/.kream/admin.conf", r.Content, 0644)

	if err != nil {
		log.Fatalf("could not write to file: %v", err)
	}

	input, err := ioutil.ReadFile(homedir + "/.kream/admin.conf")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "server:") {
			lines[i] = "    server: https://kubernetes.default:6444"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(homedir+"/.kream/admin.conf", []byte(output), 0644)
	if err != nil {
		log.Fatalf("could not write to file: %v", err)
	}

	fmt.Println("To connect to your cluster copy and paste the below line into your terminal")
	fmt.Println("export KUBECONFIG=~/.kream/admin.conf")

	return err
}
