package db

import (
	"context"
	"os"

	"github.com/SquadGO/squad-admin-panel/internal/sqlGen"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn    *pgx.Conn
	Queries *sqlGen.Queries
}

func New(ctx context.Context) (*Storage, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("BACKEND_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	queries := sqlGen.New(conn)

	return &Storage{Queries: queries}, nil
}
