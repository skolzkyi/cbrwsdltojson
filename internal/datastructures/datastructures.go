package datastructures

import (
	"encoding/xml"
	"errors"
	"time"
)

var (
	ErrBadInputDateData = errors.New("fromDate after toDate")
	ErrBadRawData       = errors.New("parse raw date error")
)

type GetCursOnDateXML struct {
	XMLName xml.Name `xml:"GetCursOnDateXML"`
	XMLNs   string   `xml:"xmlns,attr"`
	OnDate  string   `xml:"On_date"`
}

func (data *GetCursOnDateXML) Validate(inputDTLayout string) error {
	_, err := time.Parse(inputDTLayout, data.OnDate)
	if err != nil {
		return ErrBadRawData
	}
	return nil
}

type RequestBetweenDate struct {
	FromDate string
	ToDate   string
}

func (data *RequestBetweenDate) Validate(inputDTLayout string) error {
	fromDateDate, err := time.Parse(inputDTLayout, data.FromDate)
	if err != nil {
		return ErrBadRawData
	}
	toDateDate, err := time.Parse(inputDTLayout, data.ToDate)
	if err != nil {
		return ErrBadRawData
	}
	if fromDateDate.After(toDateDate) {
		return ErrBadInputDateData
	}
	return nil
}

type RequestSeld struct {
	Seld bool
}

type RequestGetCursDynamic struct {
	FromDate   string
	ToDate     string
	ValutaCode string
}

func (data *RequestGetCursDynamic) Validate(inputDTLayout string) error {
	fromDateDate, err := time.Parse(inputDTLayout, data.FromDate)
	if err != nil {
		return err
	}
	toDateDate, err := time.Parse(inputDTLayout, data.ToDate)
	if err != nil {
		return err
	}
	if fromDateDate.After(toDateDate) {
		return ErrBadInputDateData
	}
	return nil
}

type GetCursOnDateXMLResult struct {
	// ValuteData node
	OnDate           string                       `xml:"OnDate,attr"`
	ValuteCursOnDate []GetCursOnDateXMLResultElem `xml:"ValuteCursOnDate"`
}

type GetCursOnDateXMLResultElem struct {
	Vname   string `xml:"Vname"`
	Vnom    int32  `xml:"Vnom"`
	Vcurs   string `xml:"Vcurs"`
	Vcode   string `xml:"Vcode"`
	VchCode string `xml:"VchCode"`
}
