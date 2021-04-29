package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/moguchev/meloman/internal/auth"
	"github.com/moguchev/meloman/internal/models"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *implimentation) Auth(ctx context.Context, req *meloman.AuthRequest) (*meloman.AuthResponse, error) {
	const api = "service.GetUser"

	// validate input
	if err := validateCredentials(req.GetCredentials()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db := s.repo.DB()
	var (
		id               uuid.UUID
		hash, salt, role string
	)

	row := db.QueryRow(ctx, "SELECT id, password, salt, role FROM users WHERE login = $1", req.GetCredentials().GetLogin())
	if err := row.Scan(&id, &hash, &salt, &role); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, codes.NotFound.String())
		}
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	hashPassword := sha256.Sum256([]byte(req.GetCredentials().GetPassword() + salt))
	if fmt.Sprintf("%x", hashPassword[:]) != hash {
		return nil, status.Errorf(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	// set auth headers
	token, err := s.authManager.Generate(req.GetCredentials().GetLogin(), models.ParseRole(role).String())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	// create and send header
	header := metadata.Pairs(auth.Header, token)
	grpc.SendHeader(ctx, header)
	// create and set trailer
	trailer := metadata.Pairs(auth.Header, token)
	grpc.SetTrailer(ctx, trailer)

	return &meloman.AuthResponse{Id: id.String()}, nil
}
