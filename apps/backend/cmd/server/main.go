package main

import (
	"context"
	"log/slog"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/http/router"
	"github.com/SquadGO/squad-admin-panel/internal/http/server"
	"github.com/SquadGO/squad-admin-panel/internal/logger"
	"github.com/SquadGO/squad-admin-panel/internal/service"
)

func main() {
	logger.New()

	db, err := db.New(context.Background())
	if err != nil {
		slog.Error("Database", slog.Any("err", err))
		return
	}

	repo := service.NewService(db)
	r := router.New(repo)

	server.New(r)
}
