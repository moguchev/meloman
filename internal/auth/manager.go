package auth

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Manager interface {
	TokenManager
	Interceptor
}

type manager struct {
	interceptor  Interceptor
	tokenManager TokenManager
}

func NewManager(tm TokenManager, accessibleRoles map[string][]string, log *zap.Logger) Manager {
	if log == nil {
		log, _ = zap.NewDevelopment()
	}
	return &manager{
		interceptor:  NewInterceptor(tm, accessibleRoles, log),
		tokenManager: tm,
	}
}

func (m *manager) Generate(user, role string) (string, error) {
	return m.tokenManager.Generate(user, role)
}

func (m *manager) Verify(accessToken string) (*UserClaims, error) {
	return m.tokenManager.Verify(accessToken)
}

func (m *manager) Unary() grpc.UnaryServerInterceptor {
	return m.interceptor.Unary()
}

func (m *manager) Stream() grpc.StreamServerInterceptor {
	return m.interceptor.Stream()
}

func (m *manager) GetUserClaimsFromContext(ctx context.Context) (*UserClaims, bool) {
	return m.tokenManager.GetUserClaimsFromContext(ctx)
}

func (m *manager) PutUserClaimsToContext(ctx context.Context, claims *UserClaims) context.Context {
	return m.tokenManager.PutUserClaimsToContext(ctx, claims)
}
