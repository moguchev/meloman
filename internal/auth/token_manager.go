package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	Generate(user, role string) (string, error)
	Verify(accessToken string) (*UserClaims, error)
	GetUserClaimsFromContext(ctx context.Context) (*UserClaims, bool)
	PutUserClaimsToContext(ctx context.Context, claims *UserClaims) context.Context
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

type jwtManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type claimsKey struct {
	string
}

var claimsKeykey = claimsKey{"key"}

func NewJWTManager(secretKey string, tokenDuration time.Duration) TokenManager {
	return &jwtManager{secretKey, tokenDuration}
}

func (manager *jwtManager) Generate(user, role string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: user,
		Role:     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *jwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (manager *jwtManager) GetUserClaimsFromContext(ctx context.Context) (*UserClaims, bool) {
	value := ctx.Value(claimsKeykey)

	claims, ok := value.(UserClaims)
	if !ok {
		return nil, false
	}

	return &claims, true
}

func (manager *jwtManager) PutUserClaimsToContext(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, claimsKeykey, *claims)
}
