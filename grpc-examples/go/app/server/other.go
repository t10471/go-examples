package server

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

func OtherMain(myAddress string, sleep int) error {
	lis, err := net.Listen("tcp", myAddress)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	pb.RegisterOtherServer(s, &other{sleep: time.Duration(sleep) * time.Second})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}
	return nil
}

type other struct {
	pb.UnimplementedOtherServer
	sleep time.Duration
}

func (s *other) CallOther(ctx context.Context, in *pb.OtherRequest) (*pb.OtherReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Printf("sleep: %v", s.sleep)
	time.Sleep(s.sleep)
	log.Printf("sleep end")
	return &pb.OtherReply{Message: "Hello " + in.GetName()}, nil
}

func (s *other) CallOtherV2(ctx context.Context, in *pb.OtherRequest) (*pb.OtherReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.OtherReply{Message: "Hello " + in.GetName()}, nil
}

