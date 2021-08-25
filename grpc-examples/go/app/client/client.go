package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

func ClientMain(address string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return err
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return nil
}
