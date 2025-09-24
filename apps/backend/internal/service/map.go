package service

import (
	"context"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type MapService interface {
	GetMaps(ctx context.Context) ([]gen.Map, error)
	CreateMap(ctx context.Context, gameMap models.Map) (int64, error)
	UpdateMap(ctx context.Context, gameMap models.Map) error
}

type mapService struct {
	db *db.Storage
}

func NewMapService(db *db.Storage) MapService {
	return &mapService{
		db: db,
	}
}

func (m *mapService) GetMaps(ctx context.Context) ([]gen.Map, error) {
	return m.db.Queries.GetMaps(ctx)
}

func (m *mapService) CreateMap(ctx context.Context, gameMap models.Map) (int64, error) {
	return m.db.Queries.InsertMap(ctx, gen.InsertMapParams{
		ServerID: gameMap.ServerID,
		MapName:  gameMap.MapName,
	})
}

func (m *mapService) UpdateMap(ctx context.Context, gameMap models.Map) error {
	return m.db.Queries.UpdateMap(ctx, gen.UpdateMapParams{
		MapID:         gameMap.MapID,
		WinnerName:    gameMap.WinnerName,
		WinnerTeamID:  gameMap.WinnerTeamID,
		WinnerTickets: gameMap.WinnerTickets,
		LoserName:     gameMap.LoserName,
		LoserTeamID:   gameMap.LoserTeamID,
		LoserTickets:  gameMap.LoserTickets})
}
