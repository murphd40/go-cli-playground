package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	pb "github.com/murphd40/go-playground/examples/grpc/pullserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGenericStreamClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := c.Connect(ctx)
	defer stream.CloseSend()
	if err != nil {
		log.Fatalf("Connect failed %s", err.Error())
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

			switch msg.Request {
			case pb.Request_HEALTH:
				log.Printf("Request received from server: %d", msg.Request)

				body := map[string]string {
					"status": "OK",
				}
				bs, _ := json.Marshal(body)

				stream.Send(&pb.ClientMessage{
					Body: bs,
				})
			}
		}
	}()
	<-waitc
}