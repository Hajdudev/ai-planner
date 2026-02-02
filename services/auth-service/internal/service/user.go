package service

import (
	"context"

	"github.com/Hajdudev/ai-planner/services/auth-service/pkg/types"
)

func (s *AuthService) Login(ctx context.Context, email, passwordHash string) (*types.TokenPairs, error) {
	// TODO: implement
	return nil, nil
}

func (s *AuthService) Signup(ctx context.Context, email, username, passwordHash string) (verificationID string, expiresIn int32, err error) {
	// TODO: implement
	return "", 0, nil
}

func (s *AuthService) VerifyEmail(ctx context.Context, verificationID, code string) (*types.TokenPairs, error) {
	// TODO: implement
	return nil, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, email string) (resetID string, expiresIn int32, err error) {
	// TODO: implement
	return "", 0, nil
}

func (s *AuthService) VerifyPasswordReset(ctx context.Context, resetID, code, passwordHash string) (*types.TokenPairs, error) {
	// TODO: implement
	return nil, nil
}
