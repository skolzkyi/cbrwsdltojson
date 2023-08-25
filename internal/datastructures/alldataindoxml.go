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
	KEY_RATE         KEY_RATEElem         `xml:"KEY_RATE" json:"KEY_RATE"`
	KEY_RATE_FUTURE  KEY_RATE_FUTUREElem  `xml:"KEY_RATE_FUTURE" json:"KEY_RATE_FUTURE"`
	REF_RATE         REF_RATEElem         `xml:"REF_RATE" json:"REF_RATE"`
	MBRStavki        MBRStavkiElem        `xml:"MBRStavki" json:"MBRStavki"`
}

type VoVStElem struct {
	Val     string `xml:"val,attr" json:"val"`
	Old_val string `xml:"old_val,attr" json:"old_val"` //nolint:revive, stylecheck, nolintlint
}

type VStElem struct {
	Val string `xml:"val,attr" json:"val"`
}

type MainIndicatorsVRElem struct {
	Title           string              `xml:"Title,attr" json:"Title"`
	Currency        CurrencyElem        `xml:"Currency" json:"Currency"`
	Metall          MetallElem          `xml:"Metall" json:"Metall"`
	Inflation       InflationElemADI    `xml:"Inflation" json:"Inflation"`
	InflationTarget InflationTargetElem `xml:"InflationTarget" json:"InflationTarget"`
	MBK             MBKElem             `xml:"MBK" json:"MBK"`
	MosPrime        MosPrimeElem        `xml:"MosPrime" json:"MosPrime"`
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
	Title     string    `xml:"Title,attr" json:"Title"`
	OnDate    string    `xml:"OnDate,attr" json:"OnDate"`
	LUpd      string    `xml:"LUpd,attr" json:"LUpd"`
	Gold      VoVStElem `xml:"Золото" json:"Gold"`        //nolint:revive, stylecheck, nolintlint
	Silver    VoVStElem `xml:"Серебро" json:"Silver"`     //nolint:revive, stylecheck, nolintlint
	Platinum  VoVStElem `xml:"Платина" json:"Platinum"`   //nolint:revive, stylecheck, nolintlint
	Palladium VoVStElem `xml:"Палладий" json:"Palladium"` //nolint:revive, stylecheck, nolintlint
}

/*
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
}*/

type InflationElemADI struct {
	Title  string `xml:"Title,attr" json:"Title"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	LUpd   string `xml:"LUpd,attr" json:"LUpd"`
	Val    string `xml:"val,attr" json:"val"`
}

type InflationTargetElem struct {
	Title  string `xml:"Title,attr" json:"Title"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	LUpd   string `xml:"LUpd,attr" json:"LUpd"`
	Val    string `xml:"val,attr" json:"val"`
}

type MBKElem struct {
	Title   string        `xml:"Title,attr" json:"Title"`
	LUpd    string        `xml:"LUpd,attr" json:"LUpd"`
	MIBID   MBKStructElem `xml:"MIBID" json:"MIBID"`
	MIBOR   MBKStructElem `xml:"MIBOR" json:"MIBOR"`
	MIACR   MBKStructElem `xml:"MIACR" json:"MIACR"`
	MIACRIG MBKStructElem `xml:"MIACR-IG" json:"MIACRIG"`
}

type MBKStructElem struct {
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`       //nolint:revive, stylecheck, nolintlint
	D2_7   VoVStElem `xml:"D2_7" json:"D2_7"`   //nolint:revive, stylecheck, nolintlint
	D8_30  VoVStElem `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck, nolintlint
}

/*
type MIBIDElem struct {
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`       //nolint:revive, stylecheck, nolintlint
	D2_7   VoVStElem `xml:"D2_7" json:"D2_7"`   //nolint:revive, stylecheck, nolintlint
	D8_30  VoVStElem `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck, nolintlint
}

type MIBORElem struct {
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`       //nolint:revive, stylecheck, nolintlint
	D2_7   VoVStElem `xml:"D2_7" json:"D2_7"`   //nolint:revive, stylecheck, nolintlint
	D8_30  VoVStElem `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck, nolintlint
}

type MIACRElem struct {
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`       //nolint:revive, stylecheck, nolintlint
	D2_7   VoVStElem `xml:"D2_7" json:"D2_7"`   //nolint:revive, stylecheck, nolintlint
	D8_30  VoVStElem `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck, nolintlint
}

type MIACRIGElem struct {
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`       //nolint:revive, stylecheck, nolintlint
	D2_7   VoVStElem `xml:"D2_7" json:"D2_7"`   //nolint:revive, stylecheck, nolintlint
	D8_30  VoVStElem `xml:"D8_30" json:"D8_30"` //nolint:revive, stylecheck, nolintlint
}*/

/*
type D1Elem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type D2_7Elem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type D8_30Elem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}*/

type MosPrimeElem struct {
	Title  string    `xml:"Title,attr" json:"Title"`
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	LUpd   string    `xml:"LUpd,attr" json:"LUpd"`
	D1     VoVStElem `xml:"D1" json:"D1"`
	M1     VoVStElem `xml:"M1" json:"M1"`
	M3     VoVStElem `xml:"M3" json:"M3"`
}

