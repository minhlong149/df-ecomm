package util

import "log"

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type LoggerImpl struct {
	logger *log.Logger
}

func SetupLogger() *LoggerImpl {
	return &LoggerImpl{logger: log.Default()}
}

func (l *LoggerImpl) Info(args ...interface{}) {
	l.logger.Printf("%v", args...)
}

func (l *LoggerImpl) Error(args ...interface{}) {
	l.logger.Fatalf("%v", args...)
}
