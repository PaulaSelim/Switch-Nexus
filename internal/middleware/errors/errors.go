package errors

type AppError struct {
	Code    int
	Message string
}

const (
	ErrCodeConnectionFailed = iota + 1
	ErrCodeWailsInitFailed  = iota + 2
)
