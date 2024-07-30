package middleware

import "github.com/rs/zerolog/log"

type AppError struct {
	Code    int
	Message string
}

const (
	ErrCodeConnectionFailed = iota + 1
)

func (e *AppError) Error() string {
	return e.Message
}

func LogError(err error) {
	log.Error().Err(err).Msg("An error occurred")
}
