package app

import (
	"context"
	"time"

	"go.uber.org/zap"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

type App struct {
	logger        Logger
	config        Config
	soapReqSender SoapRequestSender
}

type Logger interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
	GetZapLogger() *zap.SugaredLogger
}

type Config interface {
	Init(path string) error
	GetServerURL() string
	GetAddress() string
	GetPort() string
	GetServerShutdownTimeout() time.Duration
	GetCBRWSDLTimeout() time.Duration
	GetCBRWSDLAddress() string
	GetLoggingOn() bool
	GetDateTimeResponseLayout() string
	GetDateTimeRequestLayout() string
	GetPermittedRequests() map[struct{}]string
}

type SoapRequestSender interface {
	GetCursOnDate(ctx context.Context, input datastructures.RequestOnDate) (error, datastructures.ResponseValuteCursDynamic)
}

func New(logger Logger, config Config, soapReqSender SoapRequestSender) *App {
	app := App{
		logger:        logger,
		config:        config,
		soapReqSender: soapReqSender,
	}
	return &app
}

func (a *App) GetCursOnDate(ctx context.Context, input datastructures.RequestOnDate) (error, datastructures.ResponseValuteCursDynamic) {
	var response datastructures.ResponseValuteCursDynamic
	err, response := a.soapReqSender.GetCursOnDate(ctx, input)
	if err != nil {
		return err, response
	}
	return nil, response
}
