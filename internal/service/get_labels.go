package service

import (
	"context"

	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *implimentation) GetLabels(ctx context.Context, _ *emptypb.Empty) (*meloman.GetLabelsResponse, error) {
	const api = "service.GetLabels"

	db := s.repo.DB()

	rows, err := db.Query(ctx, "SELECT id, name FROM labels")
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get labels from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	labels := []*meloman.Label{}
	for rows.Next() {
		var (
			id   int32
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		labels = append(labels, &meloman.Label{Id: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetLabelsResponse{Labels: labels}, nil
}
