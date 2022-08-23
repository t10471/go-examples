package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
	"google.golang.org/grpc"
)

const (
	client = "client"
	server = "server"
	bi     = "bi"
)

// reference code https://qiita.com/suzuki0430/items/0ee2468f21cf48cc66ac

func StreamMain(address, method string) error {
	os.Remove("s")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	switch method {
	case server:
		doServerStreaming(c)
	case client:
		doClientStreaming(c)
	case bi:
		doBiDiStreaming(c)
	default:
		return fmt.Errorf("invalid method: %s", method)
	}

	return nil
}

func doServerStreaming(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	requests := []*pb.GreetManyTimesRequest{
		&pb.GreetManyTimesRequest{
			Greeting: &pb.Greeting{
				FirstName: "Stephane",
				LastName:  "Maarek",
			},
		},
		&pb.GreetManyTimesRequest{
			Greeting: &pb.Greeting{
				FirstName: "John",
				LastName:  "Maarek",
			},
		},
	}
	for _, req := range requests {
		resStream, err := c.GreetManyTimes(context.Background(), req)
		if err != nil {
			log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
		}
		for {
			msg, err := resStream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
			}
			log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
		}
	}
}

func doClientStreaming(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*pb.LongGreetRequest{
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Stephane",
			},
		},
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "John",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBiDiStreaming(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*pb.GreetEveryoneRequest{
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "Stephane",
			},
		},
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "John",
			},
		},
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "Lucy",
			},
		},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("EOF\n")
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
