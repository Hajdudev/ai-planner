package grpc

import (
	"context"

	"github.com/Hajdudev/ai-planner/services/auth-service/internal/domain"
	"google.golang.org/grpc"

	pb "github.com/Hajdudev/ai-planner/shared/proto/auth"
)

type gRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	authService domain.AuthService
	userService domain.UserService
}

func NewGRPCHandler(server *grpc.Server, authService domain.AuthService, userService domain.UserService) *gRPCHandler {
	handler := &gRPCHandler{
		authService: authService,
		userService: userService,
	}

	pb.RegisterAuthServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) GetAccessTokenFromRefreshToken(ctx context.Context, req *pb.GetAccessTokenFromRefreshTokenRequest) (*pb.GetAccessTokenFromRefreshTokenResponse, error) {
	// TODO: implement
	return &pb.GetAccessTokenFromRefreshTokenResponse{}, nil
}

func (h *gRPCHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokens, err := h.userService.Login(ctx, req.Email, req.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Tokens: &pb.AuthTokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	}, nil
}

func (h *gRPCHandler) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	verificationID, expiresIn, err := h.userService.Signup(ctx, req.Email, req.Username, req.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &pb.SignupResponse{
		VerificationId: verificationID,
		ExpiresIn:      expiresIn,
	}, nil
}

func (h *gRPCHandler) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	tokens, err := h.userService.VerifyEmail(ctx, req.VerificationId, req.Code)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyEmailResponse{
		Tokens: &pb.AuthTokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	}, nil
}

func (h *gRPCHandler) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	resetID, expiresIn, err := h.userService.ForgotPassword(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.ForgotPasswordResponse{
		ResetId:   resetID,
		ExpiresIn: expiresIn,
	}, nil
}

func (h *gRPCHandler) VerifyPasswordReset(ctx context.Context, req *pb.VerifyPasswordResetRequest) (*pb.VerifyPasswordResetResponse, error) {
	tokens, err := h.userService.VerifyPasswordReset(ctx, req.ResetId, req.Code, req.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyPasswordResetResponse{
		Tokens: &pb.AuthTokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	}, nil
}
