package main

import (
	"bytes"
	"log"
	"net"
	"os/exec"

	pb "github.com/scue/grpc-demo/backdog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Command(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", req.Command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	e := cmd.Run()
	return &pb.Response{
		Stdout:  stdout.Bytes(),
		Stderr:  stderr.Bytes(),
		Message: "run command oK!",
	}, e
}

func main() {
	lis, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExecServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
