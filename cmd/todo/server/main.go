package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/iamNilotpal/grpc/internal/todo"
	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
)

func main() {
	shutdownCtx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()

	server := todo.NewService()
	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("TCP Error : %v", err)
	}

	serverError := make(chan error)
	go func() {
		log.Println("Starting GRPC Server at http://localhost:50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("GRPC Error : %v", err)
			serverError <- err
			close(serverError)
		}
	}()

	select {
	case <-shutdownCtx.Done():
		{
			log.Println("Stopping GRPC Server Gracefully...")
			grpcServer.GracefulStop()
			log.Println("GRPC Server Closed.")
		}
	case err := <-serverError:
		if err != nil {
			log.Printf("GRPC Server Error : %v\n", err)
		}
	}
}
