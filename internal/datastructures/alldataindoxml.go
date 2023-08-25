package datastructures

import (
	"encoding/xml"
)

type AllDataInfoXML struct {
	XMLName xml.Name `xml:"AllDataInfoXML" json:"-"`
	XMLNs   string   `xml:"xmlns,attr" json:"-"`
}

func (data *AllDataInfoXML) Init() {
	data.XMLNs = cbrNamespace
}

type AllDataInfoXMLResult struct {
	// AllData node
	MainIndicatorsVR MainIndicatorsVRElem `xml:"MainIndicatorsVR" json:"MainIndicatorsVR"`
}

type MainIndicatorsVRElem struct {
	Title    string       `xml:"Title,attr" json:"Title"`
	Currency CurrencyElem `xml:"Currency" json:"Currency"`
	Metall   MetallElem   `xml:"Metall" json:"Metall"`
}

type CurrencyElem struct {
	Title string  `xml:"Title,attr" json:"Title"`
	LUpd  string  `xml:"LUpd,attr" json:"LUpd"`
	USD   USDElem `xml:"USD" json:"USD"`
	EUR   EURElem `xml:"EUR" json:"EUR"`
	CNY   CNYElem `xml:"CNY" json:"CNY"`
}

type USDElem struct {
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Curs   string `xml:"curs" json:"curs"`
}

type EURElem struct {
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Curs   string `xml:"curs" json:"curs"`
}

type CNYElem struct {
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Curs   string `xml:"curs" json:"curs"`
}

type MetallElem struct {
	Title     string        `xml:"Title,attr" json:"Title"`
	OnDate    string        `xml:"OnDate,attr" json:"OnDate"`
	LUpd      string        `xml:"LUpd,attr" json:"LUpd"`
	Gold      GoldElem      `xml:"Золото" json:"Gold"`        //nolint:revive, stylecheck, nolintlint
	Silver    SilverElem    `xml:"Серебро" json:"Silver"`     //nolint:revive, stylecheck, nolintlint
	Platinum  PlatinumElem  `xml:"Платина" json:"Platinum"`   //nolint:revive, stylecheck, nolintlint
	Palladium PalladiumElem `xml:"Палладий" json:"Palladium"` //nolint:revive, stylecheck, nolintlint
}

type GoldElem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type SilverElem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type PlatinumElem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type PalladiumElem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

/*
type KeyRateElem struct {
	Title   string `xml:"Title,attr" json:"Title"`
	Date    string `xml:"Date,attr" json:"Date"`
	KeyRate string `xml:",chardata" json:"keyRate"`
}
*/
