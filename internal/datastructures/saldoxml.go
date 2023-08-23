package datastructures

import (
	"encoding/xml"
	"time"
)

type SaldoXML struct {
	XMLName  xml.Name `xml:"SaldoXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SaldoXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SaldoXML) Validate() error {
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

type SaldoXMLResult struct {
	// Saldo node
	So []SaldoXMLResultElem `xml:"So" json:"So"`
}

type SaldoXMLResultElem struct {
	Dt         time.Time `xml:"Dt" json:"Dt"`
	DEADLINEBS string    `xml:"DEADLINEBS" json:"DEADLINEBS"`
}
