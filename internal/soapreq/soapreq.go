package soapreq

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	gosoap "github.com/tiaguinho/gosoap"
)

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

type SoapRequestSender struct {
	address    string
	config     Config
	httpClient *http.Client
	soap       *gosoap.Client
}

func NewSoapRequestSender() SoapRequestSender {
	return SoapRequestSender{}
}

func (srs *SoapRequestSender) Init(config Config) error {
	var err error
	srs.config = config
	srs.address = config.GetCBRWSDLAddress()
	srs.httpClient = &http.Client{
		Timeout: config.GetCBRWSDLTimeout(),
	}
	srs.soap, err = gosoap.SoapClient(config.GetCBRWSDLAddress(), srs.httpClient)
	if err != nil {
		return err
	}
	return nil
}

func (srs *SoapRequestSender) GetCursOnDate(ctx context.Context, input datastructures.RequestOnDate) (error, datastructures.ResponseValuteCursDynamic) {
	var response datastructures.ResponseValuteCursDynamic
	onDateDate, err := time.Parse(srs.config.GetDateTimeRequestLayout(), input.OnDate)
	if err != nil {
		fmt.Println("1: ", err.Error())
		return err, response
	}
	onDatestrFormatForWSDL := onDateDate.Format("2006-01-02")
	params := gosoap.Params{
		"On_date": onDatestrFormatForWSDL,
	}
	fmt.Println("gosoap.Params: ", params)
	res, err := srs.soap.Call("GetCursOnDateXML", params)
	if err != nil {
		fmt.Println("2: ", err.Error())
		return err, response
	}
	fmt.Println("res: ", res)
	var GetCursOnDateXMLResponse string

	res.Unmarshal(&GetCursOnDateXMLResponse)
	err = xml.Unmarshal([]byte(GetCursOnDateXMLResponse), &response)
	if err != nil {
		fmt.Println("3: ", err.Error())
		return err, response
	}

	return nil, response
}
