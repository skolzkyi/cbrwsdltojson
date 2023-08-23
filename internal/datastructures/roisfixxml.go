package datastructures

import (
	"encoding/xml"
	"time"
)

type ROISfixXML struct {
	XMLName  xml.Name `xml:"ROISfixXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *ROISfixXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *ROISfixXML) Validate() error {
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

type ROISfixXMLResult struct {
	// ROISfix node
	Rf []ROISfixXMLResultElem `xml:"rf" json:"rf"`
}

type ROISfixXMLResultElem struct {
	D0  time.Time `xml:"D0" json:"D0"`
	R1W string    `xml:"R1W" json:"R1W"`
	R2W string    `xml:"R2W" json:"R2W"`
	R1M string    `xml:"R1M" json:"R1M"`
	R2M string    `xml:"R2M" json:"R2M"`
	R3M string    `xml:"R3M" json:"R3M"`
	R6M string    `xml:"R6M" json:"R6M"`
}
