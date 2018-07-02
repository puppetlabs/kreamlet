//Package main (stubclient) calls the grpc endpoint GetAdminCreds over tcp
package main

import (
	"log"
	"time"

	pb "github.com/puppetlabs/kreamlet/bootstrap/messaging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50091"
)

func main() {
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

	jtr, err := c.GetJoinToken(ctx, &pb.JoinTokenRequest{})
	if err != nil {
		log.Fatalf("could not invoke admin creds server: %v", err)
	}
	log.Printf("Join token response: %s", jtr)
}
