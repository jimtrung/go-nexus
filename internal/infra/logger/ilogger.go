package logger

type ILogger interface {
    Info(message string, args ...any)
    Error(message string, args ...any)
}

