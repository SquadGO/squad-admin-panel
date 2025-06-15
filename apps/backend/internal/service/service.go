package service

import (
	"github.com/SquadGO/squad-admin-panel/internal/db"
)

type Service struct {
	UserService   UserService
	PlayerService PlayerService
	ServerService ServerService
	RconService   RconService
	LogsService   LogsService
}

func NewService(db *db.Storage) *Service {
	userService := NewUserService(db)
	playerService := NewPlayerService(db)
	serverService := NewServerService(db)
	rconService := NewRconService(db, playerService)
	logsService := NewLogsService(db, playerService)

	return &Service{
		UserService:   userService,
		PlayerService: playerService,
		ServerService: serverService,
		RconService:   rconService,
		LogsService:   logsService,
	}
}
