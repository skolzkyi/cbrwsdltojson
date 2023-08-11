package datastructures

import (
	"encoding/xml"
	"time"
)

type NewsInfoXML struct {
	XMLName  xml.Name `xml:"NewsInfoXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *NewsInfoXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *NewsInfoXML) Validate() error {
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

type NewsInfoXMLResult struct {
	// NewsInfo node
	News []NewsInfoXMLResultElem `xml:"News"`
}

type NewsInfoXMLResultElem struct {
	Doc_id  int64     `xml:"Doc_id" json:"Doc_id"` //nolint:revive, stylecheck
	DocDate time.Time `xml:"DocDate" json:"DocDate"`
	Title   string    `xml:"Title" json:"Title"`
	Url     string    `xml:"Url" json:"Url"`
}
