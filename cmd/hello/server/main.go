package main

import (
	"log"
	"net"

	"github.com/iamNilotpal/grpc/internal/hello"
	pb "github.com/iamNilotpal/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	srv := hello.NewService()

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})
	defer func() {
		log.Println("Shutting down GRPC server")
		grpcServer.GracefulStop()
		log.Println("Shutdown completed GRPC server")
	}()

	pb.RegisterHelloServiceServer(grpcServer, srv)

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
