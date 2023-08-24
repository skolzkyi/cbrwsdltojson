package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapInfoSellVolXML struct {
	XMLName  xml.Name `xml:"SwapInfoSellVolXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapInfoSellVolXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapInfoSellVolXML) Validate() error {
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

type SwapInfoSellVolXMLResult struct {
	// SwapInfoSellVol node
	SSUV []SwapInfoSellVolXMLResultElem `xml:"SSUV" json:"SSUV"`
}

type SwapInfoSellVolXMLResultElem struct {
	DT       time.Time `xml:"DT" json:"DT"`
	Currency int       `xml:"Currency" json:"Currency"`
	Type     int       `xml:"type" json:"type"`
	VOL_FC   string    `xml:"VOL_FC" json:"VOL_FC"`   //nolint:revive, stylecheck, nolintlint
	VOL_RUB  string    `xml:"VOL_RUB" json:"VOL_RUB"` //nolint:revive, stylecheck, nolintlint
}
