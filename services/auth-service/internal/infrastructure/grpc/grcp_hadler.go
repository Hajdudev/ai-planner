package grpc

import (
	"github.com/Hajdudev/ai-planner/services/auth-service/internal/domain"
	"google.golang.org/grpc"

	pb "github.com/Hajdudev/ai-planner/shared/proto/auth"
)

type gRPCHandler struct {
	service domain.AuthService
}

func NewGRPCHandler(server *grpc.Server, service domain.AuthService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterAuthServiceServer(server, handler)
	return handler
}
