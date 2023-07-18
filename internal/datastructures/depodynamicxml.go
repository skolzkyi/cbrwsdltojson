package datastructures

import (
	"encoding/xml"
	"time"
)

type DepoDynamicXML struct {
	XMLName  xml.Name `xml:"DepoDynamicXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *DepoDynamicXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *DepoDynamicXML) Validate() error {
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

type DepoDynamicXMLResult struct {
	// DepoDynamic node
	Depo []DepoDynamicXMLResultElem `xml:"Depo"`
}

type DepoDynamicXMLResultElem struct {
	DateDepo  time.Time `xml:"DateDepo" json:"DateDepo"`
	Overnight string    `xml:"Overnight" json:"Overnight"`
}
