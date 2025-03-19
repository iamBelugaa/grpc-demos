package hello

import (
	"context"
	"math/rand/v2"
	"strings"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedHelloServiceServer
}

func (service) SayHello(context context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	if strings.EqualFold(strings.TrimSpace(req.FirstName), "") {
		return nil, status.Errorf(codes.InvalidArgument, "FirstName is required")
	}

	if strings.EqualFold(strings.TrimSpace(req.LastName), "") {
		return nil, status.Errorf(codes.InvalidArgument, "LastName is required")
	}

	if rand.IntN(10) < 5 {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &pb.SayHelloResponse{Message: req.FirstName + " " + req.LastName}, nil
}

func NewService() *service {
	return &service{}
}
