package models

import (
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/jackc/pgx/pgtype"
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

type Map struct {
	MapID         int64              `json:"map_id"`
	ServerID      *int32             `json:"server_id"`
	EndAt         pgtype.Timestamptz `json:"end_at"`
	MapName       string             `json:"map_name"`
	WinnerName    *string            `json:"winner_name"`
	WinnerTeamID  *int32             `json:"winner_team_id"`
	WinnerTickets *int32             `json:"winner_tickets"`
	LoserName     *string            `json:"loser_name"`
	LoserTeamID   *int32             `json:"loser_team_id"`
	LoserTickets  *int32             `json:"loser_tickets"`
}

type NewGame struct {
	MapClassname   string
	LayerClassname string
}

type RoundTickets struct {
	Team       string
	Subfaction string
	Faction    string
	Action     string
	Tickets    int
	Layer      string
	Level      string
}

type RoundEnded struct {
	VictimName               string
	Damage                   float64
	AttackerPlayerController string
	AttackerEOSID            string
	AttackerSteamID          string
	Weapon                   string
}
