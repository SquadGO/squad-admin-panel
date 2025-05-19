package service

import "github.com/SquadGO/squad-admin-panel/internal/db"

type CoreService interface {
}

type coreService struct {
	db *db.Storage
}

func NewCoreService(db *db.Storage) CoreService {
	return coreService{
		db: db,
	}
}
