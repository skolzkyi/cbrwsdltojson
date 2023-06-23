package logger

import (
	"go.uber.org/zap"
)

type LogWrap struct {
	config zap.Config
	logger *zap.SugaredLogger
}

func New(level string) (*LogWrap, error) {
	logWrap := LogWrap{}
	zlevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	logWrap.config = zap.Config{
		Level:            zlevel,
		DisableCaller:    true,
		Development:      true,
		Encoding:         "console",
		OutputPaths:      []string{"stdout", "file_log.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}
	logWrap.logger = zap.Must(logWrap.config.Build()).Sugar()
	return &logWrap, nil
}

func (l LogWrap) GetZapLogger() *zap.SugaredLogger {
	return l.logger
}

func (l LogWrap) Info(msg string) {
	l.logger.Info(msg)
}

func (l LogWrap) Warning(msg string) {
	l.logger.Warn(msg)
}

func (l LogWrap) Error(msg string) {
	l.logger.Error(msg)
}

func (l LogWrap) Fatal(msg string) {
	l.logger.Fatal(msg)
}
