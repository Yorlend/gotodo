package data

import (
	"github.com/go-pg/pg"
	"github.com/yorlend/gotodo/internal/modules/config"
)

type Client struct {
	*pg.DB
}

func NewDBClient(
	cfg *config.DBConfig,
) *Client {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Addr,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	return &Client{DB: db}
}
