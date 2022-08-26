package server

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func CancelMain(myAddress string) error {

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("failed to zap.NewDevelopment: %v", err)
		return err
	}

	lis, err := net.Listen("tcp", myAddress)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_zap.UnaryServerInterceptor(zapLogger)),
	)
	pb.RegisterGreeterServer(s, &canelServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}
	return nil
}

type canelServer struct {
	pb.UnimplementedGreeterServer
}

func (s *canelServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}, nil
}
