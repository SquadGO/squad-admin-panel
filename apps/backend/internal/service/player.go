package service

import (
	"context"
	"fmt"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type PlayerService interface {
	GetPlayerBySteamID(ctx context.Context, steamID string) (*gen.Player, error)
	GetPlayerByEosID(ctx context.Context, eosID string) (*gen.Player, error)
	CreatePlayer(ctx context.Context, player models.CreatePlayer) (int32, error)
	UpdatePlayerName(ctx context.Context, name string) error
}

type playerService struct {
	db *db.Storage
}

func NewPlayerService(db *db.Storage) PlayerService {
	return &playerService{
		db: db,
	}
}

func (p *playerService) CreatePlayer(ctx context.Context, player models.CreatePlayer) (int32, error) {
	id, err := p.db.Queries.InsertPlayer(ctx, gen.InsertPlayerParams{
		Name:    player.Name,
		EosID:   player.EosID,
		SteamID: player.SteamID,
	})

	if err != nil {
		return 0, fmt.Errorf("Failed create player: %w", err)
	}

	return id, nil
}

func (p *playerService) UpdatePlayerName(ctx context.Context, name string) error {
	err := p.db.Queries.UpdatePlayerName(ctx, name)

	if err != nil {
		return fmt.Errorf("Failed update player: %w", err)
	}

	return nil
}

func (p *playerService) GetPlayerBySteamID(ctx context.Context, steamID string) (*gen.Player, error) {
	player, err := p.db.Queries.GetPlayerBySteamID(ctx, steamID)

	if err != nil {
		return nil, fmt.Errorf("User not found: %w", err)
	}

	return &player, nil
}

func (p *playerService) GetPlayerByEosID(ctx context.Context, eosID string) (*gen.Player, error) {
	player, err := p.db.Queries.GetPlayerByEosID(ctx, eosID)

	if err != nil {
		return nil, fmt.Errorf("User not found: %w", err)
	}

	return &player, nil
}
