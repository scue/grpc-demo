package main

import (
	"context"
	"log"

	pb "github.com/scue/grpc-demo/backdog"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":2020", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExecClient(conn)

	r, err := c.Command(context.Background(), &pb.Request{Command: "echo 'hello gRPC!'"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(`stdout:`, string(r.Stdout))
	log.Println(`stderr:`, string(r.Stderr))
	log.Println(`error:`, r.Error)
	log.Println(`message:`, r.Message)
}
