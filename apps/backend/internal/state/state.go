package state

import (
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type ServerState struct {
	ID         int
	CurrentMap models.Map
}

type AppState struct {
	Servers map[int32]*ServerState
}
