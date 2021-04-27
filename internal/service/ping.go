package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	const api = "service.ping"
	if err := s.repo.DB().Ping(ctx); err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	return &emptypb.Empty{}, nil
}
