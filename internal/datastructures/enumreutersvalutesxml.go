package datastructures

import (
	"encoding/xml"
)

type EnumReutersValutesXML struct {
	XMLName xml.Name `xml:"EnumReutersValutesXML" json:"-"`
	XMLNs   string   `xml:"xmlns,attr" json:"-"`
}

func (data *EnumReutersValutesXML) Init() {
	data.XMLNs = cbrNamespace
}

type EnumReutersValutesXMLResult struct {
	// ReutersValutesList node
	EnumRValutes []EnumReutersValutesXMLResultElem `xml:"EnumRValutes"`
}

type EnumReutersValutesXMLResultElem struct {
	Num_code  int32  `xml:"num_code" json:"num_code"`
	Char_code string `xml:"char_code" json:"char_code"`
	Title_ru  string `xml:"Title_ru" json:"Title_ru"`
	Title_en  string `xml:"Title_en" json:"Title_en"`
}
