package client

import (
	"fmt"
	pb "github.com/puppetlabs/kreamlet/bootstrap/messaging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	address = "localhost:50091"
)

func Creds() error {

	fmt.Printf("Connecting to grpc server\n")

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
	log.Printf("Response: %s", r)

	err = ioutil.WriteFile(homedir+"/.kream/admin.conf", r.Content, 0644)

	if err != nil {
		log.Fatalf("could not write to file: %v", err)
	}

	return err
}
