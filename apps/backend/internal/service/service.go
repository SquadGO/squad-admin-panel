package service

import (
	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/state"
)

type Service struct {
	UserService   UserService
	PlayerService PlayerService
	ServerService ServerService
	MapService    MapService
	RconService   RconService
	LogsService   LogsService
}

func NewService(db *db.Storage, appState *state.AppState) *Service {
	userService := NewUserService(db)
	playerService := NewPlayerService(db)
	serverService := NewServerService(db)
	mapService := NewMapService(db)
	rconService := NewRconService(db, playerService)
	logsService := NewLogsService(db, appState, playerService, userService, mapService)

	return &Service{
		UserService:   userService,
		PlayerService: playerService,
		ServerService: serverService,
		MapService:    mapService,
		RconService:   rconService,
		LogsService:   logsService,
	}
}
