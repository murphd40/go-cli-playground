package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/murphd40/go-playground/examples/grpc/helloworld/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) Chat(stream pb.Greeter_ChatServer) error {
	for {
		message := fmt.Sprintf("Hello %s", time.Now())
		log.Printf("Sending message %s", message)
		stream.Send(&pb.ServerChat{
			Message: message,
		})

		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		log.Printf("Received message %s", msg.GetMessage())

		time.Sleep(5 * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

