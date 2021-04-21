package service

import (
	"github.com/moguchev/meloman/db"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
)

type service struct {
	meloman.UnimplementedMelomanServer

	repo db.Database
	log  *zap.Logger
}

func NewService(log *zap.Logger, repo db.Database) meloman.MelomanServer {
	if log == nil {
		log, _ = zap.NewDevelopment()
	}
	return &service{
		repo: repo,
		log:  log,
	}
}
