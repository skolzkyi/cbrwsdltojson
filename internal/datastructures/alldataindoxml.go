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
	KEY_RATE         KEY_RATEElem         `xml:"KEY_RATE" json:"KEY_RATE"`               //nolint:revive, stylecheck, nolintlint
	KEY_RATE_FUTURE  KEY_RATE_FUTUREElem  `xml:"KEY_RATE_FUTURE" json:"KEY_RATE_FUTURE"` //nolint:revive, stylecheck, nolintlint
	REF_RATE         TVStElem             `xml:"REF_RATE" json:"REF_RATE"`               //nolint:revive, stylecheck, nolintlint
	MBRStavki        MBRStavkiElem        `xml:"MBRStavki" json:"MBRStavki"`
	Ko               KoElem               `xml:"Ko" json:"Ko"`
	BankLikvid       BankLikvidElem       `xml:"BankLikvid" json:"BankLikvid"`
	Nor              NorElem              `xml:"Nor" json:"Nor"`
	Macro            MacroElem            `xml:"Macro" json:"Macro"`
}

type VoVStElem struct {
	Val     string `xml:"val,attr" json:"val"`
	Old_val string `xml:"old_val,attr" json:"old_val"` //nolint:revive, stylecheck, nolintlint
}

type VStElem struct {
	Val string `xml:"val,attr" json:"val"`
}

type TVStElem struct {
	Title string `xml:"Title,attr" json:"Title"`
	Val   string `xml:"val,attr" json:"val"`
}

type TLOVOStElem struct {
	Title   string `xml:"Title,attr" json:"Title"`
	LUpd    string `xml:"LUpd,attr" json:"LUpd"`
	OnDate  string `xml:"OnDate,attr" json:"OnDate"`
	Val     string `xml:"val,attr" json:"val"`
	Old_val string `xml:"old_val,attr" json:"old_val"` //nolint:revive, stylecheck, nolintlint
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
	LUpd      string    `xml:"LUpd,attr" json:"LUpd"`
	OnDate    string    `xml:"OnDate,attr" json:"OnDate"`
	Gold      VoVStElem `xml:"Золото" json:"Gold"`        //nolint:revive, stylecheck, nolintlint
	Silver    VoVStElem `xml:"Серебро" json:"Silver"`     //nolint:revive, stylecheck, nolintlint
	Platinum  VoVStElem `xml:"Платина" json:"Platinum"`   //nolint:revive, stylecheck, nolintlint
	Palladium VoVStElem `xml:"Палладий" json:"Palladium"` //nolint:revive, stylecheck, nolintlint
}

