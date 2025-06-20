package core

import (
	"fmt"
	"log/slog"

	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	logs "github.com/SquadGO/squad-logs-go"
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func NewLocalLogs(s *service.Service) {}

func NewFTPLogs(s *service.Service) {
	servers, err := s.ServerService.GetServers()
	if err != nil {
		slog.Error("Failed get servers logs", slog.Any("err", err))
		return
	}

	for _, server := range servers {
		config := logs.FTPReaderConfig{
			Host:               fmt.Sprintf("%s:22", server.Host),
			Username:           server.User,
			Password:           server.Password,
			LogPath:            server.GameLogFilePath,
			AdminsPath:         server.AdminLogFilePath,
			LogEnabled:         true,
			AutoReconnect:      true,
			AutoReconnectDelay: 5,
		}

		fr, err := logs.NewFTPReader(config)
		if err != nil {
			slog.Error("Failed logs init", slog.Any("err", err))
			return
		}

		defer fr.Close()

		slog.Info("[LOGS] Connection successful")

		/* Listeners works after first initialization */

		fr.Emitter.On(logsEvents.CONNECTED, func(_ interface{}) {
			slog.Info("[LOGS] Connection successful")
		})

		fr.Emitter.On(logsEvents.RECONNECTING, func(i interface{}) {
			slog.Info("[LOGS] Reconnecting")
		})

		fr.Emitter.On(logsEvents.CLOSE, func(_ interface{}) {
			slog.Info("[LOGS] Connection closed")
		})

		fr.Emitter.On(logsEvents.ERROR, func(err interface{}) {
			slog.Error("[LOGS]", slog.Any("err", err))
		})

		fr.Emitter.On(logsEvents.PLAYER_CONNECTED, func(i interface{}) {
			if v, ok := i.(logsTypes.PlayerConnected); ok {
				data := models.PlayerConnected{
					SteamID: v.SteamID,
					Ip:      v.Ip,
				}

				err := s.LogsService.PlayerConnected(&server.ServerID, data)
				if err != nil {
					slog.Error("Failed player connected", slog.Any("err", err))
				}
			}
		})

		fr.Emitter.On(logsEvents.PLAYER_DISCONNECTED, func(i interface{}) {
			if v, ok := i.(logsTypes.PlayerDisconnected); ok {
				data := models.PlayerDisconnected{
					EosID: v.EosID,
					Ip:    v.Ip,
				}

				err := s.LogsService.PlayerDisconnected(&server.ServerID, data)
				if err != nil {
					slog.Error("Failed player disconnected", slog.Any("err", err))
				}
			}
		})

		fr.Emitter.On(logsEvents.PLAYER_DAMAGED, func(i interface{}) {
			if data, ok := i.(logsTypes.PlayerDamaged); ok {
				slog.Info("PLAYER_DAMAGED", slog.Any("data", data))
			}
		})

		fr.Emitter.On(logsEvents.SERVER_TICKRATE, func(i interface{}) {
			if data, ok := i.(logsTypes.ServerTickrate); ok {
				slog.Info("TICKRATE", slog.Any("data", data))
			}
		})
	}
}
