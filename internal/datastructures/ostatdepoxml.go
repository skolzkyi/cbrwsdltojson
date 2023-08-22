package datastructures

import (
	"encoding/xml"
	"time"
)

type OstatDepoXML struct {
	XMLName  xml.Name `xml:"OstatDepoXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *OstatDepoXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *OstatDepoXML) Validate() error {
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

type OstatDepoXMLResult struct {
	// OD node
	Odr []OstatDepoXMLResultElem `xml:"odr" json:"odr"`
}

type OstatDepoXMLResultElem struct {
	D0    time.Time `xml:"D0" json:"D0"`
	D1_7  string    `xml:"D1_7" json:"D1_7"`   //nolint:revive, stylecheck
	D8_30 string    `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck
	Total string    `xml:"total" json:"total"` //nolint:revive, stylecheck
}
