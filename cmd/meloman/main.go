package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/moguchev/meloman/db"

	"go.uber.org/zap"
)

func ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong\n")
}

func main() {
	url := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("create logger")
	}
	defer logger.Sync()

	database, err := db.Initialize(ctx, url, logger)
	if err != nil {
		logger.Fatal("init database", zap.Error(err))
	}
	defer database.Close()

	if err = db.Migrate(url); err != nil {
		logger.Fatal("migrate database", zap.Error(err))
	}

	logger.Info("Server started")

	http.HandleFunc("/ping", ping)
	if err = http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("listen and serve", zap.Error(err))
	}
}
