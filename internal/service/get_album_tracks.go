package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) GetAlbumTracks(ctx context.Context, req *meloman.GetAlbumTracksRequest) (*meloman.GetAlbumTracksResponse, error) {
	const api = "service.GetAlbumTracks"

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	db := s.repo.DB()

	rows, err := db.Query(ctx, `SELECT id, number, title, length FROM tracks 
		WHERE album_id = $1 ORDER BY number`, id)
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get album tracks from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	tracks := []*meloman.GetAlbumTracksResponse_Track{}
	for rows.Next() {
		var (
			id            uuid.UUID
			number        int32
			title, length string
		)
		if err := rows.Scan(&id, &number, &title, &length); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		tracks = append(tracks, &meloman.GetAlbumTracksResponse_Track{
			Id:     id.String(),
			Title:  title,
			Number: number,
			Lenght: length,
		})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetAlbumTracksResponse{Tracks: tracks}, nil
}
