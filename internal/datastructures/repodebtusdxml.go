package datastructures

import (
	"encoding/xml"
	"time"
)

type RepoDebtUSDXML struct {
	XMLName  xml.Name `xml:"RepoDebtUSDXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *RepoDebtUSDXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *RepoDebtUSDXML) Validate() error {
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

type RepoDebtUSDXMLResult struct {
	// RepoDebtUSD node
	Rd []RepoDebtUSDXMLResultElem `xml:"rd" json:"rd"`
}

type RepoDebtUSDXMLResultElem struct {
	D0 time.Time `xml:"D0" json:"D0"`
	TP int       `xml:"TP" json:"TP"`
}
