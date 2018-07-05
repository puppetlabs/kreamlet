package client

import (
	"fmt"
	pb "github.com/puppetlabs/kreamlet/bootstrap/messaging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	var homedir string = os.Getenv("HOME")

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

	fmt.Printf("Creating admin.conf\n")
	file, err := os.Create(homedir + "/.kream/admin.conf")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "test")

	return err
}
