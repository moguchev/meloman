package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) CreateArtist(ctx context.Context, req *meloman.CreateArtistRequest) (*meloman.CreateArtistsResponse, error) {
	const api = "service.CreateArtist"

	if req.GetFullName() == "" {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	id := uuid.New()

	tag, err := s.repo.DB().Exec(ctx,
		"INSERT INTO artists (id, full_name, biography, image) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING",
		id, req.GetFullName(), req.GetBiography(), req.GetImage())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	return &meloman.CreateArtistsResponse{Id: id.String()}, nil
}
