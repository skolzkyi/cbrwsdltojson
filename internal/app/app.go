package app

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	//customsoap "github.com/skolzkyi/cbrwsdltojson/internal/customsoap"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	//memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
)

const cbrNamespace = "http://web.cbr.ru/"

var (
	ErrAssertionAfterXMLDecoding  = errors.New("assertion error after XML decoding")
	ErrAssertionAfterGetCacheData = errors.New("assertion error after get cached data")
	ErrMethodProhibited           = errors.New("method prohibited")
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
	SoapCall(action string, payload interface{}) ([]byte, error)
}

type AppMemCache interface {
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
		fmt.Println("prsm.permittedRequests: ", prsm.permittedRequests)
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
	return !ok
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

func (a *App) RemoveDataInMemCacheBySOAPAction(SOAPAction string) {
	a.appmemcache.RemovePayloadInCache(SOAPAction)
}

func (a *App) GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (error, datastructures.GetCursOnDateXMLResult) {
	SOAPMethod := "GetCursOnDateXML"
	if a.permittedRequests.PermittedRequestMapLength() > 0 {
		fmt.Println("111")
		if !a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
			fmt.Println("222")
			return ErrMethodProhibited, datastructures.GetCursOnDateXMLResult{}
		}
	}
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
	fmt.Println("res: ", string(res))

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
	/*
		testGetCursOnDateXMLResult := datastructures.GetCursOnDateXMLResult{
			OnDate:           "20230622",
			ValuteCursOnDate: make([]datastructures.GetCursOnDateXMLResultElem, 2),
		}
		testGetCursOnDateXMLResultElem := datastructures.GetCursOnDateXMLResultElem{
			Vname:   "Австралийский доллар",
			Vnom:    1,
			Vcurs:   "57.1445",
			Vcode:   "36",
			VchCode: "AUD",
		}
		testGetCursOnDateXMLResult.ValuteCursOnDate[0] = testGetCursOnDateXMLResultElem
		testGetCursOnDateXMLResultElem = datastructures.GetCursOnDateXMLResultElem{
			Vname:   "Азербайджанский манат",
			Vnom:    1,
			Vcurs:   "49.5569",
			Vcode:   "944",
			VchCode: "AZN",
		}
		testGetCursOnDateXMLResult.ValuteCursOnDate[1] = testGetCursOnDateXMLResultElem
		testDataXMLMarsh, err := xml.Marshal(testGetCursOnDateXMLResult)
		if err != nil {
			a.logger.Error(err.Error())
			return err, response
		}
		fmt.Println("testDataXMLMarsh: ", string(testDataXMLMarsh))
	*/
	return nil, response

}
