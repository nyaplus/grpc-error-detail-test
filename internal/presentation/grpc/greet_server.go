package grpc

import (
	"context"
	"fmt"

	greetv1 "github.com/2yanpath/grpc-error-detail-test/proto/greet/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGreetServer() *greetServer {
	return &greetServer{}
}

type greetServer struct {
	greetv1.UnimplementedGreetServiceServer
}

// grpcurl -plaintext -d '{"name": "Jane"}' localhost:8081 greet.v1.GreetService.Greet
func (s *greetServer) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {

	if req.Name == "" {
		badRequestDetail := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "username",
					Description: "should not empty",
				},
			},
		}
		stat := status.New(codes.InvalidArgument, "name is required")
		stat, _ = stat.WithDetails(badRequestDetail)
		return nil, stat.Err()
	}

	return &greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
