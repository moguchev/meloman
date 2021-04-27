package auth

import "google.golang.org/grpc"

type Manager interface {
	TokenManager
	Interceptor
}

type manager struct {
	interceptor  Interceptor
	tokenManager TokenManager
}

func NewManager(intcptr Interceptor, tm TokenManager) Manager {
	return &manager{
		interceptor:  intcptr,
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
