package datastructures

import (
	"encoding/xml"
	"time"
)

type OvernightXML struct {
	XMLName  xml.Name `xml:"OvernightXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *OvernightXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *OvernightXML) Validate() error {
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

type OvernightXMLResult struct {
	// Overnight node
	OB []OvernightXMLResultElem `xml:"OB" json:"OB"`
}

type OvernightXMLResultElem struct {
	Date   time.Time `xml:"date" json:"date"`
	Stavka string    `xml:"stavka" json:"stavka"`
}
