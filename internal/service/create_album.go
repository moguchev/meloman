package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) CreateAlbum(ctx context.Context, req *meloman.CreateAlbumRequest) (*meloman.CreateAlbumResponse, error) {
	const api = "service.CreateAlbum"

	artistID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	newAlbum := req.GetAlbum()
	if newAlbum == nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	id := uuid.New()

	tag, err := s.repo.DB().Exec(ctx,
		"INSERT INTO albums (id, title, artist_id, year, cover, format_id, label_id) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT DO NOTHING",
		id, newAlbum.GetTitle(), artistID, newAlbum.GetYear(), newAlbum.GetImage(), newAlbum.GetFormatId(), newAlbum.GetLabelId())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	return &meloman.CreateAlbumResponse{Id: id.String()}, nil
}
