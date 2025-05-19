package service

import (
	"github.com/SquadGO/squad-admin-panel/internal/db"
)

type Service struct{}

func NewService(db *db.Storage) Service {
	return Service{}
}
