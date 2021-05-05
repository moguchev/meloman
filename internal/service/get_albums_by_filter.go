package service

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/moguchev/meloman/pkg/api/meloman"
)

func (s *implimentation) GetAlbumsByFilter(ctx context.Context, req *meloman.GetAlbumsByFilterRequest) (*meloman.GetAlbumsByFilterResponse, error) {
	const api = "service.GetAlbumsByFilter"

	query := sq.Select("a.id", "a.title", "a.year", "a.cover",
		"f.id", "f.name", "l.id", "l.name", "art.id", "art.full_name").
		From("albums a").
		Join("formats f ON a.format_id = f.id").
		Join("labels l ON a.label_id = l.id").
		Join("artists art ON a.artist_id = art.id").
		Limit(100).PlaceholderFormat(sq.Dollar)

	if req.AlbumTitle != nil {
		query = query.Where(sq.Expr("a.title ILIKE ?", "%"+req.GetAlbumTitle()+"%"))
	}
	if req.ArtistName != nil {
		query = query.Where(sq.Expr("art.full_name ILIKE ?", "%"+req.GetArtistName()+"%"))
	}
	if req.TrackTitle != nil {
		query = query.Join("tracks t ON a.id = t.album_id").
			Where(sq.Expr("t.title ILIKE ?", "%"+req.GetTrackTitle()+"%"))
	}
	if req.Limit != nil {
		query = query.Limit(req.GetLimit())
	}
	if req.Offset != nil {
		query = query.Offset(req.GetOffset())
	}

	raw, args, err := query.ToSql()
	if err != nil {
		s.log.Sugar().Errorf("%s: to sql: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	db := s.repo.DB()
	rows, err := db.Query(ctx, raw, args...)
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get albums from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	albums := []*meloman.GetAlbumsByFilterResponse_Album{}
	for rows.Next() {
		var (
			albumID, artistID                 uuid.UUID
			albumTitle, format, label, artist string
			formatID, labelID                 int32
			albumCover                        sql.NullString
			albumYear                         sql.NullInt32
		)
		if err := rows.Scan(&albumID, &albumTitle, &albumYear, &albumCover,
			&formatID, &format, &labelID, &label, &artistID, &artist); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}

		albums = append(albums, &meloman.GetAlbumsByFilterResponse_Album{
			Id:    albumID.String(),
			Title: albumTitle,
			Year:  albumYear.Int32,
			Image: albumCover.String,
			Format: &meloman.Format{
				Id:   formatID,
				Name: format,
			},
			Label: &meloman.Label{
				Id:   labelID,
				Name: label,
			},
			Artist: &meloman.GetAlbumsByFilterResponse_Album_Artist{
				Id:   artistID.String(),
				Name: artist,
			},
		})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetAlbumsByFilterResponse{Albums: albums}, nil
}
