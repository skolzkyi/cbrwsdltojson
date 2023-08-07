package datastructures

import (
	"encoding/xml"
)

type EnumValutesXML struct {
	XMLName xml.Name `xml:"EnumValutesXML" json:"-"`
	XMLNs   string   `xml:"xmlns,attr" json:"-"`
	Seld    bool     `xml:"Seld"`
}

func (data *EnumValutesXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *EnumValutesXML) Validate() error {
	return nil
}

type EnumValutesXMLResult struct {
	// ValuteData node
	EnumValutes []EnumValutesXMLResultElem `xml:"EnumValutes"`
}

type EnumValutesXMLResultElem struct {
	Vcode       string `xml:"Vcode" json:"Vcode"`
	Vname       string `xml:"Vname" json:"Vname"`
	VEngname    string `xml:"VEngname" json:"VEngname"`
	Vnom        int32  `xml:"Vnom" json:"Vnom"`
	VcommonCode string `xml:"VcommonCode" json:"VcommonCode"`
	VnumCode    int32  `xml:"VnumCode" json:"VnumCode"`
	VcharCode   string `xml:"VcharCode" json:"VcharCode"`
}
