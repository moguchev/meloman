package service

import (
	"context"

	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *implimentation) GetFormats(ctx context.Context, _ *emptypb.Empty) (*meloman.GetFormatsResponse, error) {
	const api = "service.GetFormats"

	db := s.repo.DB()

	rows, err := db.Query(ctx, "SELECT id, name FROM formats")
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get formats from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	formats := []*meloman.Format{}
	for rows.Next() {
		var (
			id   int32
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		formats = append(formats, &meloman.Format{Id: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetFormatsResponse{Formats: formats}, nil
}
