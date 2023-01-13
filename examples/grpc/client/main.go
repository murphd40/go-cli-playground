package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/murphd40/go-playground/examples/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name = flag.String("name", "World", "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	stream, err := c.Chat(ctx)
	defer stream.CloseSend()
	if err != nil {
		log.Fatalf("Chat failed %s", err.Error())
	}

	// in this example the client establishes the connection but the server sends the first message
	waitc := make(chan int)
	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			log.Printf("Message received from Server: %s", msg.Message)

			stream.Send(&pb.ClientChat{
				Message: fmt.Sprintf("ACK: %s", msg.Message),
			})
		}
	}()
	<-waitc
}