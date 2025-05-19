package service

import (
	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type PlayerService interface{}

type playerService struct {
	db *db.Storage
}

func NewPlayerService(db *db.Storage) PlayerService {
	return playerService{
		db: db,
	}
}

func (p *playerService) CreatePlayer(player models.CreatePlayer) {

}
