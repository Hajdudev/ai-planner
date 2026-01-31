package service

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/Hajdudev/ai-planner/services/auth-service/internal/domain"
	"github.com/Hajdudev/ai-planner/services/auth-service/pkg/types"
	"github.com/Hajdudev/ai-planner/shared/auth"
)

var hashSecret = make([]byte, 32)

type AuthService struct {
	repo domain.AuthRepository
}

func (s *AuthService) CreateTokens(userID string) (*types.TokenPairs, error) {
	ctx := context.Background()
	accessToken, err := auth.NewJWTUser(userID, 15*time.Minute)
	if err != nil {
		return &types.TokenPairs{}, err
	}

	refreshToken, err := s.CreateRefreshToken(ctx)
	if err != nil {
		return &types.TokenPairs{}, err
	}

	return &types.TokenPairs{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) CreateRefreshToken(ctx context.Context) (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	refreshToken := base64.URLEncoding.EncodeToString(bytes)

	// Hash the token before storing
	hmacHasher := hmac.New(sha256.New, hashSecret)
	hmacHasher.Write([]byte(refreshToken))
	hashedToken := base64.URLEncoding.EncodeToString(hmacHasher.Sum(nil))

	err = s.repo.StoreRefreshToken(ctx, domain.RefreshToken(hashedToken))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
