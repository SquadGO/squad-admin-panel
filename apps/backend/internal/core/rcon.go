package core

import (
	"log/slog"
	"time"

	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	rcon "github.com/SquadGO/squad-rcon-go/v2"
	"github.com/SquadGO/squad-rcon-go/v2/rconEvents"
	"github.com/SquadGO/squad-rcon-go/v2/rconTypes"
)

func NewRcon(s *service.Service) {
	servers, err := s.ServerService.GetServers()
	if err != nil {
		slog.Error("Failed get servers rcon", slog.Any("err", err))
		return
	}

	if len(servers) == 0 {
		slog.Error("Servers not found")
		return
	}

	for _, server := range servers {
		r, err := rcon.NewRcon(rcon.RconConfig{Host: server.Host, Password: server.RconPassword, Port: server.RconPort, AutoReconnect: true, AutoReconnectDelay: 5})
		if err != nil {
			slog.Error("Failed rcon init", slog.Any("serverID", server.ServerID), slog.Any("err", err))
			return
		}

		slog.Info("[RCON] Connection successful")

		/* Listeners works after first initialization */

		r.Emitter.On(rconEvents.CONNECTED, func(_ interface{}) {
			slog.Info("[RCON] Connection successful", slog.Any("serverID", server.ServerID))
		})

		r.Emitter.On(rconEvents.RECONNECTING, func(_ interface{}) {
			slog.Info("[RCON] Reconnecting", slog.Any("serverID", server.ServerID))
		})

		r.Emitter.On(rconEvents.CLOSE, func(_ interface{}) {
			slog.Info("[RCON] Connection closed", slog.Any("serverID", server.ServerID))
		})

		r.Emitter.On(rconEvents.ERROR, func(err interface{}) {
			slog.Error("[RCON]", slog.Any("serverID", server.ServerID), slog.Any("err", err))
		})

		r.Emitter.On(rconEvents.CHAT_MESSAGE, func(data interface{}) {
			if v, ok := data.(rconTypes.Message); ok {
				data := models.ChatMessage{
					Message:  v.Message,
					SteamID:  v.SteamID,
					ChatType: gen.ChatType(v.ChatType),
				}

				err := s.LogsService.ChatMessage(&server.ServerID, data)
				if err != nil {
					slog.Error("Failed create message", slog.Any("err", err))
				}
			}
		})

		r.Emitter.On(rconEvents.SQUAD_CREATED, func(data interface{}) {
			if v, ok := data.(rconTypes.SquadCreated); ok {
				data := models.SquadCreated{
					SteamID:   v.SteamID,
					SquadID:   v.SquadID,
					SquadName: v.SquadName,
				}

				err := s.LogsService.SquadCreated(&server.ServerID, data)
				if err != nil {
					slog.Error("Failed squad created", slog.Any("err", err))
				}
			}
		})

		r.Emitter.On(rconEvents.POSSESSED_ADMIN_CAMERA, func(data interface{}) {
			if v, ok := data.(rconTypes.PosAdminCam); ok {
				data := models.AdminCamera{
					SteamID:   v.SteamID,
					AdminName: v.AdminName,
				}

				err := s.LogsService.AdminCam(&server.ServerID, rconEvents.POSSESSED_ADMIN_CAMERA, data)
				if err != nil {
					slog.Error("Failed pos admin created", slog.Any("err", err))
				}
			}
		})

		r.Emitter.On(rconEvents.UNPOSSESSED_ADMIN_CAMERA, func(data interface{}) {
			if v, ok := data.(rconTypes.UnposAdminCam); ok {
				data := models.AdminCamera{
					SteamID:   v.SteamID,
					AdminName: v.AdminName,
				}

				err := s.LogsService.AdminCam(&server.ServerID, rconEvents.UNPOSSESSED_ADMIN_CAMERA, data)
				if err != nil {
					slog.Error("Failed unpos admin created", slog.Any("err", err))
				}
			}
		})

		r.Emitter.On(rconEvents.LIST_PLAYERS, func(data interface{}) {
			if v, ok := data.(rconTypes.Players); ok {
				s.RconService.UpdatePlayers(v)
			}
		})

		ticker := time.NewTicker(5 * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					r.Execute(rconEvents.LIST_PLAYERS)
				}
			}
		}()
	}
}
