package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

var ErrAssertionGetFullRequestTimeout = errors.New("error of data assertion on get full request timeout")

type Server struct {
	serv              *http.Server
	logg              Logger
	app               Application
	Config            Config
	fullRequestTimeot atomic.Value
}

type Config interface {
	Init(path string) error
	GetServerURL() string
	GetAddress() string
	GetPort() string
	GetServerShutdownTimeout() time.Duration
	GetCBRWSDLTimeout() time.Duration
	GetInfoExpirTime() time.Duration
	GetCBRWSDLAddress() string
	GetLoggingOn() bool
	GetPermittedRequests() map[string]struct{}
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
	GetCursOnDateXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	BiCurBaseXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	BliquidityXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	DepoDynamicXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	DragMetDynamicXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	DVXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	EnumReutersValutesXML(ctx context.Context) (interface{}, error)
	EnumValutesXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	KeyRateXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
	MainInfoXML(ctx context.Context) (interface{}, error)
	Mrrf7DXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error)
}

func NewServer(logger Logger, app Application, config Config) *Server {
	server := Server{}
	server.logg = logger
	server.app = app
	server.Config = config
	server.fullRequestTimeot.Store(config.GetCBRWSDLTimeout())
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

func (s *Server) GetFullRequestTimeout() (time.Duration, error) {
	var timeout time.Duration
	timeoutAny := s.fullRequestTimeot.Load()
	timeout, ok := timeoutAny.(time.Duration)
	if !ok {
		return timeout, ErrAssertionGetFullRequestTimeout
	}
	return timeout, nil
}

func (s *Server) ReadDataFromInputJSON(pointerOnStruct interface{}, r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.logg.Error("server ReadDataFromInputJSON error: " + err.Error())
		return "", err
	}

	err = json.Unmarshal(body, pointerOnStruct)
	if err != nil {
		s.logg.Error("server ReadDataFromInputJSON error: " + err.Error())
		return "", err
	}

	return string(body), nil
}

func (s *Server) WriteDataToOutputJSON(marshallingObject interface{}, w http.ResponseWriter) error {
	jsonstring, err := json.Marshal(marshallingObject)
	if err != nil {
		s.logg.Error("server WriteDataToOutputJSON error: " + err.Error())
		return err
	}
	_, err = w.Write(jsonstring)
	if err != nil {
		s.logg.Error("server WriteDataToOutputJSON error: " + err.Error())
		return err
	}
	return nil
}
