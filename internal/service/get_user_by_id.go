package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/moguchev/meloman/internal/models"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *implimentation) GetUserByID(ctx context.Context, req *meloman.GetUserByIDRequest) (*meloman.GetUserByIDResponse, error) {
	const api = "service.GetUserByID"

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
		login     string
		createdAt time.Time
		updatedAt time.Time
		role      string
		id        uuid.UUID
	)

	row := db.QueryRow(ctx, "SELECT id, login, created_at, updated_at, role FROM users WHERE id = $1", req.GetId())
	if err := row.Scan(&id, &login, &createdAt, &updatedAt, &role); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, codes.NotFound.String())
		}
		s.log.Sugar().Errorf("%s: can't get user from db: %s", api, err.Error())
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return &meloman.GetUserByIDResponse{
		User: &meloman.User{
			Login:     login,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
			Role:      models.ParseRole(role).Proto(),
			Id:        id.String(),
		},
	}, nil
}
