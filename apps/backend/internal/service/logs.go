package service

import (
	"context"
	"fmt"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-rcon-go/v2/rconEvents"
)

type LogsService interface {
	AdminCam(serverID *int32, logType string, data models.AdminCamera) error
	ChatMessage(serverID *int32, data models.ChatMessage) error
	SquadCreated(serverID *int32, data models.SquadCreated) error
	PlayerConnected(serverID *int32, data models.PlayerConnected) error
	PlayerDisconnected(serverID *int32, data models.PlayerDisconnected) error
}

type logsService struct {
	db            *db.Storage
	playerService PlayerService
}

func NewLogsService(db *db.Storage, playerService PlayerService) LogsService {
	return &logsService{
		db:            db,
		playerService: playerService,
	}
}

// TODO ADD USER
func (l *logsService) AdminCam(serverID *int32, logType string, data models.AdminCamera) error {
	op := "service.logs.createPosAdminCam"

	player, err := l.playerService.GetPlayerBySteamID(data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: serverID,
		LogType:  logType,
		PlayerID: &player.PlayerID,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

// TODO TEAM ID
func (l *logsService) SquadCreated(serverID *int32, data models.SquadCreated) error {
	op := "service.logs.createSquadCreated"

	player, err := l.playerService.GetPlayerBySteamID(data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID:  serverID,
		LogType:   rconEvents.SQUAD_CREATED,
		PlayerID:  &player.PlayerID,
		SquadID:   &data.SquadID,
		SquadName: &data.SquadName,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (l *logsService) ChatMessage(serverID *int32, data models.ChatMessage) error {
	op := "service.logs.createChatMessage"

	player, err := l.playerService.GetPlayerBySteamID(data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: serverID,
		LogType:  rconEvents.CHAT_MESSAGE,
		PlayerID: &player.PlayerID,
		Message:  &data.Message,
		ChatType: gen.NullChatType{
			ChatType: data.ChatType,
			Valid:    true,
		},
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (l *logsService) PlayerConnected(serverID *int32, data models.PlayerConnected) error {
	op := "service.logs.playerConnected"

	player, err := l.playerService.GetPlayerBySteamID(data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: serverID,
		LogType:  logsEvents.PLAYER_CONNECTED,
		PlayerID: &player.PlayerID,
		PlayerIp: &data.Ip,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (l *logsService) PlayerDisconnected(serverID *int32, data models.PlayerDisconnected) error {
	op := "service.logs.playerDisconnected"

	player, err := l.playerService.GetPlayerByEosID(data.EosID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: serverID,
		LogType:  logsEvents.PLAYER_DISCONNECTED,
		PlayerID: &player.PlayerID,
		PlayerIp: &data.Ip,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
