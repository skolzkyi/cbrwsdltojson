package app

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	helpers "github.com/skolzkyi/cbrwsdltojson/helpers"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
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
	Appmemcache       AppMemCache
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
	GetInfoExpirTime() time.Duration
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
	GetCacheDataInCache(tag string) (memcache.CacheInfo, bool)
	PrintAllCacheKeys()
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
		Appmemcache:       memcache,
		permittedRequests: NewPermittedReqSyncMap(),
	}
	app.permittedRequests.Init(permReqMap)
	return &app
}

func (a *App) ProcessRequest(ctx context.Context, SOAPMethod string, startNodeName string, inputData interface{}, responseData interface{}, pointerToResponseData interface{}) error {
	res, err := a.soapSender.SoapCall(ctx, SOAPMethod, inputData)
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}
	err = a.XMLToStructDecoder(res, startNodeName, pointerToResponseData)
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}

	err = a.AddOrUpdateDataInCache(SOAPMethod, inputData, responseData)
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}
	return nil
}

func (a *App) XMLToStructDecoder(data []byte, startNodeName string, pointerToStruct interface{}) error {
	xmlData := bytes.NewBuffer(data)

	d := xml.NewDecoder(xmlData)

	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) { //nolint: gocritic
		case xml.StartElement:
			if se.Name.Local == startNodeName {
				err := d.DecodeElement(pointerToStruct, &se)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *App) GetDataInCacheIfExisting(SOAPMethod string, rawBodyIn string) (interface{}, bool) { //nolint: gocritic
	rawBody := helpers.ClearStringByWhitespaceAndLinebreak(rawBodyIn)
	cachedData, ok := a.Appmemcache.GetCacheDataInCache(SOAPMethod + rawBody)
	if ok {
		if cachedData.InfoDTStamp.Add(a.config.GetInfoExpirTime()).After(time.Now()) {
			return cachedData.Payload, true
		}
	}
	return nil, false
}

func (a *App) AddOrUpdateDataInCache(SOAPMethod string, request interface{}, response interface{}) error { //nolint: gocritic
	jsonstring, err := json.Marshal(request)
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}

	rawBody := helpers.ClearStringByWhitespaceAndLinebreak(string(jsonstring))
	a.Appmemcache.AddOrUpdatePayloadInCache(SOAPMethod+rawBody, response)
	return nil
}

func (a *App) RemoveDataInMemCacheBySOAPAction(tag string) {
	a.Appmemcache.RemovePayloadInCache(tag)
}

func (a *App) GetCursOnDateXML(ctx context.Context, input datastructures.GetCursOnDateXML, rawBody string) (datastructures.GetCursOnDateXMLResult, error) {
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

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
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

		err = a.XMLToStructDecoder(res, startNodeName, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		for i := range response.ValuteCursOnDate {
			response.ValuteCursOnDate[i].Vname = strings.TrimSpace(response.ValuteCursOnDate[i].Vname)
			response.ValuteCursOnDate[i].Vname = strings.Trim(response.ValuteCursOnDate[i].Vname, "\r\n")
		}
		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, err
}

func (a *App) BiCurBaseXML(ctx context.Context, input datastructures.BiCurBaseXML, rawBody string) (datastructures.BiCurBaseXMLResult, error) {
	var err error
	var response datastructures.BiCurBaseXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
	default:
		SOAPMethod := "BiCurBaseXML"
		startNodeName := "BiCurBase"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.BiCurBaseXMLResult{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
		if ok {
			response, ok = cachedData.(datastructures.BiCurBaseXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		input.XMLNs = cbrNamespace
		/*
			err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, input, response, &response)
			if err != nil {
				a.logger.Error(err.Error())
				return response, err
			}*/
		res, err := a.soapSender.SoapCall(ctx, SOAPMethod, input)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.XMLToStructDecoder(res, startNodeName, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, err
}
