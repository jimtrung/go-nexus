package zap

import "go.uber.org/zap"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(message string, args ...any) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	if len(args) > 0 {
		sugar.Info(message, args)
	} else {
		sugar.Info(message)
	}
}

func (lo Logger) Error(message string, args ...any) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	if len(args) > 0 {
		sugar.Error(message, args)
	} else {
		sugar.Error(message)
	}
}

