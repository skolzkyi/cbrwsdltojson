package datastructures

import (
	"encoding/xml"
)

type MainInfoXML struct {
	XMLName xml.Name `xml:"MainInfoXML" json:"-"`
	XMLNs   string   `xml:"xmlns,attr" json:"-"`
}

func (data *MainInfoXML) Init() {
	data.XMLNs = cbrNamespace
}

type MainInfoXMLResult struct {
	// RegData node
	KeyRate    keyRateElem    `xml:"keyRate" json:"keyRate"`
	Inflation  InflationElem  `xml:"Inflation" json:"Inflation"`
	Stavka_ref stavka_refElem `xml:"stavka_ref" json:"stavka_ref"`
	GoldBaks   GoldBaksElem   `xml:"GoldBaks" json:"GoldBaks"`
}

type keyRateElem struct {
	Title   string `xml:"Title,attr" json:"Title"`
	Date    string `xml:"Date,attr" json:"Date"`
	KeyRate string `xml:",chardata" json:"keyRate"`
}

type InflationElem struct {
	Title     string `xml:"Title,attr" json:"Title"`
	Date      string `xml:"Date,attr" json:"Date"`
	Inflation string `xml:",chardata" json:"Inflation"`
}

type stavka_refElem struct {
	Title      string `xml:"Title,attr" json:"Title"`
	Date       string `xml:"Date,attr" json:"Date"`
	Stavka_ref string `xml:",chardata" json:"stavka_ref"`
}

type GoldBaksElem struct {
	Title    string `xml:"Title,attr" json:"Title"`
	Date     string `xml:"Date,attr" json:"Date"`
	GoldBaks int32  `xml:",chardata" json:"GoldBaks"`
}
