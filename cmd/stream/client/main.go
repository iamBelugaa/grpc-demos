package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// stream, err := client.StreamServerTime(context.Background(), &pb.StreamTimeRequest{IntervalSeconds: 2})
	// if err != nil {
	// 	log.Fatalf("[StreamServerTime] error : %v", err)
	// }

	// for {
	// 	resp, err := stream.Recv()
	// 	if err != nil {
	// if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatalf("[stream.Recv()] error : %v", err)
	// 	}
	// 	log.Printf("Current time - %s", resp.CurrentTime.AsTime().String())
	// }

	stream, err := client.StreamServerLog(context.Background())
	if err != nil {
		log.Fatalf("[StreamServerLog] error : %v", err)
	}

	for i := range 10 {
		if err := stream.Send(
			&pb.LogStreamRequest{
				Timestamp: timestamppb.Now(),
				LogLevel:  pb.LogLevel_INFO,
				Message:   fmt.Sprintf("Message #%d", i+1),
			},
		); err != nil {
			log.Fatalf("[stream.Send] error : %v", err)
		}

		time.Sleep(time.Second)
	}

	if resp, err := stream.CloseAndRecv(); err != nil {
		log.Fatalf("[CloseAndRecv] error : %v", err)
	} else {
		log.Printf("Entries logged - %d\n", resp.EntiresLogged)
	}
}
