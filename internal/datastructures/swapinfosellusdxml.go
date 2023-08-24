package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapInfoSellUSDXML struct {
	XMLName  xml.Name `xml:"SwapInfoSellUSDXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapInfoSellUSDXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapInfoSellUSDXML) Validate() error {
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

type SwapInfoSellUSDXMLResult struct {
	// swapinfosellusd node
	SSU []SwapInfoSellUSDXMLResultElem `xml:"SSU" json:"SSU"`
}

type SwapInfoSellUSDXMLResultElem struct {
	DateBuy  time.Time `xml:"DateBuy" json:"DateBuy"`
	DateSell time.Time `xml:"DateSell" json:"DateSell"`
	DateSPOT time.Time `xml:"DateSPOT" json:"DateSPOT"`
	Type     int       `xml:"Type" json:"Type"`
	BaseRate string    `xml:"BaseRate" json:"BaseRate"`
	SD       string    `xml:"SD" json:"SD"`
	TIR      string    `xml:"TIR" json:"TIR"`
	Stavka   string    `xml:"Stavka" json:"Stavka"`
	Limit    string    `xml:"limit" json:"limit"`
}
