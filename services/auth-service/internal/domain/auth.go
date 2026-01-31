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
}
