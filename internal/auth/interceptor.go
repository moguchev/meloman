package auth

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const Header = "authorization"

type Interceptor interface {
	Unary() grpc.UnaryServerInterceptor
	Stream() grpc.StreamServerInterceptor
}

type interceptor struct {
	tokenManager    TokenManager
	accessibleRoles map[string][]string
	logger          *zap.Logger
}

func NewInterceptor(manager TokenManager, accessibleRoles map[string][]string, log *zap.Logger) Interceptor {
	if log == nil {
		log, _ = zap.NewDevelopment()
	}
	return &interceptor{
		tokenManager:    manager,
		accessibleRoles: accessibleRoles,
		logger:          log,
	}
}

func (i *interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		const api = "interceptor.Unary"
		i.logger.Debug(api, zap.String("method", info.FullMethod))

		if err := i.authorize(ctx, info.FullMethod); err != nil {
			i.logger.Error(api, zap.Error(err))
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *interceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		const api = "interceptor.Stream"
		i.logger.Debug(api, zap.String("method", info.FullMethod))

		if err := i.authorize(stream.Context(), info.FullMethod); err != nil {
			i.logger.Error(api, zap.Error(err))
			return err
		}

		return handler(srv, stream)
	}
}

func (i *interceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := i.accessibleRoles[method]
	if !ok || len(accessibleRoles) == 0 {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md[Header]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := i.tokenManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
