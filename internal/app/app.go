package app

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

const cbrNamespace = "http://web.cbr.ru/"

var (
	ErrAssertionAfterXMLDecoding  = errors.New("assertion error after XML decoding")
	ErrAssertionAfterGetCacheData = errors.New("assertion error after get cached data")
	ErrMethodProhibited           = errors.New("method prohibited")
	ErrContextWSReqExpired        = errors.New("context of request to CBR WS expired")
)

type App struct {
	logger            Logger
	config            Config
	soapSender        SoapRequestSender
	appmemcache       AppMemCache
	permittedRequests PermittedReqSyncMap
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
	GetPermittedRequests() map[string]struct{}
}

type SoapRequestSender interface {
	SoapCall(ctx context.Context, action string, payload interface{}) ([]byte, error)
}

type AppMemCache interface { //nolint: revive
	Init()
	AddOrUpdatePayloadInCache(tag string, payload interface{}) bool
	RemovePayloadInCache(tag string)
	GetPayloadInCache(tag string) (interface{}, bool)
}

type PermittedReqSyncMap struct {
	mu                sync.RWMutex
	permittedRequests map[string]struct{}
}

func NewPermittedReqSyncMap() PermittedReqSyncMap {
	return PermittedReqSyncMap{}
}

func (prsm *PermittedReqSyncMap) Init(initMap map[string]struct{}) {
	prsm.mu.Lock()
	defer prsm.mu.Unlock()
	prsm.permittedRequests = make(map[string]struct{})
	if len(initMap) > 0 {
		prsm.permittedRequests = initMap
	}
}

func (prsm *PermittedReqSyncMap) AddPermittedRequest(permittedRequest string) {
	prsm.mu.Lock()
	defer prsm.mu.Unlock()
	prsm.permittedRequests[permittedRequest] = struct{}{}
}

func (prsm *PermittedReqSyncMap) IsPermittedRequestInMap(permittedRequest string) bool {
	prsm.mu.RLock()
	defer prsm.mu.RUnlock()
	_, ok := prsm.permittedRequests[permittedRequest]
	return ok
}

func (prsm *PermittedReqSyncMap) PermittedRequestMapLength() int {
	prsm.mu.RLock()
	defer prsm.mu.RUnlock()
	length := len(prsm.permittedRequests)
	return length
}

func New(logger Logger, config Config, sender SoapRequestSender, memcache AppMemCache, permReqMap map[string]struct{}) *App {
	app := App{
		logger:            logger,
		config:            config,
		soapSender:        sender,
		appmemcache:       memcache,
		permittedRequests: NewPermittedReqSyncMap(),
	}
	app.permittedRequests.Init(permReqMap)
	return &app
}

func (a *App) RemoveDataInMemCacheBySOAPAction(SOAPAction string) { //nolint: gocritic
	a.appmemcache.RemovePayloadInCache(SOAPAction)
}

func (a *App) GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (datastructures.GetCursOnDateXMLResult, error) {
	var err error
	var response datastructures.GetCursOnDateXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
	default:
		SOAPMethod := "GetCursOnDateXML"
		startNodeName := "ValuteData"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.GetCursOnDateXMLResult{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.appmemcache.GetPayloadInCache(SOAPMethod)
		if ok {
			response, ok = cachedData.(datastructures.GetCursOnDateXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		input.XMLNs = cbrNamespace

		res, err := a.soapSender.SoapCall(ctx, SOAPMethod, input)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		xmlData := bytes.NewBuffer(res)

		d := xml.NewDecoder(xmlData)

		for t, _ := d.Token(); t != nil; t, _ = d.Token() {
			switch se := t.(type) { //nolint: gocritic
			case xml.StartElement:
				if se.Name.Local == startNodeName {
					err = d.DecodeElement(&response, &se)
					if err != nil {
						return response, err
					}
				}
			}
		}

		for i := range response.ValuteCursOnDate {
			response.ValuteCursOnDate[i].Vname = strings.TrimSpace(response.ValuteCursOnDate[i].Vname)
			response.ValuteCursOnDate[i].Vname = strings.Trim(response.ValuteCursOnDate[i].Vname, "\r\n")
		}
		a.appmemcache.AddOrUpdatePayloadInCache(SOAPMethod, response)
	}
	return response, err
}
