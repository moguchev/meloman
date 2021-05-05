package service

import (
	"context"
	"errors"

	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) GetArtistByID(ctx context.Context, req *meloman.GetArtistByIDRequest) (*meloman.GetArtistByIDResponse, error) {
	const api = "service.GetArtistByID"

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	db := s.repo.DB()
	// ______________artists________________
	// id         uuid          NOT NULL,
	// full_name  varchar(256)  NOT NULL,
	// biography  text          NOT NULL DEFAULT '',
	// image      text,
	var (
		fullName, biography string
		image               sql.NullString
	)

	row := db.QueryRow(ctx, "SELECT full_name, biography, image FROM artists WHERE id = $1", id)
	if err := row.Scan(&fullName, &biography, &image); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, codes.NotFound.String())
		}
		s.log.Sugar().Errorf("%s: can't get artist from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetArtistByIDResponse{
		Id:        id.String(),
		FullName:  fullName,
		Biography: biography,
		Image:     image.String,
	}, nil
}
