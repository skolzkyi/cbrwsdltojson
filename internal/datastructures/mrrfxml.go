package datastructures

import (
	"encoding/xml"
	"time"
)

type MrrfXML struct {
	XMLName  xml.Name `xml:"mrrfXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *MrrfXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *MrrfXML) Validate() error {
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

type MrrfXMLResult struct {
	// mmrf7d node
	Mr []MrrfXMLResultElem `xml:"mr" json:"mr"`
}

type MrrfXMLResultElem struct {
	D0 time.Time `xml:"D0" json:"D0"`
	P1 string    `xml:"p1" json:"p1"`
	P2 string    `xml:"p2" json:"p2"`
	P3 string    `xml:"p3" json:"p3"`
	P4 string    `xml:"p4" json:"p4"`
	P5 string    `xml:"p5" json:"p5"`
	P6 string    `xml:"p6" json:"p6"`
}
