package middleware

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	file, err := os.Create("logs.txt")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create log file")
	}
	log.Logger = log.Output(zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, file))
	log.Info().Msg("Logger created")
}
