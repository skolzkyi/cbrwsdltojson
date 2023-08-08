package datastructures

import (
	"encoding/xml"
	"time"
)

type Mrrf7DXML struct {
	XMLName  xml.Name `xml:"mrrf7DXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *Mrrf7DXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *Mrrf7DXML) Validate() error {
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

type Mrrf7DXMLResult struct {
	// mmrf7d node
	Mr []Mrrf7DXMLResultElem `xml:"mr" json:"mr"`
}

type Mrrf7DXMLResultElem struct {
	D0  time.Time `xml:"D0" json:"D0"`
	Val string    `xml:"val" json:"val"`
}
