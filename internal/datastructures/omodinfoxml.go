package datastructures

import (
	"encoding/xml"
)

type OmodInfoXML struct {
	XMLName xml.Name `xml:"OmodInfoXML" json:"-"`
	XMLNs   string   `xml:"xmlns,attr" json:"-"`
}

func (data *OmodInfoXML) Init() {
	data.XMLNs = cbrNamespace
}

type OmodInfoXMLResult struct {
	// OMO node
	Date            string         `xml:"Date,attr" json:"Date"`
	DirectRepo      DirectRepoElem `xml:"DirectRepo" json:"DirectRepo"`
	RevRepo         RevRepoElem    `xml:"RevRepo" json:"RevRepo"`
	OBR             OBRElem        `xml:"OBR" json:"OBR"`
	Deposit         string         `xml:"Deposit" json:"Deposit"`
	Credit          string         `xml:"Credit" json:"Credit"`
	VolNom          string         `xml:"VolNom" json:"VolNom"`
	TotalFixRepoVol string         `xml:"TotalFixRepoVol" json:"TotalFixRepoVol"`
	FixRepoDate     string         `xml:"FixRepoDate" json:"FixRepoDate"`
	FixRepo1D       FixRepo1DElem  `xml:"FixRepo1D" json:"FixRepo1D"`
	FixRepo7D       FixRepo7DElem  `xml:"FixRepo7D" json:"FixRepo7D"`
	FixRepo1Y       FixRepo1YElem  `xml:"FixRepo1Y" json:"FixRepo1Y"`
}

type DirectRepoElem struct {
	Time      string `xml:"Time,attr" json:"Time"`
	Debt      string `xml:"debt" json:"debt"`
	Rate      string `xml:"rate" json:"rate"`
	Minrate1D string `xml:"minrate1D" json:"minrate1D"`
	Minrate7D string `xml:"minrate7D" json:"minrate7D"`
}

type RevRepoElem struct {
	Time     string `xml:"Time,attr" json:"Time"`
	Debt     string `xml:"debt" json:"debt"`
	Rate     string `xml:"rate" json:"rate"`
	Sum_debt string `xml:"sum_debt" json:"sum_debt"` //nolint:revive, stylecheck
}

type OBRElem struct {
	Time string `xml:"Time,attr" json:"Time"`
	Debt string `xml:"debt" json:"debt"`
	Rate string `xml:"rate" json:"rate"`
}

type FixRepo1DElem struct {
	Debt string `xml:"debt" json:"debt"`
	Rate string `xml:"rate" json:"rate"`
}

type FixRepo7DElem struct {
	Debt string `xml:"debt" json:"debt"`
	Rate string `xml:"rate" json:"rate"`
}

type FixRepo1YElem struct {
	Rate string `xml:"rate" json:"rate"`
}