/*
type D1MPElem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type M1Elem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type M3Elem struct {
	Val     string `xml:"val,attr" json:"Val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}*/

type KEY_RATEElem struct { //nolint:revive, stylecheck, nolintlint
	Title string `xml:"Title,attr" json:"Title"`
	Val   string `xml:"val,attr" json:"val"`
	Date  string `xml:"date,attr" json:"date"`
}

type KEY_RATE_FUTUREElem struct { //nolint:revive, stylecheck, nolintlint
	Title   string `xml:"Title,attr" json:"Title"`
	Val     string `xml:"val,attr" json:"val"`
	NewDate string `xml:"newdate,attr" json:"newdate"`
}

type REF_RATEElem struct { //nolint:revive, stylecheck, nolintlint
	Title string `xml:"Title,attr" json:"Title"`
	Val   string `xml:"val,attr" json:"val"`
}

type MBRStavkiElem struct {
	Title               string               `xml:"Title,attr" json:"Title"`
	Overnight_rate      Overnight_rateElem   `xml:"Overnight_rate" json:"Overnight_rate"` //nolint:revive, stylecheck, nolintlint
	FixedLomb           FixedLombElem        `xml:"FixedLomb" json:"FixedLomb"`
	DepoRates           DepoRatesElem        `xml:"DepoRates" json:"DepoRates"`
	SWAP                SWAPElem             `xml:"SWAP" json:"SWAP"`
	FixedRepoRate       FixedRepoRateElem    `xml:"FixedRepoRate" json:"FixedRepoRate"`
	MinimalRepoRates    MinimalRepoRatesElem `xml:"MinimalRepoRates" json:"MinimalRepoRates"`
	MaxVolRepoOnAuction MaxVolMBRelem        `xml:"MaxVolRepoOnAuction" json:"MaxVolRepoOnAuction"`
	MaxVolSwap          MaxVolMBRelem        `xml:"MaxVolSwap" json:"MaxVolSwap"`
}

type Overnight_rateElem struct {
	Title string    `xml:"Title,attr" json:"Title"`
	LUpd  string    `xml:"LUpd,attr" json:"LUpd"`
	Val1  ValORElem `xml:"Val1" json:"Val1"`
	Val2  ValORElem `xml:"Val2" json:"Val2"`
}

type ValORElem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"val"`
}

type FixedLombElem struct {
	Title string `xml:"Title,attr" json:"Title"`
	LUpd  string `xml:"LUpd,attr" json:"LUpd"`
	D30   FLElem `xml:"D30" json:"D30"`
	D1    FLElem `xml:"D1" json:"D1"`
	D7    FLElem `xml:"D7" json:"D7"`
}

type FLElem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"val"`
}

/*
type D30Elem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"Val"`
}

type D1FLElem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"Val"`
}

type D7FLElem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"Val"`
}*/

type DepoRatesElem struct {
	Title       string    `xml:"Title,attr" json:"Title"`
	LUpd        string    `xml:"LUpd,attr" json:"LUpd"`
	OnDate      string    `xml:"OnDate,attr" json:"OnDate"`
	TomNext     VoVStElem `xml:"TomNext" json:"TomNext"`
	SpotNext    VoVStElem `xml:"SpotNext" json:"SpotNext"`
	W1          VoVStElem `xml:"W1" json:"W1"`
	W1_SPOT     VoVStElem `xml:"W1_SPOT" json:"W1_SPOT"` //nolint:revive, stylecheck, nolintlint
	CallDeposit VoVStElem `xml:"CallDeposit" json:"CallDeposit"`
}

type SWAPElem struct {
	Title   string      `xml:"Title,attr" json:"Title"`
	USD_RUB SWAPCurElem `xml:"USD_RUB" json:"USD_RUB"` //nolint:revive, stylecheck, nolintlint
	EUR_RUB SWAPCurElem `xml:"EUR_RUB" json:"EUR_RUB"` //nolint:revive, stylecheck, nolintlint
}

type SWAPCurElem struct {
	LUpd    string `xml:"LUpd,attr" json:"LUpd"`
	Val     string `xml:"val,attr" json:"val"`
	Old_val string `xml:"old_val,attr" json:"Old_val"` //nolint:revive, stylecheck, nolintlint
}

type FixedRepoRateElem struct {
	Title string  `xml:"Title,attr" json:"Title"`
	D1    VStElem `xml:"D1" json:"D1"`
	D7    VStElem `xml:"D7" json:"D7"`
}

type MinimalRepoRatesElem struct {
	Title  string  `xml:"Title,attr" json:"Title"`
	LUpd   string  `xml:"LUpd,attr" json:"LUpd"`
	OnDate string  `xml:"OnDate,attr" json:"OnDate"`
	D1     VStElem `xml:"D1" json:"D1"`
	D7     VStElem `xml:"D7" json:"D7"`
}

type MaxVolMBRelem struct {
	Title  string `xml:"Title,attr" json:"Title"`
	LUpd   string `xml:"LUpd,attr" json:"LUpd"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Val    string `xml:"val,attr" json:"val"`
}

/*
type KeyRateElem struct {
	Title   string `xml:"Title,attr" json:"Title"`
	Date    string `xml:"Date,attr" json:"Date"`
	KeyRate string `xml:",chardata" json:"keyRate"`
}
*/
