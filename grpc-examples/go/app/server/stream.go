package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

// reference code https://qiita.com/suzuki0430/items/0ee2468f21cf48cc66ac

func StreamMain(myAddress string) error {
	lis, err := net.Listen("tcp", myAddress)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &stream{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}
	return nil
}

type stream struct {
	pb.UnimplementedGreetServiceServer
}

func (*stream) GreetManyTimes(req *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 2; i++ {
		res := &pb.GreetManytimesResponse{
			Result: fmt.Sprintf("Hello %s number %d", firstName, i),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*stream) LongGreet(stream pb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*stream) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		time.Sleep(2000 * time.Millisecond)
		firstName := req.GetGreeting().GetFirstName()
		result = "Hello " + firstName + "! "
		if sendErr := stream.Send(&pb.GreetEveryoneResponse{Result: result}); sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
		time.Sleep(2000 * time.Millisecond)
		result = "And hello " + firstName + "! "
		if sendErr := stream.Send(&pb.GreetEveryoneResponse{Result: result}); sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}
