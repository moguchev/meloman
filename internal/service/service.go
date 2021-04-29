package service

import (
	"github.com/moguchev/meloman/db"
	"github.com/moguchev/meloman/internal/auth"
	"github.com/moguchev/meloman/pkg/api/meloman"
	"go.uber.org/zap"
)

type implimentation struct {
	meloman.UnimplementedMelomanServer

	repo        db.Database
	log         *zap.Logger
	authManager auth.Manager
}

func NewService(log *zap.Logger, repo db.Database, am auth.Manager) meloman.MelomanServer {
	if log == nil {
		log, _ = zap.NewDevelopment()
	}
	return &implimentation{
		repo:        repo,
		log:         log,
		authManager: am,
	}
}
