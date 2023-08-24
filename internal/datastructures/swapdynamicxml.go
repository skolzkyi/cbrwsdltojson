package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapDynamicXML struct {
	XMLName  xml.Name `xml:"SwapDynamicXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapDynamicXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapDynamicXML) Validate() error {
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

type SwapDynamicXMLResult struct {
	// SwapDynamic node
	Swap []SwapDynamicXMLResultElem `xml:"Swap" json:"Swap"`
}

type SwapDynamicXMLResultElem struct {
	DateBuy  time.Time `xml:"DateBuy" json:"DateBuy"`
	DateSell time.Time `xml:"DateSell" json:"DateSell"`
	BaseRate string    `xml:"BaseRate" json:"BaseRate"`
	SD       string    `xml:"SD" json:"SD"`
	TIR      string    `xml:"TIR" json:"TIR"`
	Stavka   string    `xml:"Stavka" json:"Stavka"`
	Currency int       `xml:"Currency" json:"Currency"`
}
