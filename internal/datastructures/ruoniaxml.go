package datastructures

import (
	"encoding/xml"
	"time"
)

type RuoniaXML struct {
	XMLName  xml.Name `xml:"RuoniaXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *RuoniaXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *RuoniaXML) Validate() error {
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

type RuoniaXMLResult struct {
	// Ruonia node
	Ro []RuoniaXMLResultElem `xml:"ro" json:"ro"`
}

type RuoniaXMLResultElem struct {
	D0         time.Time `xml:"D0" json:"D0"`
	Ruo        string    `xml:"ruo" json:"ruo"`
	Vol        string    `xml:"vol" json:"vol"`
	DateUpdate time.Time `xml:"DateUpdate" json:"DateUpdate"`
}
