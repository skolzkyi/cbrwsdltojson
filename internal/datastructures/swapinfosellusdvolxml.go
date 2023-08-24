package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapInfoSellUSDVolXML struct {
	XMLName  xml.Name `xml:"SwapInfoSellUSDVolXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapInfoSellUSDVolXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapInfoSellUSDVolXML) Validate() error {
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

type SwapInfoSellUSDVolXMLResult struct {
	// SwapInfoSellUSDVol node
	SSUV []SwapInfoSellUSDVolXMLResultElem `xml:"SSUV" json:"SSUV"`
}

type SwapInfoSellUSDVolXMLResultElem struct {
	DT           time.Time `xml:"DT" json:"DT"`
	TODTOMrubvol string    `xml:"TODTOMrubvol" json:"TODTOMrubvol"`
	TODTOMusdvol string    `xml:"TODTOMusdvol" json:"TODTOMusdvol"`
	TOMSPTrubvol string    `xml:"TOMSPTrubvol" json:"TOMSPTrubvol"`
	TOMSPTusdvol string    `xml:"TOMSPTusdvol" json:"TOMSPTusdvol"`
}
