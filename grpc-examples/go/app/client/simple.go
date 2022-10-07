package client

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

func SimpleMain(address string) error {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Printf("context Done")
			break LOOP
		default:
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
			if err != nil {
				log.Printf("could not greet: %v", err)
			} else {
				log.Printf("greeting: %s", r.GetMessage())
			}
			time.Sleep(1 * time.Second)
		}
	}
	log.Printf("finished")

	return nil
}
