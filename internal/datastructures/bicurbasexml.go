package datastructures

import (
	"encoding/xml"
	"time"
)

type BiCurBaseXML struct {
	XMLName  xml.Name `xml:"BiCurBaseXML"`
	XMLNs    string   `xml:"xmlns,attr"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *BiCurBaseXML) Validate(inputDTLayout string) error {
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

type BiCurBaseXMLResult struct {
	// BiCurBase node
	BCB []BiCurBaseXMLResultElem `xml:"BCB"`
}

type BiCurBaseXMLResultElem struct {
	D0  time.Time `xml:"D0"`
	VAL string    `xml:"VAL"`
}
