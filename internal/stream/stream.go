package stream

import (
	"time"

	pb "github.com/iamNilotpal/grpc/proto/__generated__"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service struct {
	pb.UnimplementedStreamTimeServiceServer
}

func NewService() *service {
	return &service{}
}

func (service) StreamServerTime(req *pb.StreamTimeRequest, stream grpc.ServerStreamingServer[pb.StreamTimeResponse]) error {
	if req.IntervalSeconds <= 0 {
		return status.Errorf(codes.InvalidArgument, "IntervalSeconds must be positive and greater than 0")
	}

	duration := time.Duration(req.IntervalSeconds) * time.Second
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			{
				return stream.Context().Err()
			}
		case currentTime := <-ticker.C:
			{
				if err := stream.Send(&pb.StreamTimeResponse{CurrentTime: timestamppb.New(currentTime)}); err != nil {
					return err
				}
			}
		}
	}
}
