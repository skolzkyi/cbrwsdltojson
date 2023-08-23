package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapDayTotalXML struct {
	XMLName  xml.Name `xml:"SwapDayTotalXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapDayTotalXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapDayTotalXML) Validate() error {
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

type SwapDayTotalXMLResult struct {
	// SwapDayTotal node
	SDT []SwapDayTotalXMLResultElem `xml:"SDT" json:"SDT"`
}

type SwapDayTotalXMLResultElem struct {
	DT   time.Time `xml:"DT" json:"DT"`
	Swap string    `xml:"Swap" json:"Swap"`
}
