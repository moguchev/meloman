package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *implimentation) RemoveAlbum(ctx context.Context, req *meloman.RemoveAlbumRequest) (*emptypb.Empty, error) {
	const api = "service.RemoveAlbum"

	db := s.repo.DB()

	claims, ok := s.authManager.GetUserClaimsFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// check user access to collection
	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(1) FROM users WHERE id = $1 AND login = $2", userID, claims.Username).
		Scan(&count); err != nil || count == 0 {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	albumID, err := uuid.Parse(req.GetAlbumId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// add album to user collection
	tag, err := db.Exec(ctx, "DELETE FROM users_albums WHERE user_id = $1 AND album_id = $2", userID, albumID)
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &emptypb.Empty{}, nil
}
