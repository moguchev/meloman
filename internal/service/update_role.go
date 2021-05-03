package service

import (
	"context"

	"github.com/moguchev/meloman/internal/models"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *implimentation) UpdateUserRole(ctx context.Context, req *meloman.UpdateUserRoleRequest) (*emptypb.Empty, error) {
	const api = "service.UpdateUserRole"

	role := models.ParseRole(req.GetRole().String())
	tag, err := s.repo.DB().Exec(ctx, "UPDATE users SET role = $1 WHERE id = $2", role, req.GetId())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &emptypb.Empty{}, nil
}
