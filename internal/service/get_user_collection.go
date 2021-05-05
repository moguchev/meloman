package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *implimentation) GetUserCollection(ctx context.Context,
	req *meloman.GetUserCollectionRequest) (*meloman.GetUserCollectionResponse, error) {
	const api = "service.GetUserCollection"

	db := s.repo.DB()

	claims, ok := s.authManager.GetUserClaimsFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// check user access to collection
	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(1) FROM users WHERE id = $1 AND login = $2", userID, claims.Username).
		Scan(&count); err != nil || count == 0 {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	var (
		stat                  = new(meloman.GetUserCollectionResponse_Statistic)
		lastAdded, firstAdded time.Time
	)

	// get statistic
	if err := db.QueryRow(ctx, `SELECT
		COUNT(DISTINCT a.id) AS albums_num,
		COUNT(DISTINCT a2.id) AS artist_num,
		COUNT(t.id) AS tracks_num,
		MAX(added_at) AS last_added,
		MIN(added_at) AS first_added
	FROM users_albums
	JOIN albums a on users_albums.album_id = a.id
	JOIN artists a2 on a.artist_id = a2.id
	JOIN tracks t on a.id = t.album_id
	WHERE user_id = $1
	GROUP BY user_id`, userID).Scan(
		&stat.AlbumNum, &stat.ArtistNum, &stat.TrackNum, &lastAdded, &firstAdded,
	); err != nil {
		s.log.Sugar().Errorf("%s: can't get stat: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	stat.LastAdded = timestamppb.New(lastAdded)
	stat.FirstAdded = timestamppb.New(firstAdded)
	// get albums
	rows, err := db.Query(ctx, `SELECT 
		a.id, a.title, a.year, a.cover,
		f.id, f.name, l.id, l.name,
		a2.id, a2.full_name
	FROM users_albums
	JOIN albums a on users_albums.album_id = a.id
	JOIN artists a2 on a.artist_id = a2.id
	JOIN tracks t on a.id = t.album_id
	JOIN formats f on a.format_id = f.id
	JOIN labels l on a.label_id = l.id
	WHERE user_id = $1`, userID)
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get albums from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	albums := []*meloman.GetUserCollectionResponse_Album{}
	for rows.Next() {
		album := new(meloman.GetUserCollectionResponse_Album)
		format := new(meloman.Format)
		label := new(meloman.Label)
		artist := new(meloman.GetUserCollectionResponse_Album_Artist)

		if err = rows.Scan(&album.Id, &album.Title, &album.Year, &album.Title,
			&format.Id, &format.Name, &label.Id, &label.Name, &artist.Id, &artist.Name); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		album.Format = format
		album.Artist = artist
		album.Label = label
		albums = append(albums, album)
	}

	return &meloman.GetUserCollectionResponse{Statistic: stat, Artists: nil, Albums: albums}, nil
}
