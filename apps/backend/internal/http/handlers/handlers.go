package handlers

import "github.com/SquadGO/squad-admin-panel/internal/service"

type handlers struct {
	AuthHandler AuthHandler
}

func NewHandlers(s *service.Service) handlers {
	return handlers{
		AuthHandler: NewAuthHandler(),
	}
}
