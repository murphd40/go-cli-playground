package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/murphd40/go-playground/examples/grpc/pullserver/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGenericStreamServer
	tick <-chan time.Time
}

func (s *server) Connect(stream pb.GenericStream_ConnectServer) error {
	for range s.tick {
		err := stream.Send(&pb.ServerMessage{
			Request: pb.Request_ECHO,
		})

		if err != nil {
			return err
		}

		msg, err := stream.Recv()

		if err != nil {
			return err
		}

		log.Printf("Message from client: %s", msg.Body)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tick := time.Tick(5 * time.Second)

	server := &server{
		tick: tick,
	}

	s := grpc.NewServer()
	pb.RegisterGenericStreamServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}