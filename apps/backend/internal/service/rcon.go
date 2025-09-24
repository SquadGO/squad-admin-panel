package service

import (
	"context"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-rcon-go/v2/rconTypes"
)

type RconService interface {
	UpdatePlayers(players rconTypes.Players)
}

type rconService struct {
	db            *db.Storage
	playerService PlayerService
}

func NewRconService(db *db.Storage, playerService PlayerService) RconService {
	return &rconService{
		db:            db,
		playerService: playerService,
	}
}

func (r *rconService) UpdatePlayers(players rconTypes.Players) {
	for _, player := range players {
		findPlayer, err := r.playerService.GetPlayerBySteamID(context.Background(), player.SteamID)
		if err != nil {
			r.playerService.CreatePlayer(context.Background(), models.CreatePlayer{
				Name:    player.PlayerName,
				EosID:   player.EosID,
				SteamID: player.SteamID,
			})
		} else if findPlayer.Name != player.PlayerName {
			r.playerService.UpdatePlayerName(context.Background(), player.PlayerName)
		}
	}
}
