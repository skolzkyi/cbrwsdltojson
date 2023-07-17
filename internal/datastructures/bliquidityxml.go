package datastructures

import (
	"encoding/xml"
	"time"
)

type BliquidityXML struct {
	XMLName  xml.Name `xml:"BliquidityXML" json:"-"`
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *BliquidityXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *BliquidityXML) Validate(inputDTLayout string) error {
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

type BliquidityXMLResult struct {
	// Bliquidity node
	BL []BliquidityXMLResultElem `xml:"BL"`
}

type BliquidityXMLResultElem struct {
	DT                            time.Time `xml:"DT" json:"DT"`
	StrLiDef                      string    `xml:"StrLiDef" json:"StrLiDef"`
	Claims                        string    `xml:"claims" json:"claims"`
	ActionBasedRepoFX             string    `xml:"actionBasedRepoFX" json:"actionBasedRepoFX"`
	ActionBasedSecureLoans        string    `xml:"actionBasedSecureLoans" json:"actionBasedSecureLoans"`
	StandingFacilitiesRepoFX      string    `xml:"standingFacilitiesRepoFX" json:"standingFacilitiesRepoFX"`
	StandingFacilitiesSecureLoans string    `xml:"standingFacilitiesSecureLoans" json:"standingFacilitiesSecureLoans"`
	Liabilities                   string    `xml:"liabilities" json:"liabilities"`
	DepositAuctionBased           string    `xml:"depositAuctionBased" json:"depositAuctionBased"`
	DepositStandingFacilities     string    `xml:"depositStandingFacilities" json:"depositStandingFacilities"`
	CBRbonds                      string    `xml:"CBRbonds" json:"CBRbonds"`
	NetCBRclaims                  string    `xml:"netCBRclaims" json:"netCBRclaims"`
}
