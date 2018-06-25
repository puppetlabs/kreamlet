package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"strings"
	"time"

	"crypto/rand"

	"github.com/puppetlabs/kreamlet/bootstrap/kubelet"
	pb "github.com/puppetlabs/kreamlet/bootstrap/messaging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port                    = ":50091"
	adminCredsFromLocalFile = true
	pathToCredsFile         = "/etc/kubernetes/admin.conf"
)

func main() {
	//in a background go routine start a tcp listener for grpc connections
	go startListening()

	// Time to wait for the kubelet container to start
	time.Sleep(5 * time.Second)

	taskRoot := random()

	err := initKubeAdm(taskRoot)
	if err != nil {
		//if initialising kube admin doesn't succeed, there is nothing we can do here, so just exit
		log.Fatal(err)
	}

	//for now keep the main thread alive whilst we wait for a tcp connection
	//(we should be make a listener channel and waiting for it to complete?)
	for true {
		fmt.Println("Listening.....")
		time.Sleep(time.Minute)
	}
}

func initKubeAdm(taskRoot string) error {
	var output string
	var err error

	output, err = kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", []string{"kubeadm-init.sh"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is \n%v\n", output)
	joinToken, err := getJoinToken(output)
	fmt.Printf("join token is \n%v\n", joinToken)
	return err

}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// GetAdminCreds implements AdminCredsServer.GetAdminCreds
func (s *server) GetAdminCreds(ctx context.Context, in *pb.AdminCredsRequest) (*pb.AdminCredsResponse, error) {
	r := &pb.AdminCredsResponse{}

	var content []byte
	var err error

	if adminCredsFromLocalFile {
		content, err = getAdminCredsFromLocalFile()
	} else {
		content, err = getAdminCredsViaContainerd()

	}

	if err == nil {
		r.Content = content
		r.StatusCode = pb.StatusCode_Ok
	} else {
		r.StatusCode = pb.StatusCode_Failed
		r.Message = err.Error()
	}
	fmt.Printf("Returning %v and err of %v\n", r, err)
	return r, err
}

func getAdminCredsViaContainerd() ([]byte, error) {
	var output string
	var err error
	output, err = kubelet.Run("services.linuxkit", nextExecID(random()), "kubelet", []string{"cat", "/etc/kubernetes/admin.conf"})
	fmt.Printf("output is \n%v and err is %v\n", output, err)
	return []byte(output), err
}

func getAdminCredsFromLocalFile() ([]byte, error) {
	return ioutil.ReadFile(pathToCredsFile)
}

func startListening() {
	fmt.Printf("Entered startListening about to listen on port %v\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminCredsServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("About to listen on port %v\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("Listening on port %v\n", port)
}

func random() string {
	n := 3
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

func nextExecID(taskRoot string) string {
	execIDCounter = execIDCounter + 1
	return fmt.Sprintf("%v_%v", taskRoot, execIDCounter)
}

var execIDCounter = 0

func getJoinToken(output string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	re := regexp.MustCompile("kubeadm join .* --token ([^ ]+) ")

	for scanner.Scan() {
		s := scanner.Text()
		matches := re.FindStringSubmatch(s)
		if len(matches) == 2 {
			return matches[1], nil
		}
	}

	return "", nil
}
