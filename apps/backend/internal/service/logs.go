package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-admin-panel/internal/state"
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-rcon-go/v2/rconEvents"
)

type LogsService interface {
	AdminCam(serverID int32, logType string, data models.AdminCamera) error
	ChatMessage(serverID int32, data models.ChatMessage) error
	SquadCreated(serverID int32, data models.SquadCreated) error
	PlayerConnected(serverID int32, data models.PlayerConnected) error
	PlayerDisconnected(serverID int32, data models.PlayerDisconnected) error
	NewGame(serverID int32, data models.NewGame) error
	RoundTickets(serverID int32, data models.RoundTickets) error
	RoundEnded(serverID int32) error
}

type logsService struct {
	db            *db.Storage
	appState      *state.AppState
	playerService PlayerService
	userService   UserService
	mapService    MapService
}

func NewLogsService(db *db.Storage, appState *state.AppState, playerService PlayerService, userService UserService, mapService MapService) LogsService {
	return &logsService{
		db:            db,
		appState:      appState,
		playerService: playerService,
		userService:   userService,
		mapService:    mapService,
	}
}

func (l *logsService) AdminCam(serverID int32, logType string, data models.AdminCamera) error {
	op := "service.logs.createPosAdminCam"

	player, err := l.playerService.GetPlayerBySteamID(context.Background(), data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	user, err := l.userService.GetUserBySteamID(context.Background(), data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: &serverID,
		LogType:  logType,
		PlayerID: &player.PlayerID,
		UserID:   &user.UserID,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

// TODO TEAM ID
func (l *logsService) SquadCreated(serverID int32, data models.SquadCreated) error {
	op := "service.logs.createSquadCreated"

	player, err := l.playerService.GetPlayerBySteamID(context.Background(), data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID:  &serverID,
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

func (l *logsService) ChatMessage(serverID int32, data models.ChatMessage) error {
	op := "service.logs.createChatMessage"

	player, err := l.playerService.GetPlayerBySteamID(context.Background(), data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: &serverID,
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

func (l *logsService) PlayerConnected(serverID int32, data models.PlayerConnected) error {
	op := "service.logs.playerConnected"

	player, err := l.playerService.GetPlayerBySteamID(context.Background(), data.SteamID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: &serverID,
		LogType:  logsEvents.PLAYER_CONNECTED,
		PlayerID: &player.PlayerID,
		PlayerIp: &data.Ip,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (l *logsService) PlayerDisconnected(serverID int32, data models.PlayerDisconnected) error {
	op := "service.logs.playerDisconnected"

	player, err := l.playerService.GetPlayerByEosID(context.Background(), data.EosID)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	_, err = l.db.Queries.InsertLog(context.Background(), gen.InsertLogParams{
		ServerID: &serverID,
		LogType:  logsEvents.PLAYER_DISCONNECTED,
		PlayerID: &player.PlayerID,
		PlayerIp: &data.Ip,
	})

	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (l *logsService) NewGame(serverID int32, data models.NewGame) error {
	op := "service.logs.newGame"

	l.appState.Servers[serverID] = &state.ServerState{
		CurrentMap: models.Map{
			ServerID: &serverID,
			MapName:  data.MapClassname,
		},
	}

	if v, ok := l.appState.Servers[serverID]; ok {
		_, err := l.mapService.CreateMap(context.Background(), v.CurrentMap)
		if err != nil {
			return fmt.Errorf("%s %w", op, err)
		}
	}

	return nil
}

func (l *logsService) RoundTickets(serverID int32, data models.RoundTickets) error {
	op := "service.logs.roundTickets"

	server, ok := l.appState.Servers[serverID]
	if !ok {
		return fmt.Errorf("%s: server %d not found", op, serverID)
	}

	teamID, err := strconv.Atoi(data.Team)
	if err != nil {
		teamID = 1
	}

	convertedTeamID := int32(teamID)
	convertedTickets := int32(data.Tickets)

	switch data.Action {
	case "won":
		server.CurrentMap.WinnerName = &data.Faction
		server.CurrentMap.WinnerTeamID = &convertedTeamID
		server.CurrentMap.WinnerTickets = &convertedTickets

	case "lost":
		server.CurrentMap.LoserName = &data.Faction
		server.CurrentMap.LoserTeamID = &convertedTeamID
		server.CurrentMap.LoserTickets = &convertedTickets
	default:
		return fmt.Errorf("%s: unknown action %q", op, data.Action)
	}

	return nil
}

func (l *logsService) RoundEnded(serverID int32) error {
	op := "service.logs.roundEnded"

	server, ok := l.appState.Servers[serverID]
	if !ok {
		return fmt.Errorf("%s: server %d not found", op, serverID)
	}

	err := l.mapService.UpdateMap(context.Background(), server.CurrentMap)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
