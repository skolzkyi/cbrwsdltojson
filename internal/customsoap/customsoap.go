package customsoap

// based on: https://fale.io/blog/2018/12/03/calling-a-soap-service-in-go

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

var ErrBadLenEnvelopeSlice = errors.New("Bad length of slice with element of envelope string")

type CBRSOAPSender struct {
	InclLogger Logger
	InclConfig Config
	HTTPClient http.Client
	WSAddress  string
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

type soapRQ struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XMLNsXsi  string   `xml:"xmlns:xsi,attr"`
	XMLNsXsd  string   `xml:"xmlns:xsd,attr"`
	XMLNsSoap string   `xml:"xmlns:soap,attr"`
	Body      soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Desc    string   `xml:"BODY_DESCRIPTOR"`
}

func New(logger Logger, config Config) *CBRSOAPSender {
	CBRSOAPSender := CBRSOAPSender{
		InclLogger: logger,
		InclConfig: config,
		HTTPClient: http.Client{
			Timeout: config.GetCBRWSDLTimeout(),
		},
		WSAddress: config.GetCBRWSDLAddress(),
	}
	return &CBRSOAPSender
}

func (soapSender *CBRSOAPSender) SoapCall(action string, payload interface{}) ([]byte, error) {
	bodyRequest := make([]byte, 0)
	bodyInnXML, err := xml.MarshalIndent(payload, "", "  ")
	if err != nil {
		soapSender.InclLogger.Error(err.Error())
		return nil, err
	}
	fmt.Println("bodyInnXML: ", string(bodyInnXML))
	v := soapRQ{
		XMLNsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
		XMLNsXsd:  "http://www.w3.org/2001/XMLSchema",
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	payloadBase, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		soapSender.InclLogger.Error(err.Error())
		return nil, err
	}

	payloadBaseSl := strings.Split(string(payloadBase), "<BODY_DESCRIPTOR></BODY_DESCRIPTOR>")
	if len(payloadBaseSl) != 2 {
		soapSender.InclLogger.Error(err.Error())
		return nil, ErrBadLenEnvelopeSlice
	}
	bodyRequest = append(bodyRequest, []byte(payloadBaseSl[0])...)
	bodyRequest = append(bodyRequest, bodyInnXML...)
	bodyRequest = append(bodyRequest, []byte(payloadBaseSl[1])...)

	req, err := http.NewRequest("POST", soapSender.InclConfig.GetCBRWSDLAddress(), bytes.NewBuffer(bodyRequest))
	if err != nil {
		soapSender.InclLogger.Error(err.Error())
		return nil, err
	}

	req.Header.Set("SOAPAction", `"`+`http://web.cbr.ru/`+action+`"`)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	response, err := soapSender.HTTPClient.Do(req)
	if err != nil {
		soapSender.InclLogger.Error(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		soapSender.InclLogger.Error(err.Error())
		return nil, err
	}

	return bodyBytes, nil
}
