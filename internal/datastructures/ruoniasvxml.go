package datastructures

import (
	"encoding/xml"
	"time"
)

type RuoniaSVXML struct {
	XMLName  xml.Name `xml:"RuoniaSVXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *RuoniaSVXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *RuoniaSVXML) Validate() error {
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

type RuoniaSVXMLResult struct {
	// RuoniaSV node
	Ra []RuoniaSVXMLResultElem `xml:"ra" json:"ra"`
}

type RuoniaSVXMLResultElem struct {
	DT            time.Time `xml:"DT" json:"DT"`
	RUONIA_Index  string    `xml:"RUONIA_Index" json:"RUONIA_Index"`   //nolint:revive, stylecheck, nolintlint
	RUONIA_AVG_1M string    `xml:"RUONIA_AVG_1M" json:"RUONIA_AVG_1M"` //nolint:revive, stylecheck, nolintlint
	RUONIA_AVG_3M string    `xml:"RUONIA_AVG_3M" json:"RUONIA_AVG_3M"` //nolint:revive, stylecheck, nolintlint
	RUONIA_AVG_6M string    `xml:"RUONIA_AVG_6M" json:"RUONIA_AVG_6M"` //nolint:revive, stylecheck, nolintlint
}