type InflationElemADI struct {
	Title  string `xml:"Title,attr" json:"Title"`
	LUpd   string `xml:"LUpd,attr" json:"LUpd"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Val    string `xml:"val,attr" json:"val"`
}

type InflationTargetElem struct {
	Title  string `xml:"Title,attr" json:"Title"`
	LUpd   string `xml:"LUpd,attr" json:"LUpd"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
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

type MosPrimeElem struct {
	Title  string    `xml:"Title,attr" json:"Title"`
	LUpd   string    `xml:"LUpd,attr" json:"LUpd"`
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	D1     VoVStElem `xml:"D1" json:"D1"`
	M1     VoVStElem `xml:"M1" json:"M1"`
	M3     VoVStElem `xml:"M3" json:"M3"`
}

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

type Overnight_rateElem struct { //nolint:revive, stylecheck, nolintlint
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
	D7    FLElem `xml:"D7" json:"D7"`
	D1    FLElem `xml:"D1" json:"D1"`
}

type FLElem struct {
	Date string `xml:"Date,attr" json:"Date"`
	Val  string `xml:"val,attr" json:"val"`
}

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

type KoElem struct {
	Title             string           `xml:"Title,attr" json:"Title"`
	OnOvernightCredit TLOVOStElem      `xml:"OnOvernightCredit" json:"OnOvernightCredit"`
	OnLombardCredit   TLOVOStElem      `xml:"OnLombardCredit" json:"OnLombardCredit"`
	OnOtherCredit     TLOVOStElem      `xml:"OnOtherCredit" json:"OnOtherCredit"`
	OnDirectRepo      OnDirectRepoElem `xml:"OnDirectRepo" json:"OnDirectRepo"`
	UnsecLoans        TLOVOStElem      `xml:"UnsecLoans" json:"UnsecLoans"`
}

type OnDirectRepoElem struct {
	Title     string   `xml:"Title,attr" json:"Title"`
	OnDate    string   `xml:"OnDate,attr" json:"OnDate"`
	OnAuction TVStElem `xml:"OnAuction" json:"OnAuction"`
	OnFixed   TVStElem `xml:"OnFixed" json:"OnFixed"`
}

type BankLikvidElem struct {
	Title     string      `xml:"Title,attr" json:"Title"`
	OstatKO   OstatKOElem `xml:"OstatKO" json:"OstatKO"`
	InDCredit TLOVOStElem `xml:"InDCredit" json:"InDCredit"`
	DepoBR    TLOVOStElem `xml:"DepoBR" json:"DepoBR"`
	Saldo     TLOVOStElem `xml:"Saldo" json:"Saldo"`
	VolOBR    TVStElem    `xml:"VolOBR" json:"VolOBR"`
	VolDepo   VolDepoElem `xml:"VolDepo" json:"VolDepo"`
}

type OstatKOElem struct {
	Title  string    `xml:"Title,attr" json:"Title"`
	OnDate string    `xml:"OnDate,attr" json:"OnDate"`
	LUpd   string    `xml:"LUpd,attr" json:"LUpd"`
	Russ   VoVStElem `xml:"Russ" json:"Russ"`
	Msk    VoVStElem `xml:"Msk" json:"Msk"`
}

type VolDepoElem struct {
	Title  string `xml:"Title,attr" json:"Title"`
	OnDate string `xml:"OnDate,attr" json:"OnDate"`
	Val    string `xml:"val,attr" json:"val"`
}

type NorElem struct {
	Date  string   `xml:"date,attr" json:"date"`
	Title string   `xml:"Title,attr" json:"Title"`
	Ob_1  Ob_1Elem `xml:"Ob_1" json:"Ob_1"` //nolint:revive, stylecheck, nolintlint
	Ob_2  Ob_2Elem `xml:"Ob_2" json:"Ob_2"` //nolint:revive, stylecheck, nolintlint
	Ob_3  Ob_3Elem `xml:"Ob_3" json:"Ob_3"` //nolint:revive, stylecheck, nolintlint
	Kor   KorElem  `xml:"Kor" json:"Kor"`
}

type Ob_1Elem struct { //nolint:revive, stylecheck, nolintlint
	Title  string        `xml:"Title,attr" json:"Title"`
	Ob_1_1 NorTLevelelem `xml:"Ob_1_1" json:"Ob_1_1"` //nolint:revive, stylecheck, nolintlint
	Ob_1_2 NorTLevelelem `xml:"Ob_1_2" json:"Ob_1_2"` //nolint:revive, stylecheck, nolintlint
	Ob_1_3 NorTLevelelem `xml:"Ob_1_3" json:"Ob_1_3"` //nolint:revive, stylecheck, nolintlint
}

type Ob_2Elem struct { //nolint:revive, stylecheck, nolintlint
	Title  string        `xml:"Title,attr" json:"Title"`
	Ob_2_1 NorTLevelelem `xml:"Ob_2_1" json:"Ob_2_1"` //nolint:revive, stylecheck, nolintlint
	Ob_2_2 NorTLevelelem `xml:"Ob_2_2" json:"Ob_2_2"` //nolint:revive, stylecheck, nolintlint
	Ob_2_3 NorTLevelelem `xml:"Ob_2_3" json:"Ob_2_3"` //nolint:revive, stylecheck, nolintlint
}

type Ob_3Elem struct { //nolint:revive, stylecheck, nolintlint
	Title  string        `xml:"Title,attr" json:"Title"`
	Ob_3_1 NorTLevelelem `xml:"Ob_3_1" json:"Ob_3_1"` //nolint:revive, stylecheck, nolintlint
	Ob_3_2 NorTLevelelem `xml:"Ob_3_2" json:"Ob_3_2"` //nolint:revive, stylecheck, nolintlint
	Ob_3_3 NorTLevelelem `xml:"Ob_3_3" json:"Ob_3_3"` //nolint:revive, stylecheck, nolintlint
}

type NorTLevelelem struct {
	Title            string `xml:"Title,attr" json:"Title"`
	Val_rub          string `xml:"val_rub,attr" json:"val_rub"`                   //nolint:revive, stylecheck, nolintlint
	Val_usd          string `xml:"val_usd,attr" json:"val_usd"`                   //nolint:revive, stylecheck, nolintlint
	Val_usd_excludUC string `xml:"val_usd_excludUC,attr" json:"val_usd_excludUC"` //nolint:revive, stylecheck, nolintlint
}

type KorElem struct {
	Title string   `xml:"Title,attr" json:"Title"`
	Ku_1  TVStElem `xml:"Ku_1" json:"Ku_1"` //nolint:revive, stylecheck, nolintlint
	Ku_2  TVStElem `xml:"Ku_2" json:"Ku_2"` //nolint:revive, stylecheck, nolintlint
}

type MacroElem struct {
	Title       string    `xml:"Title,attr" json:"Title"`
	DB          TVStElem  `xml:"DB" json:"DB"`
	DM          TVStElem  `xml:"DM" json:"DM"`
	M_rez       M_rezElem `xml:"M_rez" json:"M_rez"`             //nolint:revive, stylecheck, nolintlint
	Vol_GKO_OFZ TVStElem  `xml:"Vol_GKO_OFZ" json:"Vol_GKO_OFZ"` //nolint:revive, stylecheck, nolintlint
}

type M_rezElem struct { //nolint:revive, stylecheck, nolintlint
	Title string `xml:"Title,attr" json:"Title"`
	Val   string `xml:"val,attr" json:"val"`
	Date  string `xml:"date,attr" json:"date"`
}
