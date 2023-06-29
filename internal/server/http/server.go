package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"time"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	"go.uber.org/zap"
)

type Server struct {
	serv   *http.Server
	logg   Logger
	app    Application
	Config Config
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

type Logger interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
	GetZapLogger() *zap.SugaredLogger
}

type Application interface {
	RemoveDataInMemCacheBySOAPAction(SOAPAction string)
	GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (error, datastructures.GetCursOnDateXMLResult)
}

func NewServer(logger Logger, app Application, config Config) *Server {
	server := Server{}
	server.logg = logger
	server.app = app
	server.Config = config
	server.serv = &http.Server{
		Addr:              config.GetServerURL(),
		Handler:           server.routes(),
		ReadHeaderTimeout: 2 * time.Second,
	}

	return &server
}

func (s *Server) Start(ctx context.Context) error {
	s.logg.Info("cbrwsdltojson is running...")
	err := s.serv.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logg.Error("server start error: " + err.Error())
			return err
		}
	}
	<-ctx.Done()
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.serv.Shutdown(ctx)
	if err != nil {
		s.logg.Error("server shutdown error: " + err.Error())
		return err
	}
	s.logg.Info("server graceful shutdown")
	return err
}
