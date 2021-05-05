package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *implimentation) GetArtistAlbums(ctx context.Context, req *meloman.GetArtistAlbumsRequest) (*meloman.GetArtistAlbumsResponse, error) {
	const api = "service.GetArtistAlbums"

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	db := s.repo.DB()

	rows, err := db.Query(ctx, `SELECT a.id, a.title, a.year, a.cover, a.format_id, a.label_id, f.name, l.name
		FROM albums a
		JOIN formats f ON a.format_id = f.id
		JOIN labels l ON a.label_id = l.id
		WHERE a.artist_id = $1
		ORDER BY a.year`, id)
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get artist albums from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	albums := []*meloman.GetArtistAlbumsResponse_Album{}
	for rows.Next() {
		var (
			id                   uuid.UUID
			title, label, format string
			year                 sql.NullInt32
			cover                sql.NullString
			formatID, labelID    int32
		)
		if err := rows.Scan(&id, &title, &year, &cover, &formatID, &labelID, &format, &label); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		albums = append(albums, &meloman.GetArtistAlbumsResponse_Album{
			Id:    id.String(),
			Title: title,
			Year:  year.Int32,
			Image: cover.String,
			Format: &meloman.Format{
				Id:   formatID,
				Name: format,
			},
			Label: &meloman.Label{
				Id:   labelID,
				Name: label,
			},
		})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetArtistAlbumsResponse{Albums: albums}, nil
}
