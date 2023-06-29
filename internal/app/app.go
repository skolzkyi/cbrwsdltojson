package app

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"

	//customsoap "github.com/skolzkyi/cbrwsdltojson/internal/customsoap"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	//memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
)

const cbrNamespace = "http://web.cbr.ru/"

var ErrAssertionAfterXMLDecoding = errors.New("assertion error after XML decoding")
var ErrAssertionAfterGetCacheData = errors.New("assertion error after get cached data")

type App struct {
	logger      Logger
	config      Config
	soapSender  SoapRequestSender
	appmemcache AppMemCache
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
	SoapCall(action string, payload interface{}) ([]byte, error)
}

type AppMemCache interface {
	Init()
	AddOrUpdatePayloadInCache(tag string, payload interface{}) bool
	RemovePayloadInCache(tag string)
	GetPayloadInCache(tag string) (interface{}, bool)
}

func New(logger Logger, config Config, sender SoapRequestSender, memcache AppMemCache) *App {
	app := App{
		logger:      logger,
		config:      config,
		soapSender:  sender,
		appmemcache: memcache,
	}
	return &app
}

func (a *App) RemoveDataInMemCacheBySOAPAction(SOAPAction string) {
	a.appmemcache.RemovePayloadInCache(SOAPAction)
}

func (a *App) GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (error, datastructures.GetCursOnDateXMLResult) {
	SOAPMethod := "GetCursOnDateXML"
	startNodeName := "ValuteData"
	var response datastructures.GetCursOnDateXMLResult
	var err error

	cachedData, ok := a.appmemcache.GetPayloadInCache(SOAPMethod)
	fmt.Println("cachedData: ", ok)
	if ok {
		response, ok = cachedData.(datastructures.GetCursOnDateXMLResult)
		if !ok {
			err = ErrAssertionAfterGetCacheData
			a.logger.Error(err.Error())
		} else {
			return err, response
		}
	}

	input.XMLNs = cbrNamespace

	res, err := a.soapSender.SoapCall(SOAPMethod, input)

	if err != nil {
		a.logger.Error(err.Error())
		return err, response
	}
	//fmt.Println("res: ", string(res))

	xmlData := bytes.NewBuffer(res)

	d := xml.NewDecoder(xmlData)

	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			fmt.Println("curElement: ", se.Name.Local)
			if se.Name.Local == startNodeName {
				err = d.DecodeElement(&response, &se)
				if err != nil {
					fmt.Println("d.DecodeElement err: ", err.Error())
					return err, response
				}
			}
		}
	}

	for i, _ := range response.ValuteCursOnDate {
		response.ValuteCursOnDate[i].Vname = strings.TrimSpace(response.ValuteCursOnDate[i].Vname)
		response.ValuteCursOnDate[i].Vname = strings.Trim(response.ValuteCursOnDate[i].Vname, "\r\n")
	}

	a.appmemcache.AddOrUpdatePayloadInCache(SOAPMethod, response)

	return nil, response

}
