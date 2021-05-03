package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/moguchev/meloman/internal/auth"
	"github.com/moguchev/meloman/internal/models"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"github.com/moguchev/meloman/pkg/random"
)

func (s *implimentation) CreateUser(ctx context.Context, req *meloman.CreateUserRequest) (*meloman.CreateUserResponse, error) {
	const api = "service.CreateUser"

	credentials := req.GetCredentials()

	// validate input
	if err := validateCredentials(credentials); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db := s.repo.DB()

	// check if exists
	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(1) FROM users WHERE login = $1", credentials.GetLogin()).Scan(&count); err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if count > 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	// hash password
	salt, err := random.GenerateRandomString(32)
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	hashPassword := sha256.Sum256([]byte(credentials.GetPassword() + salt))
	id := uuid.New()

	// create user in db
	tag, err := db.Exec(ctx,
		"INSERT INTO users (id, login, password, salt, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT DO NOTHING",
		id, credentials.GetLogin(), fmt.Sprintf("%x", hashPassword[:]), salt, time.Now(), time.Now())
	if err != nil {
		s.log.Error(api, zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	if tag.RowsAffected() == 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	// set auth headers
	token, err := s.authManager.Generate(credentials.GetLogin(), models.RoleUser.String())
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

	return &meloman.CreateUserResponse{Id: id.String()}, nil
}

func validateCredentials(c *meloman.Credentials) error {
	password := c.GetPassword()
	if len(password) == 0 {
		return errors.New("invalid password")
	}

	login := c.GetLogin()
	if len(login) == 0 {
		return errors.New("invalid login")
	}
	return nil
}
