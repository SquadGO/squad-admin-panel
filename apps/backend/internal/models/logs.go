package models

import (
	"github.com/SquadGO/squad-admin-panel/internal/gen"
)

type ChatMessage struct {
	Message  string
	SteamID  string
	ChatType gen.ChatType
}

type SquadCreated struct {
	SteamID   string
	SquadID   string
	SquadName string
}

type AdminCamera struct {
	SteamID   string
	AdminName string
}

type PlayerConnected struct {
	SteamID string
	Ip      string
}

type PlayerDisconnected struct {
	EosID string
	Ip    string
}
