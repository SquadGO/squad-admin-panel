package service

import (
	"context"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
)

type ServerService interface {
	GetServers() ([]gen.Server, error)
}

type serverService struct {
	db *db.Storage
}

func NewServerService(db *db.Storage) ServerService {
	return &serverService{
		db: db,
	}
}

func (s *serverService) GetServers() ([]gen.Server, error) {
	return s.db.Queries.GetServers(context.Background())
}
