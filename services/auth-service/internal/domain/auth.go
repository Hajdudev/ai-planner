package domain

import (
	"context"

	"github.com/Hajdudev/ai-planner/services/auth-service/pkg/types"
)

type (
	RefreshToken string
	AccessToken  string
	UserID       string
)

type AuthRepository interface {
	StoreRefreshToken(ctx context.Context, refreshToken RefreshToken) error
	GetRefreshToken(ctx context.Context, userID UserID) (RefreshToken, error)
}

type AuthService interface {
	CreateTokens(userID UserID) (*types.TokenPairs, error)
	GetAccessTokenFromRefreshToken(ctx context.Context, userID, refreshToken string) (string, error)
}

type UserService interface {
	Login(ctx context.Context, email, passwordHash string) (*types.TokenPairs, error)
	Signup(ctx context.Context, email, username, passwordHash string) (verificationID string, expiresIn int32, err error)
	VerifyEmail(ctx context.Context, verificationID, code string) (*types.TokenPairs, error)
	ForgotPassword(ctx context.Context, email string) (resetID string, expiresIn int32, err error)
	VerifyPasswordReset(ctx context.Context, resetID, code, passwordHash string) (*types.TokenPairs, error)
}
