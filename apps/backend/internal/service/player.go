package service

import (
	"context"
	"fmt"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type PlayerService interface {
	GetPlayerBySteamID(steamID string) (*gen.Player, error)
	GetPlayerByEosID(eosID string) (*gen.Player, error)
	CreatePlayer(player models.CreatePlayer) (int32, error)
	UpdatePlayerName(name string) error
}

type playerService struct {
	db *db.Storage
}

func NewPlayerService(db *db.Storage) PlayerService {
	return &playerService{
		db: db,
	}
}

func (p *playerService) CreatePlayer(player models.CreatePlayer) (int32, error) {
	id, err := p.db.Queries.InsertPlayer(context.Background(), gen.InsertPlayerParams{
		Name:    player.Name,
		EosID:   player.EosID,
		SteamID: player.SteamID,
	})

	if err != nil {
		return 0, fmt.Errorf("Failed create player: %w", err)
	}

	return id, nil
}

func (p *playerService) UpdatePlayerName(name string) error {
	err := p.db.Queries.UpdatePlayerName(context.Background(), name)

	if err != nil {
		return fmt.Errorf("Failed update player: %w", err)
	}

	return nil
}

func (p *playerService) GetPlayerBySteamID(steamID string) (*gen.Player, error) {
	player, err := p.db.Queries.GetPlayerBySteamID(context.Background(), steamID)

	if err != nil {
		return nil, fmt.Errorf("User not found: %w", err)
	}

	return &player, nil
}

func (p *playerService) GetPlayerByEosID(eosID string) (*gen.Player, error) {
	player, err := p.db.Queries.GetPlayerByEosID(context.Background(), eosID)

	if err != nil {
		return nil, fmt.Errorf("User not found: %w", err)
	}

	return &player, nil
}
