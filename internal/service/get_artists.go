package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) GetArtists(ctx context.Context, req *meloman.GetArtistsRequest) (*meloman.GetArtistsResponse, error) {
	const api = "service.GetArtists"

	db := s.repo.DB()

	// ______________artists________________
	// id         uuid          NOT NULL,
	// full_name  varchar(256)  NOT NULL,
	// biography  text          NOT NULL DEFAULT '',
	// image      text,
	var (
		limit uint32 = 100
		count uint32
	)

	if err := db.QueryRow(ctx, "SELECT COUNT(1) FROM artists").Scan(&count); err != nil {
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	if count == 0 {
		return &meloman.GetArtistsResponse{Total: count, Artists: []*meloman.GetArtistsResponse_Artist{}}, nil
	}

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	rows, err := db.Query(ctx, "SELECT id, full_name, image FROM artists LIMIT $1 OFFSET $2", limit, req.GetOffset())
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get artists from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	artists := []*meloman.GetArtistsResponse_Artist{}
	for rows.Next() {
		var (
			id       uuid.UUID
			fullName string
			image    sql.NullString
		)
		if err := rows.Scan(&id, &fullName, &image); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		artists = append(artists, &meloman.GetArtistsResponse_Artist{
			Id:       id.String(),
			FullName: fullName,
			Image:    image.String,
		})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetArtistsResponse{
		Total:   count,
		Artists: artists,
	}, nil
}
