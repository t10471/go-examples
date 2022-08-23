package server

import (
	"context"
	"log"
	"net"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

func HWSimpleMain(myAddress, otherAddress string, unwrapError bool) error {

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
	pb.RegisterGreeterServer(s, &hwSimple{otherAddress: otherAddress, unwrapError: unwrapError})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}
	return nil
}

type hwSimple struct {
	pb.UnimplementedGreeterServer
	otherAddress string
	unwrapError  bool
}

func (s *hwSimple) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	msg, err := s.callOther(ctx, in.Name)
	if err != nil {
		if s.unwrapError {
			return nil, toGRPCError(err)
		}
		return nil, err
	}

	return &pb.HelloReply{Message: msg}, nil
}

func (s *hwSimple) callOther(ctx context.Context, name string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(s.otherAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	c := pb.NewOtherClient(conn)

	var retStr string
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		r, err := c.CallOther(ctx, &pb.OtherRequest{Name: name})
		if err != nil {
			if !s.unwrapError {
				s, ok := status.FromError(err)
				if ok && s.Code() == codes.Canceled {
					log.Printf("CallOther canceled")
					return nil
				}
			}
			log.Printf("could not call other: %v", err)
			return xerrors.Errorf("failed to CallOther :%w", err)
			// return err
		}
		log.Printf("CallOther: %s", r.GetMessage())
		retStr = r.GetMessage()
		return nil

	})
	eg.Go(func() (err error) {
		r, err := c.CallOtherV2(ctx, &pb.OtherRequest{Name: name})
		if err != nil {
			if !s.unwrapError {
				s, ok := status.FromError(err)
				if ok && s.Code() == codes.Canceled {
					log.Printf("CallOtherV2 canceled")
					return nil
				}

			}
			log.Printf("could not call other: %v", err)
			return xerrors.Errorf("failed to CallOtherV2 :%w", err)
		}
		log.Printf("CallOtherV2: %s", r.GetMessage())
		return nil

	})
	if err := eg.Wait(); err != nil {
		return "", err
	}
	return retStr, nil
}

func toGRPCError(err error) error {
	if err == nil {
		return nil
	}
	s, ok := status.FromError(err)
	if ok {
		return s.Err()
	}
	unwrappedErr := xerrors.Unwrap(err)

	if unwrappedErr == nil {
		return status.Error(codes.Internal, err.Error())
	}
	return toGRPCError(unwrappedErr)
}
