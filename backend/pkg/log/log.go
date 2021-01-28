package log

import (
	"go.uber.org/zap"
)

var (
	global *Logger
)

type Logger struct {
	*zap.SugaredLogger
}

func New() *Logger {
	return &Logger{
		SugaredLogger: zap.NewNop().Sugar(),
	}
}

func Global() *Logger {
	if global == nil {
		global = New()
	}

	return global
}
