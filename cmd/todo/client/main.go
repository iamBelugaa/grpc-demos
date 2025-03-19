package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Client conn error : %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Client conn close error : %v", err)
		}
	}()

	client := pb.NewTodoServiceClient(conn)

	client.AddTodo(context.Background(), &pb.AddTodoRequest{Title: "Task #1"})
	client.AddTodo(context.Background(), &pb.AddTodoRequest{Title: "Task #2"})
	client.AddTodo(context.Background(), &pb.AddTodoRequest{Title: "Task #3"})
	client.AddTodo(context.Background(), &pb.AddTodoRequest{Title: "Task #4"})

	resp, err := client.AddTodo(context.Background(), &pb.AddTodoRequest{Title: "Task #5"})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			log.Fatalf("%v", err)
		} else {
			log.Fatalf("%v", s)
		}
	}

	done := true
	if _, err := client.UpdateTodo(context.Background(), &pb.UpdateTodoRequest{Id: resp.Id, Done: &done}); err != nil {
		s, ok := status.FromError(err)
		if !ok {
			log.Fatalf("%v", err)
		} else {
			log.Fatalf("%v", s)
		}
	}

	if resp, err := client.ListTodos(context.Background(), &pb.ListTodoRequest{}); err != nil {
		s, ok := status.FromError(err)
		if !ok {
			log.Fatalf("%v", err)
		} else {
			log.Fatalf("%v", s)
		}
	} else {
		for i, todo := range resp.Todos {
			if i == 0 {
				println("---------------------------")
			}
			fmt.Printf("ID : %d\n", todo.Id)
			fmt.Printf("Title : %s\n", todo.Title)
			fmt.Printf("Done : %v\n", todo.Done)
			fmt.Printf("Created At : %+v\n", time.Unix(0, todo.CreatedAt).Local().String())
			println("---------------------------")
		}
	}

	if _, err := client.DeleteTodo(context.Background(), &pb.DeleteTodoRequest{Id: resp.Id}); err != nil {
		s, ok := status.FromError(err)
		if !ok {
			log.Fatalf("%v", err)
		} else {
			log.Fatalf("%v", s)
		}
	}
}
