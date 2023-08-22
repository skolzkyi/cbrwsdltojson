package datastructures

import (
	"encoding/xml"
	"time"
)

type OstatDynamicXML struct {
	XMLName  xml.Name `xml:"OstatDynamicXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *OstatDynamicXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *OstatDynamicXML) Validate() error {
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

type OstatDynamicXMLResult struct {
	// OstatDynamic node
	Ostat []OstatDynamicXMLResultElem `xml:"Ostat" json:"Ostat"`
}

type OstatDynamicXMLResultElem struct {
	DateOst  time.Time `xml:"DateOst" json:"DateOst"`
	InRuss   string    `xml:"InRuss" json:"InRuss"`
	InMoscow string    `xml:"InMoscow" json:"InMoscow"`
}
