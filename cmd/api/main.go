package main

import (
	"github.com/billyboar/bytely/api"
	"github.com/billyboar/bytely/api/config"
	"github.com/billyboar/bytely/internal/storage/pg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	appName = "bytely"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
	}

	setGlobalLevel(cfg.LogLevel)

	logger := log.With().Str("service_name", appName).Logger()

	logger.Debug().Msg("Starting service...")

	pgClient, err := pg.NewClient(cfg.DatabaseConnectionStr)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot connect to the database")
	}
	if err := pgClient.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("cannot ping the database")
	}

	logger.Debug().Msg("Initialized database connection")

	srv := api.NewServer(pgClient, cfg)
	if err := srv.Start(); err != nil {
		logger.Fatal().Err(err).Msg("cannot start the service")
	}

	defer srv.Shutdown()
}

func setGlobalLevel(logLevel string) {
	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
