package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
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

	runCmd(taskRoot, []string{"ls"}, true)
	err := initKubeAdm(taskRoot)
	if err != nil {
		//if initialising kube admin doesn't succeed, there is nothing we can do here, so just exit
		log.Fatalf("kube admin initialisation failed %+v", err)
	}

	// joinToken, err := getJoinToken(taskRoot)
	// fmt.Printf("Join token is %v", joinToken)

	//for now keep the main thread alive whilst we wait for a tcp connection
	//(we should be make a listener channel and waiting for it to complete?)
	for true {
		fmt.Println("Listening.....")
		time.Sleep(time.Minute)
	}
}

func runCmd(taskRoot string, cmd []string, captureOutput bool) {
	output, err := kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", cmd, captureOutput)
	if err != nil {
		fmt.Printf("runCmd::Ran: %v with captureOutput [%v] and output: %v\nErr: %v\n\n\n\n\n", cmd, captureOutput, output, err)
	}
}

func initKubeAdm(taskRoot string) error {
	var output string
	var err error

	output, err = kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", []string{"kubeadm-init.sh"}, true)
	if err != nil {
		fmt.Printf("initKubeAdm::Error occured running kubeadm-init.sh - %v", err)
		os.Exit(1)
	}
	fmt.Printf("initKubeAdm::output is [%v]\n\n\n", output)
	// joinToken, err := extractJoinTokenFromInitOutput(output)
	// fmt.Printf("initKubeAdm::extractJoinToken returning: \n output %v, \n jt %v, \n err %v.", output, joinToken, err)
	return err
}

func getJoinToken(taskRoot string) (string, error) {
	var output, joinToken string
	var err error

	output, err = kubelet.Run("services.linuxkit", nextExecID(taskRoot), "kubelet", []string{"kubeadm", "token", "create"}, true)
	fmt.Printf("getJoinToken::the output is [%v] err is %v\n", output, err)

	if err == nil {
		joinToken, err = extractJoinTokenFromTokenCreate(output)
		fmt.Printf("getJoinToken::jt is [%v] err is %v\n", joinToken, err)
	}
	fmt.Printf("getJoinToken::returning: \n output [%v], \n jt [%v], \n err [%v].", output, joinToken, err)
	return joinToken, err
}

// server is used to implement AdminCredsServer
type server struct{}

// GetJoinToken implements AdminCredsServer.GetJoinToken
func (s *server) GetJoinToken(ctx context.Context, in *pb.JoinTokenRequest) (*pb.JoinTokenResponse, error) {
	jt, err := getJoinToken(nextExecID(random()))
	r := &pb.JoinTokenResponse{}
	r.JoinToken = jt
	return r, err
}

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
	output, err = kubelet.Run("services.linuxkit", nextExecID(random()), "kubelet", []string{"cat", "/etc/kubernetes/admin.conf"}, true)
	fmt.Printf("output is \n%v and err is %v\n", output, err)
	return []byte(output), err
}

func getAdminCredsFromLocalFile() ([]byte, error) {
	return ioutil.ReadFile(pathToCredsFile)
}

func startListening() {
	fmt.Printf("startListening::Entered startListening about to listen on port %v\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("startListening::failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminCredsServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("startListening::About to listen on port %v\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("startListening::failed to serve: %v", err)
	}
	fmt.Printf("startListening::Listening on port %v\n", port)
}

func extractJoinTokenFromInitOutput(output string) (string, error) {
	return extractRegex("kubeadm join .* --token ([^ ]+) ", output)
}

func extractJoinTokenFromTokenCreate(output string) (string, error) {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(output, ""), nil
}
func extractRegex(regex string, output string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	re := regexp.MustCompile(regex)

	for scanner.Scan() {
		s := scanner.Text()
		matches := re.FindStringSubmatch(s)
		if len(matches) == 2 {
			return matches[1], nil
		}
	}

	return "", nil
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
