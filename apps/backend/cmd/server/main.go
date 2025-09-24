package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/SquadGO/squad-admin-panel/internal/core"
	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/http/router"
	"github.com/SquadGO/squad-admin-panel/internal/http/server"
	"github.com/SquadGO/squad-admin-panel/internal/logger"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	"github.com/SquadGO/squad-admin-panel/internal/state"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
)

func main() {
	logger.New()

	db, err := db.New(context.Background())
	if err != nil {
		slog.Error("Database init", slog.Any("err", err))
		return
	}

	goth.UseProviders(
		steam.New(os.Getenv("STEAM_KEY"), os.Getenv("CALLBACK_URL")),
	)

	appState := &state.AppState{}

	s := service.NewService(db, appState)
	r := router.New(s)

	go core.NewRcon(s, appState)
	go core.NewFTPLogs(s, appState)

	server.New(r)
}
