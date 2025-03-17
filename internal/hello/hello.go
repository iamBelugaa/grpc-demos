package hello

import (
	"context"
	"math/rand/v2"

	pb "github.com/iamNilotpal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedHelloServiceServer
}

func (service) SayHello(context context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	if rand.IntN(10) < 5 {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return &pb.SayHelloResponse{Message: req.FirstName + " " + req.LastName}, nil
}

func NewService() *service {
	return &service{}
}
