package todo

import (
	"context"
	"math/rand/v2"
	"strings"
	"sync"
	"time"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	mu    sync.RWMutex
	todos map[int64]*pb.Todo
	pb.UnimplementedTodoServiceServer
}

func NewService() *service {
	return &service{
		mu:    sync.RWMutex{},
		todos: make(map[int64]*pb.Todo),
	}
}

func (s *service) AddTodo(context context.Context, req *pb.AddTodoRequest) (*pb.AddTodoResponse, error) {
	newTodo := &pb.Todo{
		Id:        rand.Int64(),
		Title:     strings.TrimSpace(req.Title),
		Done:      false,
		CreatedAt: time.Now().UnixNano(),
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.todos[newTodo.Id] = newTodo
	return &pb.AddTodoResponse{
			Id:        newTodo.Id,
			Title:     newTodo.Title,
			Done:      newTodo.Done,
			CreatedAt: newTodo.CreatedAt,
		},
		nil
}

func (s *service) UpdateTodo(context context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	s.mu.RLock()
	todo, ok := s.todos[req.Id]
	s.mu.RUnlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "Todo with this id doesn't exists")
	}

	if req.Done != nil {
		todo.Done = *req.Done
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.todos[req.Id] = todo

	return &pb.UpdateTodoResponse{Success: true}, nil
}

func (s *service) DeleteTodo(context context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	s.mu.RLock()
	_, ok := s.todos[req.Id]
	s.mu.RUnlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "Todo with this id doesn't exists")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.todos, req.Id)

	return &pb.DeleteTodoResponse{}, nil
}

func (s *service) ListTodos(context context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*pb.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return &pb.ListTodoResponse{Todos: todos}, nil
}
