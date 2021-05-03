package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/moguchev/meloman/internal/models"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *implimentation) GetUsers(ctx context.Context, req *meloman.GetUsersRequest) (*meloman.GetUsersResponse, error) {
	const api = "service.GetUsers"

	db := s.repo.DB()

	// ______________users________________
	// id         uuid          NOT NULL,
	// login      varchar(64)   NOT NULL,
	// password   varchar(256)  NOT NULL,
	// salt       varchar(256)  NOT NULL,
	// created_at timestamp     NOT NULL,
	// updated_at timestamp     NOT NULL,
	// role   	  user_role     NOT NULL DEFAULT 'user',
	var (
		limit uint32 = 100
		count uint32
	)

	if err := db.QueryRow(ctx, "SELECT COUNT(1) FROM users").Scan(&count); err != nil {
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	if count == 0 {
		return &meloman.GetUsersResponse{Total: count, Users: []*meloman.User{}}, nil
	}

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	rows, err := db.Query(ctx, "SELECT id, login, created_at, updated_at, role FROM users LIMIT $1 OFFSET $2", limit, req.GetOffset())
	if err != nil {
		s.log.Sugar().Errorf("%s: can't get users from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	defer rows.Close()

	users := []*meloman.User{}
	for rows.Next() {
		var (
			id                   uuid.UUID
			login, role          string
			createdAt, updatedAt time.Time
		)
		if err := rows.Scan(&id, &login, &createdAt, &updatedAt, &role); err != nil {
			s.log.Sugar().Errorf("%s: scan: %s", api, err.Error())
			return nil, status.Errorf(codes.Internal, codes.Internal.String())
		}
		users = append(users, &meloman.User{
			Login:     login,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
			Role:      models.ParseRole(role).Proto(),
			Id:        id.String(),
		})
	}

	if err := rows.Err(); err != nil {
		s.log.Sugar().Errorf("%s: error from rows: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetUsersResponse{
		Total: count,
		Users: users,
	}, nil
}
