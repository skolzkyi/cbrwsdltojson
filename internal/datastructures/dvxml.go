package datastructures

import (
	"encoding/xml"
	"time"
)

type DVXML struct {
	XMLName  xml.Name `xml:"DVXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *DVXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *DVXML) Validate() error {
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

type DVXMLResult struct {
	// DV_base node
	DV []DVXMLResultElem `xml:"DV"`
}

type DVXMLResultElem struct {
	Date     time.Time `xml:"Date" json:"Date"`
	VOvern   string    `xml:"VOvern" json:"VOvern"`
	VLomb    string    `xml:"VLomb" json:"VLomb"`
	VIDay    string    `xml:"VIDay" json:"VIDay"`
	VOther   string    `xml:"VOther" json:"VOther"`
	Vol_Gold string    `xml:"Vol_Gold" json:"Vol_Gold"`
	VIDate   time.Time `xml:"VIDate" json:"VIDate"`
}
