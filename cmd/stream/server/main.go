package main

import (
	"log"
	"net"

	"github.com/iamNilotpal/grpc/internal/stream"
	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	defer func() {
		log.Println("Shutting down GRPC server")
		grpcServer.GracefulStop()
		log.Println("Shutdown completed GRPC server")
	}()

	streamService := stream.NewService()
	pb.RegisterStreamTimeServiceServer(grpcServer, streamService)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("TCP Error : %v", err)
	}
	defer func() {
		log.Println("Closing TCP connection")
		if err := lis.Close(); err != nil {
			log.Fatalln("Error closing TCP connection")
		}
		log.Println("TCP connection closed")
	}()

	log.Println("Starting GRPC server at http://localhost:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("GRPC Error : %v", err)
	}
}
