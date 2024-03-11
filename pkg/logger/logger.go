package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Logger struct {
	Sugar *zap.SugaredLogger
}

var _ LoggerApplication = (*Logger)(nil)

func NewLogger(viper *viper.Viper) *Logger {
	zap := zap.NewProductionConfig()

	if env := viper.GetString("app.env"); env != "" && env != "production" {
		zap.Development = true
	}

	build, _ := zap.Build()
	sugar := build.Sugar()

	sugar.Error()
	return &Logger{
		Sugar: sugar,
	}
}

type LoggerApplication interface {
	Close() error
	Debug(msg string, args any)
	Info(msg string, args any)
	Warn(msg string, args any)
	Error(msg string, args any)
}

func (l Logger) Close() error {
	return l.Sugar.Sync()
}

func (l Logger) Info(msg string, args any) {
	l.Sugar.Info(msg, args)
}

func (l Logger) Error(msg string, args any) {
	l.Sugar.Error(msg, args)
}

func (l Logger) Debug(msg string, args any) {
	l.Sugar.Debug(msg, args)
}

func (l Logger) Warn(msg string, args any) {
	l.Sugar.Warn(msg, args)
}
