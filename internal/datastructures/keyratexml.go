package datastructures

import (
	"encoding/xml"
	"time"
)

type KeyRateXML struct {
	XMLName  xml.Name `xml:"KeyRateXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *KeyRateXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *KeyRateXML) Validate() error {
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

type KeyRateXMLResult struct {
	// KeyRate node
	KR []KeyRateXMLResultElem `xml:"KR"`
}

type KeyRateXMLResultElem struct {
	DT   time.Time `xml:"DT" json:"DT"`
	Rate string    `xml:"Rate" json:"Rate"`
}
