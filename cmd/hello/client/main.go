package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.SayHello(
		context.Background(), &pb.SayHelloRequest{FirstName: "", LastName: "Deka"},
	)

	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			log.Fatalf("SayHello request error : %v", err)
		} else {
			log.Fatalf("SayHello request error : %v", s)
		}
	}
	fmt.Printf("Response : %s\n", resp.GetMessage())
}
