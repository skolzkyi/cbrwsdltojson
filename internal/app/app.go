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

	customsoap "github.com/skolzkyi/cbrwsdltojson/internal/customsoap"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

const cbrNamespace = "http://web.cbr.ru/"

var ErrAssertionAfterXMLDecoding = errors.New("assertion error after XML decoding")

type App struct {
	logger     Logger
	config     Config
	soapSender customsoap.CBRSOAPSender
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
	GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (error, datastructures.GetCursOnDateXMLResult)
}

func New(logger Logger, config Config, sender customsoap.CBRSOAPSender) *App {
	app := App{
		logger:     logger,
		config:     config,
		soapSender: sender,
	}
	return &app
}

func (a *App) GetCursOnDate(ctx context.Context, input datastructures.GetCursOnDateXML) (error, datastructures.GetCursOnDateXMLResult) {
	SOAPMethod := "GetCursOnDateXML"
	startNodeName := "ValuteData"
	var response datastructures.GetCursOnDateXMLResult

	input.XMLNs = cbrNamespace

	response.ValuteCursOnDate = make([]datastructures.GetCursOnDateXMLResultElem, 0)

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
	return nil, response

}
