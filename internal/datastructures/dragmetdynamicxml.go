package datastructures

import (
	"encoding/xml"
	"time"
)

type DragMetDynamicXML struct {
	XMLName  xml.Name `xml:"DragMetDynamicXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *DragMetDynamicXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *DragMetDynamicXML) Validate() error {
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

type DragMetDynamicXMLResult struct {
	// DragMetall node
	DrgMet []DragMetDynamicXMLResultElem `xml:"DrgMet"`
}

type DragMetDynamicXMLResultElem struct {
	DateMet time.Time `xml:"DateMet" json:"DateMet"`
	CodMet  string    `xml:"CodMet" json:"CodMet"`
	Price   string    `xml:"price" json:"price"`
}
