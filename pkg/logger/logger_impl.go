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
	logger := zap.NewProductionConfig()

	if env := viper.GetString("app.env"); env != "" && env != "production" {
		logger.Development = true
	}

	build, _ := logger.Build()
	sugar := build.Sugar()

	return &Logger{
		Sugar: sugar,
	}
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
