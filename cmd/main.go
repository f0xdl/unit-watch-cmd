package main

import (
	"github.com/f0xdl/unit-watch-cmd/internal/uwcli"
	"github.com/f0xdl/unit-watch-lib/storage"
	"github.com/f0xdl/unit-watch-lib/utils"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

type Config struct {
	BotDb string `env:"BOT_DB"`
}

func main() {
	cfg := &Config{}
	err := utils.LoadConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config file")
	}
	store, err := storage.NewGormStorage(cfg.BotDb, nil, false)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	c := uwcli.NewUWCli(store)
	app := &cli.App{
		Name:     "unit-watch-cmd",
		Usage:    "Утилита для управления устройствами",
		Commands: c.BuildCommands(),
	}

	if err = app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("run error")
	}
}
