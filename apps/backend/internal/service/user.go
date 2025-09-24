package service

import (
	"context"

	"github.com/SquadGO/squad-admin-panel/internal/db"
	"github.com/SquadGO/squad-admin-panel/internal/gen"
	"github.com/SquadGO/squad-admin-panel/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.CreateUser) (int32, error)
	GetUserBySteamID(ctx context.Context, steamID string) (gen.User, error)
}

type userService struct {
	db *db.Storage
}

func NewUserService(db *db.Storage) UserService {
	return &userService{
		db: db,
	}
}

func (us *userService) CreateUser(ctx context.Context, user models.CreateUser) (int32, error) {
	return us.db.Queries.InsertUser(ctx, gen.InsertUserParams{
		SteamID: user.SteamID,
		Name:    user.Name,
		Avatar:  user.Avatar,
	})
}

func (us *userService) GetUserBySteamID(ctx context.Context, steamID string) (gen.User, error) {
	return us.db.Queries.GetUserBySteamID(ctx, steamID)
}
