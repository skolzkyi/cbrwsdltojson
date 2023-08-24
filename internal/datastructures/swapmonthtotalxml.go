package datastructures

import (
	"encoding/xml"
	"time"
)

type SwapMonthTotalXML struct {
	XMLName  xml.Name `xml:"SwapMonthTotalXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *SwapMonthTotalXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *SwapMonthTotalXML) Validate() error {
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

type SwapMonthTotalXMLResult struct {
	// SwapMonthTotal node
	SMT []SwapMonthTotalXMLResultElem `xml:"SMT" json:"SMT"`
}

type SwapMonthTotalXMLResultElem struct {
	D0  time.Time `xml:"D0" json:"D0"`
	RUB string    `xml:"RUB" json:"RUB"`
	USD string    `xml:"USD" json:"USD"`
}
