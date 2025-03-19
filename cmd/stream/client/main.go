package main

import (
	"context"
	"errors"
	"io"
	"log"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client conn error : %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Client conn close error : %v", err)
		}
	}()

	client := pb.NewStreamTimeServiceClient(conn)

	stream, err := client.StreamServerTime(context.Background(), &pb.StreamTimeRequest{IntervalSeconds: 2})
	if err != nil {
		log.Fatalf("[StreamServerTime] error : %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalf("[stream.Recv()] error : %v", err)
		}
		log.Printf("Current time - %s", resp.CurrentTime.AsTime().String())
	}
}
