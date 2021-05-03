package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) CreateTrack(ctx context.Context, req *meloman.CreateTrackRequest) (*meloman.CreateTrackResponse, error) {
	const api = "service.CreateTrack"

	albumID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	newTrack := req.GetTrack()
	if newTrack == nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	id := uuid.New()

	tag, err := s.repo.DB().Exec(ctx,
		"INSERT INTO tracks (id, number, album_id, title, length) VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING",
		id, newTrack.GetNumber(), albumID, newTrack.GetTitle(), newTrack.GetLenght())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	return &meloman.CreateTrackResponse{Id: id.String()}, nil
}
