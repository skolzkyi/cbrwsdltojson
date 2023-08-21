package datastructures

import (
	"encoding/xml"
	"time"
)

type OstatDepoNewXML struct {
	XMLName  xml.Name `xml:"OstatDepoNewXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *OstatDepoNewXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *OstatDepoNewXML) Validate() error {
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

type OstatDepoNewXMLResult struct {
	// OD node
	Odn []OstatDepoNewXMLResultElem `xml:"odn" json:"odn"`
}

type OstatDepoNewXMLResultElem struct {
	DT     time.Time `xml:"DT" json:"DT"`
	TOTAL  string    `xml:"TOTAL" json:"TOTAL"`
	AUC_1W string    `xml:"AUC_1W" json:"AUC_1W"` //nolint:revive, stylecheck
	OV_P   string    `xml:"OV_P" json:"OV_P"`     //nolint:revive, stylecheck
}
