package models

import "time"

type CreateUser struct {
	SteamID   string    `json:"steam_id"`
	Name      string    `json:"name"`
	Avatar    *string   `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
