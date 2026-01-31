package auth

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	serviceSecret  = []byte("13243dsf") // TODO: load from SERVICE_JWT_SECRET env
	userPrivateKey ed25519.PrivateKey   // TODO: load from USER_JWT_PRIVATE_KEY env
	userPublicKey  ed25519.PublicKey    // TODO: load from USER_JWT_PUBLIC_KEY env
)

var (
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

func newClaims(userID string, expiration time.Duration) jwt.RegisteredClaims {
	now := time.Now()
	return jwt.RegisteredClaims{
		Issuer:    "auth-service",
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(now.Add(expiration)),
		IssuedAt:  jwt.NewNumericDate(now),
	}
}

// NewJWTService creates a token for service-to-service communication (HS256)
func NewJWTService(userID string, expiration time.Duration) (string, error) {
	claims := newClaims(userID, expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(serviceSecret)
}

// NewJWTUser creates a token for user authentication (EdDSA)
func NewJWTUser(userID string, expiration time.Duration) (string, error) {
	if userPrivateKey == nil {
		return "", errors.New("user private key not configured")
	}
	claims := newClaims(userID, expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(userPrivateKey)
}

// ValidateJWT validates a token and returns its claims
func ValidateJWT(tokenString string, forServices bool) (*jwt.RegisteredClaims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if forServices {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidSigningMethod
			}
			return serviceSecret, nil
		}

		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return userPublicKey, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
