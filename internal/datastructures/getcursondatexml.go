package datastructures

import (
	"encoding/xml"
	"time"
)

type GetCursOnDateXML struct {
	XMLName xml.Name `xml:"GetCursOnDateXML"`
	XMLNs   string   `xml:"xmlns,attr"`
	OnDate  string   `xml:"On_date"`
}

func (data *GetCursOnDateXML) Validate(inputDTLayout string) error {
	_, err := time.Parse(inputDTLayout, data.OnDate)
	if err != nil {
		return ErrBadRawData
	}
	return nil
}

type GetCursOnDateXMLResult struct {
	// ValuteData node
	OnDate           string                       `xml:"OnDate,attr"`
	ValuteCursOnDate []GetCursOnDateXMLResultElem `xml:"ValuteCursOnDate"`
	//InfoDTStamp      time.Time                    `json:"-" xml:"-"`
}

type GetCursOnDateXMLResultElem struct {
	Vname   string `xml:"Vname"`
	Vnom    int32  `xml:"Vnom"`
	Vcurs   string `xml:"Vcurs"`
	Vcode   string `xml:"Vcode"`
	VchCode string `xml:"VchCode"`
}
