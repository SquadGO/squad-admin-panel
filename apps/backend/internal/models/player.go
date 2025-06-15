package models

type CreatePlayer struct {
	Name    string `json:"name" db:"name" binding:"required"`
	EosID   string `json:"eos_id" db:"eos_id" binding:"required,len=32"`
	SteamID string `json:"steam_id" db:"steam_id" binding:"required,len=17"`
}

type UpdatePlayer struct {
	Name string `json:"name" db:"name" binding:"required"`
}
