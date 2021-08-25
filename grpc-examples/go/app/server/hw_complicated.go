package server

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/t10471/go-examples/grpc-examples/go/proto"
)

func HWComplicatedMain(myAddress, otherAddress string) error {
	lis, err := net.Listen("tcp", myAddress)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &hwComplicated{otherAddress: otherAddress})
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

type hwComplicated struct {
	pb.UnimplementedGreeterServer
	otherAddress string
}

type result struct {
	retStrCh   chan string
	errCh      chan error
	canceledCh chan error
	canceled   bool
}

func newResult() result {
	return result{
		retStrCh:   make(chan string),
		errCh:      make(chan error),
		canceledCh: make(chan error),
		canceled:   false,
	}
}

func (s *hwComplicated) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	result := newResult()
	defer func() {
		close(result.retStrCh)
		close(result.errCh)
		close(result.canceledCh)
	}()
	go func() {
		stopCh := make(chan interface{})
		ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer func() {
			close(stopCh)
			cancel()
		}()
		go func() {
			msg, err := s.callOther(ctx2, in.Name)
			if result.canceled {
				log.Printf("return  but already canceled")
				log.Printf("return error : %v", err)
				return
			}
			if err != nil {
				log.Printf("return error : %v", err)
				result.errCh <- err
			}
			stopCh <- struct{}{}
			result.retStrCh <- msg
		}()
		select {
		case <-ctx.Done():
			log.Printf("context done: %v", ctx.Err())
			result.canceledCh <- ctx.Err()
		case <-stopCh:
			break
		}
	}()
	select {
	case msg := <-result.retStrCh:
		return &pb.HelloReply{Message: msg}, nil
	case <-result.canceledCh:
		result.canceled = true
		log.Printf("return empty msg")
		return &pb.HelloReply{Message: ""}, nil
	case err := <-result.errCh:
		log.Printf("Received error: %v", err)
		return nil, err
	}
}

func (s *hwComplicated) callOther(ctx context.Context, name string) (string, error) {
	conn, err := grpc.Dial(s.otherAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	c := pb.NewOtherClient(conn)

	r, err := c.CallOther(ctx, &pb.OtherRequest{Name: name})
	if err != nil {
		log.Printf("could not call other: %v", err)
		return "", err
	}
	log.Printf("CallOther: %s", r.GetMessage())
	return r.GetMessage(), nil
}

