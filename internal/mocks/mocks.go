package mocks

import (
	"context"
	"errors"
	"fmt"
	"time"

	customsoap "github.com/skolzkyi/cbrwsdltojson/internal/customsoap"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	"go.uber.org/zap"
)

const cFromDate = "2023-06-22"

const cToDate = "2023-06-23"

var ErrAssertion = errors.New("assertion error")

type ConfigMock struct{}

func (config *ConfigMock) Init(_ string) error {
	return nil
}

func (config *ConfigMock) GetServerURL() string {
	return "localhost:4000"
}

func (config *ConfigMock) GetAddress() string {
	return "localhost"
}

func (config *ConfigMock) GetPort() string {
	return "4000"
}

func (config *ConfigMock) GetServerShutdownTimeout() time.Duration {
	return 30 * time.Second
}

func (config *ConfigMock) GetCBRWSDLTimeout() time.Duration {
	return 5 * time.Second
}

func (config *ConfigMock) GetInfoExpirTime() time.Duration {
	return time.Second
}

func (config *ConfigMock) GetCBRWSDLAddress() string {
	return ""
}

func (config *ConfigMock) GetLoggingOn() bool {
	return true
}

func (config *ConfigMock) GetPermittedRequests() map[string]struct{} {
	return nil
}

type LoggerMock struct {
	loggingOn bool
}

func (l *LoggerMock) GetZapLogger() *zap.SugaredLogger {
	voidSugLogger := zap.SugaredLogger{}
	return &voidSugLogger
}

func NewLoggerMock(loggingOn bool) (*LoggerMock, error) {
	logMock := LoggerMock{}
	logMock.loggingOn = loggingOn
	return &logMock, nil
}

func (l *LoggerMock) Info(msg string) {
	if l.loggingOn {
		fmt.Println("LoggerMock[INFO]: ", msg)
	}
}

func (l *LoggerMock) Warning(msg string) {
	if l.loggingOn {
		fmt.Println("LoggerMock[WARNING]: ", msg)
	}
}

func (l *LoggerMock) Error(msg string) {
	if l.loggingOn {
		fmt.Println("LoggerMock[ERROR]: ", msg)
	}
}

func (l *LoggerMock) Fatal(msg string) {
	if l.loggingOn {
		fmt.Println("LoggerMock[FATAL]: ", msg)
	}
}

type SoapRequestSenderMock struct{}

func (srsm *SoapRequestSenderMock) SoapCall(_ context.Context, action string, input interface{}) ([]byte, error) { // nolint:gocognit, nolintlint, gocyclo, funlen
	switch action {
	case "GetCursOnDateXML":
		inputData, ok := input.(datastructures.GetCursOnDateXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.OnDate == cFromDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><GetCursOnDateXMLResponse xmlns="http://web.cbr.ru/"><GetCursOnDateXMLResult><ValuteData OnDate="20230622" xmlns=""><ValuteCursOnDate><Vname>Австралийский доллар      </Vname><Vnom>1</Vnom><Vcurs>57.1445</Vcurs><Vcode>36</Vcode><VchCode>AUD</VchCode></ValuteCursOnDate><ValuteCursOnDate><Vname>Азербайджанский манат         </Vname><Vnom>1</Vnom><Vcurs>49.5569</Vcurs><Vcode>944</Vcode><VchCode>AZN</VchCode></ValuteCursOnDate></ValuteData></GetCursOnDateXMLResult></GetCursOnDateXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "BiCurBaseXML":
		inputData, ok := input.(datastructures.BiCurBaseXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><BiCurBaseXMLResponse xmlns="http://web.cbr.ru/"><BiCurBaseXMLResult><BiCurBase xmlns=""><BCB><D0>2023-06-22T00:00:00Z</D0><VAL>87.736315</VAL></BCB><BCB><D0>2023-06-23T00:00:00Z</D0><VAL>87.358585</VAL></BCB></BiCurBase></BiCurBaseXMLResult></BiCurBaseXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "BliquidityXML":
		inputData, ok := input.(datastructures.BliquidityXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><BliquidityXMLResponse xmlns="http://web.cbr.ru/"><BliquidityXMLResult><Bliquidity xmlns=""><BL><DT>2023-06-22T00:00:00Z</DT><StrLiDef>-1022.50</StrLiDef><claims>1533.70</claims><actionBasedRepoFX>1378.40</actionBasedRepoFX><actionBasedSecureLoans>0.00</actionBasedSecureLoans><standingFacilitiesRepoFX>0.00</standingFacilitiesRepoFX><standingFacilitiesSecureLoans>155.30</standingFacilitiesSecureLoans><liabilities>-2890.20</liabilities><depositAuctionBased>-1828.30</depositAuctionBased><depositStandingFacilities>-1061.90</depositStandingFacilities><CBRbonds>0.00</CBRbonds><netCBRclaims>334.10</netCBRclaims></BL><BL><DT>2023-06-23T00:00:00Z</DT><StrLiDef>-980.70</StrLiDef><claims>1558.80</claims><actionBasedRepoFX>1378.40</actionBasedRepoFX><actionBasedSecureLoans>0.00</actionBasedSecureLoans><standingFacilitiesRepoFX>0.00</standingFacilitiesRepoFX><standingFacilitiesSecureLoans>180.40</standingFacilitiesSecureLoans><liabilities>-2873.00</liabilities><depositAuctionBased>-1828.30</depositAuctionBased><depositStandingFacilities>-1044.60</depositStandingFacilities><CBRbonds>0.00</CBRbonds><netCBRclaims>333.40</netCBRclaims></BL></Bliquidity></BliquidityXMLResult></BliquidityXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "DepoDynamicXML":
		inputData, ok := input.(datastructures.DepoDynamicXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><DepoDynamicXMLResponse xmlns="http://web.cbr.ru/"><DepoDynamicXMLResult><DepoDynamic xmlns=""><Depo><DateDepo>2023-06-22T00:00:00Z</DateDepo><Overnight>6.50</Overnight></Depo><Depo><DateDepo>2023-06-23T00:00:00Z</DateDepo><Overnight>6.50</Overnight></Depo></DepoDynamic></DepoDynamicXMLResult></DepoDynamicXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "DragMetDynamicXML":
		inputData, ok := input.(datastructures.DragMetDynamicXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><DragMetDynamicXMLResponse xmlns="http://web.cbr.ru/"><DragMetDynamicXMLResult><DragMetall xmlns=""><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>1</CodMet><price>5228.8000</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>2</CodMet><price>64.3800</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>3</CodMet><price>2611.0800</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>4</CodMet><price>3786.6100</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>1</CodMet><price>5176.2400</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>2</CodMet><price>62.0300</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>3</CodMet><price>2550.9600</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>4</CodMet><price>3610.0500</price></DrgMet></DragMetall></DragMetDynamicXMLResult></DragMetDynamicXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "DVXML":
		inputData, ok := input.(datastructures.DVXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><DVXMLResponse xmlns="http://web.cbr.ru/"><DVXMLResult><DV_base xmlns=""><DV><Date>2023-06-22T00:00:00Z</Date><VOvern>0.0000</VOvern><VLomb>9051.4000</VLomb><VIDay>281.3800</VIDay><VOther>504831.8300</VOther><Vol_Gold>0.0000</Vol_Gold><VIDate>2023-06-21T00:00:00Z</VIDate></DV><DV><Date>2023-06-23T00:00:00Z</Date><VOvern>0.0000</VOvern><VLomb>8851.4000</VLomb><VIDay>118.5300</VIDay><VOther>480499.1600</VOther><Vol_Gold>0.0000</Vol_Gold><VIDate>2023-06-22T00:00:00Z</VIDate></DV></DV_base></DVXMLResult></DVXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "EnumReutersValutesXML":
		_, ok := input.(datastructures.EnumReutersValutesXML)
		if !ok {
			return nil, ErrAssertion
		}
		return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><EnumReutersValutesXMLResponse xmlns="http://web.cbr.ru/"><EnumReutersValutesXMLResult><ReutersValutesList xmlns=""><EnumRValutes><num_code>8</num_code><char_code>ALL</char_code><Title_ru>Албанский лек</Title_ru><Title_en>Albanian Lek</Title_en></EnumRValutes><EnumRValutes><num_code>12</num_code><char_code>DZD</char_code><Title_ru>Алжирский динар</Title_ru><Title_en>Algerian Dinar</Title_en></EnumRValutes></ReutersValutesList></EnumReutersValutesXMLResult></EnumReutersValutesXMLResponse></Body></soap:Envelope>`), nil
	case "EnumValutesXML":
		_, ok := input.(datastructures.EnumValutesXML)
		if !ok {
			return nil, ErrAssertion
		}
		return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><EnumReutersValutesXMLResponse xmlns="http://web.cbr.ru/"><EnumValutesXMLResult><ValuteData xmlns=""><EnumValutes><Vcode>R01010</Vcode><Vname>Австралийский доллар</Vname><VEngname>Australian Dollar</VEngname><Vnom>1</Vnom><VcommonCode>R01010</VcommonCode><VnumCode>36</VnumCode><VcharCode>AUD</VcharCode></EnumValutes><EnumValutes><Vcode>R01015</Vcode><Vname>Австрийский шиллинг</Vname><VEngname>Austrian Shilling</VEngname><Vnom>1000</Vnom><VcommonCode>R01015</VcommonCode><VnumCode>40</VnumCode><VcharCode>ATS</VcharCode></EnumValutes></ValuteData></EnumValutesXMLResult></EnumReutersValutesXMLResponse></Body></soap:Envelope>`), nil
	case "KeyRateXML":
		inputData, ok := input.(datastructures.KeyRateXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><KeyRateXMLResponse xmlns="http://web.cbr.ru/"><KeyRateXMLResult><KeyRate xmlns=""><KR><DT>2023-06-22T00:00:00Z</DT><Rate>7.50</Rate></KR><KR><DT>2023-06-23T00:00:00Z</DT><Rate>7.50</Rate></KR></KeyRate></KeyRateXMLResult></KeyRateXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "MainInfoXML":
		_, ok := input.(datastructures.MainInfoXML)
		if !ok {
			return nil, ErrAssertion
		}
		return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><MainInfoXMLResponse xmlns="http://web.cbr.ru/"><MainInfoXMLResult><RegData xmlns=""><keyRate Title="Ключевая ставка" Date="24.07.2023">8.50</keyRate><Inflation Title="Инфляция" Date="01.06.2023">3.25</Inflation><stavka_ref Title="Ставка рефинансирования" Date="24.07.2023">8.50</stavka_ref><GoldBaks Title="Международные резервы" Date="28.07.2023">594</GoldBaks></RegData></MainInfoXMLResult></MainInfoXMLResponse></soap:Body></soap:Envelope>`), nil
	case "mrrf7DXML":
		inputData, ok := input.(datastructures.Mrrf7DXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2023-06-15" && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><mrrf7DXMLResponse xmlns="http://web.cbr.ru/"><mrrf7DXMLResult><mmrf7d xmlns=""><mr><D0>2023-06-16T00:00:00Z</D0><val>587.50</val></mr><mr><D0>2023-06-23T00:00:00Z</D0><val>586.90</val></mr></mmrf7d></mrrf7DXMLResult></mrrf7DXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "mrrfXML":
		inputData, ok := input.(datastructures.MrrfXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2023-05-01" && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><mrrfXMLResponse xmlns="http://web.cbr.ru/"><mrrfXMLResult><mmrf xmlns=""><mr><D0>2023-05-01T00:00:00Z</D0><p1>595787.00</p1><p2>447187.00</p2><p3>418628.00</p3><p4>23559.00</p4><p5>5000.00</p5><p6>148599.00</p6></mr><mr><D0>2023-06-01T00:00:00Z</D0><p1>584175.00</p1><p2>438344.00</p2><p3>410313.00</p3><p4>23127.00</p4><p5>4903.00</p5><p6>145832.00</p6></mr></mmrf></mrrfXMLResult></mrrfXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "NewsInfoXML":
		inputData, ok := input.(datastructures.NewsInfoXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><NewsInfoXMLResponse xmlns="http://web.cbr.ru/"><NewsInfoXMLResult><NewsInfo xmlns=""><News><Doc_id>35498</Doc_id><DocDate>2023-06-22T00:00:00Z</DocDate><Title>О развитии банковского сектора Российской Федерации в мае 2023 года</Title><Url>/analytics/bank_sector/develop/#a_48876</Url></News><News><Doc_id>35495</Doc_id><DocDate>2023-06-22T00:00:00Z</DocDate><Title>Указание Банка России от 10.01.2023 № 6356-У</Title><Url>/Queries/UniDbQuery/File/90134/2803</Url></News></NewsInfo></NewsInfoXMLResult></NewsInfoXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "OmodInfoXML":
		_, ok := input.(datastructures.OmodInfoXML)
		if !ok {
			return nil, ErrAssertion
		}
		return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><OmodInfoXMLResponse xmlns="http://web.cbr.ru/"><OmodInfoXMLResult><OMO Date="05.03.2018" xmlns=""><DirectRepo Time="10:00"><debt>0</debt><rate>0</rate><minrate1D>7.5</minrate1D><minrate7D>7.5</minrate7D></DirectRepo><RevRepo Time="10:00"><debt>0</debt><rate>4.97</rate><sum_debt>0</sum_debt></RevRepo><OBR Time="10:00"><debt>0</debt><rate>3.55</rate></OBR><Deposit>0</Deposit><Credit>0</Credit><VolNom>6741.11</VolNom><TotalFixRepoVol>3132.2</TotalFixRepoVol><FixRepoDate>02.03.2018</FixRepoDate><FixRepo1D><debt>3130.1</debt><rate>8.5</rate></FixRepo1D><FixRepo7D><debt>0</debt><rate>8.5</rate></FixRepo7D><FixRepo1Y><rate>8.5</rate></FixRepo1Y></OMO></OmodInfoXMLResult></OmodInfoXMLResponse></soap:Body></soap:Envelope>`), nil
	case "OstatDepoNewXML":
		inputData, ok := input.(datastructures.OstatDepoNewXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><OstatDepoNewXMLResponse xmlns="http://web.cbr.ru/"><OstatDepoNewXMLResult><OD xmlns=""><odn><DT>2023-06-22T00:00:00Z</DT><TOTAL>2872966.59</TOTAL><AUC_1W>1828340.00</AUC_1W><OV_P>1044626.59</OV_P></odn><odn><DT>2023-06-23T00:00:00Z</DT><TOTAL>2890199.16</TOTAL><AUC_1W>1828340.00</AUC_1W><OV_P>1061859.16</OV_P></odn></OD></OstatDepoNewXMLResult></OstatDepoNewXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "OstatDepoXML":
		inputData, ok := input.(datastructures.OstatDepoXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-12-29" && inputData.ToDate == "2022-12-30" {
			return []byte(`<OstatDepoXMLResult><OD xmlns=""><odr><D0>2022-12-29T00:00:00Z</D0><D1_7>1747362.67</D1_7><D8_30>2515151.15</D8_30><total>4262513.81</total></odr><odr><D0>2022-12-30T00:00:00Z</D0><D1_7>1387715.38</D1_7><D8_30>2515151.15</D8_30><total>3897866.53</total></odr></OD></OstatDepoXMLResult>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "OstatDynamicXML":
		inputData, ok := input.(datastructures.OstatDynamicXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><OstatDynamicXMLResponse xmlns="http://web.cbr.ru/"><OstatDynamicXMLResult><OstatDynamic xmlns=""><Ostat><DateOst>2023-06-22T00:00:00Z</DateOst><InRuss>3756300.00</InRuss><InMoscow>3528600.00</InMoscow></Ostat><Ostat><DateOst>2023-06-23T00:00:00Z</DateOst><InRuss>3688300.00</InRuss><InMoscow>3441000.00</InMoscow></Ostat></OstatDynamic></OstatDynamicXMLResult></OstatDynamicXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "OvernightXML":
		inputData, ok := input.(datastructures.OvernightXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2023-07-22" && inputData.ToDate == "2023-08-16" {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><OvernightXMLResponse xmlns="http://web.cbr.ru/"><OvernightXMLResult><Overnight xmlns=""><OB><date>2023-07-24T00:00:00Z</date><stavka>9.50</stavka></OB><OB><date>2023-08-15T00:00:00Z</date><stavka>13.00</stavka></OB></Overnight></OvernightXMLResult></OvernightXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "Repo_debtXML":
		inputData, ok := input.(datastructures.Repo_debtXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><Repo_debtXMLResponse xmlns="http://web.cbr.ru/"><Repo_debtXMLResult><Repo_debt xmlns=""><RD><Date>2023-06-22T00:00:00Z</Date><debt>1378387.6</debt><debt_auc>1378387.6</debt_auc><debt_fix>0.0</debt_fix></RD><RD><Date>2023-06-23T00:00:00Z</Date><debt>1378379.7</debt><debt_auc>1378379.7</debt_auc><debt_fix>0.0</debt_fix></RD></Repo_debt></Repo_debtXMLResult></Repo_debtXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "RepoDebtUSDXML":
		inputData, ok := input.(datastructures.RepoDebtUSDXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><RepoDebtUSDXMLResponse xmlns="http://web.cbr.ru/"><RepoDebtUSDXMLResult><RepoDebtUSD xmlns=""><rd><D0>2023-06-22T00:00:00Z</D0><TP>0</TP></rd><rd><D0>2023-06-22T00:00:00Z</D0><TP>1</TP></rd><rd><D0>2023-06-23T00:00:00Z</D0><TP>0</TP></rd><rd><D0>2023-06-23T00:00:00Z</D0><TP>1</TP></rd></RepoDebtUSD></RepoDebtUSDXMLResult></RepoDebtUSDXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "ROISfixXML":
		inputData, ok := input.(datastructures.ROISfixXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-27" && inputData.ToDate == "2022-03-02" {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><ROISfixXMLResponse xmlns="http://web.cbr.ru/"><ROISfixXMLResult><ROISfix xmlns=""><rf><D0>2022-02-28T00:00:00Z</D0><R1W>17.83</R1W><R2W>18.00</R2W><R1M>20.65</R1M><R2M>21.96</R2M><R3M>23.23</R3M><R6M>24.52</R6M></rf><rf><D0>2022-03-01T00:00:00Z</D0><R1W>19.85</R1W><R2W>19.91</R2W><R1M>22.63</R1M><R2M>23.79</R2M><R3M>24.49</R3M><R6M>25.71</R6M></rf></ROISfix></ROISfixXMLResult></ROISfixXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "RuoniaSVXML":
		inputData, ok := input.(datastructures.RuoniaSVXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><RuoniaSVXMLResponse xmlns="http://web.cbr.ru/"><RuoniaSVXMLResult><RuoniaSV xmlns=""><ra><DT>2023-06-22T00:00:00Z</DT><RUONIA_Index>2.65003371140540</RUONIA_Index><RUONIA_AVG_1M>7.33031817626889</RUONIA_AVG_1M><RUONIA_AVG_3M>7.28023580262342</RUONIA_AVG_3M><RUONIA_AVG_6M>7.34479164787354</RUONIA_AVG_6M></ra><ra><DT>2023-06-23T00:00:00Z</DT><RUONIA_Index>2.65055282759819</RUONIA_Index><RUONIA_AVG_1M>7.32512579295002</RUONIA_AVG_1M><RUONIA_AVG_3M>7.27890778428907</RUONIA_AVG_3M><RUONIA_AVG_6M>7.34359578515310</RUONIA_AVG_6M></ra></RuoniaSV></RuoniaSVXMLResult></RuoniaSVXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "RuoniaXML":
		inputData, ok := input.(datastructures.RuoniaXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><RuoniaXMLResponse xmlns="http://web.cbr.ru/"><RuoniaXMLResult><Ruonia xmlns=""><ro><D0>2023-06-22T00:00:00Z</D0><ruo>7.1500</ruo><vol>367.9500</vol><DateUpdate>2023-06-23T00:00:00Z</DateUpdate></ro><ro><D0>2023-06-23T00:00:00Z</D0><ruo>7.1300</ruo><vol>388.4500</vol><DateUpdate>2023-06-26T00:00:00Z</DateUpdate></ro></Ruonia></RuoniaXMLResult></RuoniaXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SaldoXML":
		inputData, ok := input.(datastructures.SaldoXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == cFromDate && inputData.ToDate == cToDate {
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SaldoXMLResponse xmlns="http://web.cbr.ru/"><SaldoXMLResult><Saldo xmlns=""><So><Dt>2023-06-22T00:00:00Z</Dt><DEADLINEBS>1044.60</DEADLINEBS></So><So><Dt>2023-06-23T00:00:00Z</Dt><DEADLINEBS>1061.30</DEADLINEBS></So></Saldo></SaldoXMLResult></SaldoXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapDayTotalXML":
		inputData, ok := input.(datastructures.SwapDayTotalXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-25" && inputData.ToDate == "2022-02-28" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapDayTotalXMLResponse xmlns="http://web.cbr.ru/"><SwapDayTotalXMLResult><SwapDayTotal xmlns=""><SDT><DT>2022-02-28T00:00:00Z</DT><Swap>0.0</Swap></SDT><SDT><DT>2022-02-25T00:00:00Z</DT><Swap>24120.4</Swap></SDT></SwapDayTotal></SwapDayTotalXMLResult></SwapDayTotalXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapDynamicXML":
		inputData, ok := input.(datastructures.SwapDynamicXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-25" && inputData.ToDate == "2022-02-28" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapDynamicXMLResponse xmlns="http://web.cbr.ru/"><SwapDynamicXMLResult><SwapDynamic xmlns=""><Swap><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><BaseRate>96.8252</BaseRate><SD>0.0882</SD><TIR>10.5000</TIR><Stavka>-0.576000</Stavka><Currency>1</Currency></Swap><Swap><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><BaseRate>87.1154</BaseRate><SD>0.0748</SD><TIR>10.5000</TIR><Stavka>0.050000</Stavka><Currency>0</Currency></Swap></SwapDynamic></SwapDynamicXMLResult></SwapDynamicXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapInfoSellUSDVolXML":
		inputData, ok := input.(datastructures.SwapInfoSellUSDVolXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-24" && inputData.ToDate == "2022-02-28" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapInfoSellUSDVolXMLResponse xmlns="http://web.cbr.ru/"><SwapInfoSellUSDVolXMLResult><SwapInfoSellUSDVol xmlns=""><SSUV><DT>2022-02-25T00:00:00Z</DT><TODTOMrubvol>435577.0</TODTOMrubvol><TODTOMusdvol>5000.0</TODTOMusdvol><TOMSPTrubvol>128974.3</TOMSPTrubvol><TOMSPTusdvol>1480.5</TOMSPTusdvol></SSUV><SSUV><DT>2022-02-24T00:00:00Z</DT><TODTOMrubvol>403236.5</TODTOMrubvol><TODTOMusdvol>5000.0</TODTOMusdvol><TOMSPTrubvol>32299.2</TOMSPTrubvol><TOMSPTusdvol>400.5</TOMSPTusdvol></SSUV></SwapInfoSellUSDVol></SwapInfoSellUSDVolXMLResult></SwapInfoSellUSDVolXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapInfoSellUSDXML":
		inputData, ok := input.(datastructures.SwapInfoSellUSDXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-25" && inputData.ToDate == "2022-02-28" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapInfoSellUSDXMLResponse xmlns="http://web.cbr.ru/"><SwapInfoSellUSDXMLResult><swapinfosellusd xmlns=""><SSU><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><DateSPOT>2022-03-01T00:00:00Z</DateSPOT><Type>1</Type><BaseRate>87.115400</BaseRate><SD>0.016500</SD><TIR>8.5000</TIR><Stavka>1.5500</Stavka><limit>2.0000</limit></SSU><SSU><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-25T00:00:00Z</DateSell><DateSPOT>2022-02-28T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>87.115400</BaseRate><SD>0.049600</SD><TIR>8.5000</TIR><Stavka>1.5500</Stavka><limit>5.0000</limit></SSU></swapinfosellusd></SwapInfoSellUSDXMLResult></SwapInfoSellUSDXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapInfoSellVolXML":
		inputData, ok := input.(datastructures.SwapInfoSellVolXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2023-05-05" && inputData.ToDate == "2023-05-10" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapInfoSellVolXMLResponse xmlns="http://web.cbr.ru/"><SwapInfoSellVolXMLResult><SwapInfoSellVol xmlns=""><SSUV><DT>2023-05-10T00:00:00Z</DT><Currency>2</Currency><type>0</type><VOL_FC>1113.5</VOL_FC><VOL_RUB>12512.6</VOL_RUB></SSUV><SSUV><DT>2023-05-05T00:00:00Z</DT><Currency>2</Currency><type>0</type><VOL_FC>4583.7</VOL_FC><VOL_RUB>51606.0</VOL_RUB></SSUV></SwapInfoSellVol></SwapInfoSellVolXMLResult></SwapInfoSellVolXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapInfoSellXML":
		inputData, ok := input.(datastructures.SwapInfoSellXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2023-06-20" && inputData.ToDate == "2023-06-21" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapInfoSellXMLResponse xmlns="http://web.cbr.ru/"><SwapInfoSellXMLResult><SwapInfoSell xmlns=""><SSU><Currency>2</Currency><DateBuy>2023-06-21T00:00:00Z</DateBuy><DateSell>2023-06-21T00:00:00Z</DateSell><DateSPOT>2023-06-26T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>11.764246</BaseRate><SD>0.003375</SD><TIR>6.5000</TIR><Stavka>4.3440</Stavka><limit>10.0000</limit></SSU><SSU><Currency>2</Currency><DateBuy>2023-06-20T00:00:00Z</DateBuy><DateSell>2023-06-20T00:00:00Z</DateSell><DateSPOT>2023-06-21T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>11.730496</BaseRate><SD>0.000626</SD><TIR>6.5000</TIR><Stavka>4.4890</Stavka><limit>10.0000</limit></SSU></SwapInfoSell></SwapInfoSellXMLResult></SwapInfoSellXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "SwapMonthTotalXML":
		inputData, ok := input.(datastructures.SwapMonthTotalXML)
		if !ok {
			return nil, ErrAssertion
		}
		if inputData.FromDate == "2022-02-11" && inputData.ToDate == "2022-02-24" { // nolint: goconst, nolintlint
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SwapMonthTotalXMLResponse xmlns="http://web.cbr.ru/"><SwapMonthTotalXMLResult><SwapMonthTotal xmlns=""><SMT><D0>2022-02-11T00:00:00Z</D0><RUB>41208.1</RUB><USD>553.3</USD></SMT><SMT><D0>2022-02-24T00:00:00Z</D0><RUB>24113.5</RUB><USD>299.0</USD></SMT></SwapMonthTotal></SwapMonthTotalXMLResult></SwapMonthTotalXMLResponse></soap:Body></soap:Envelope`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	case "AllDataInfoXML":
		_, ok := input.(datastructures.AllDataInfoXML)
		if !ok {
			return nil, ErrAssertion
		}
		return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><MainInfoXMLResponse xmlns="http://web.cbr.ru/"><AllDataInfoXMLResult><AllData xmlns=""><MainIndicatorsVR Title="Основные индикаторы финансового рынка"><Currency Title="Курсы валют" LUpd=""><USD OnDate="29.08.2023"><curs>95.4717</curs></USD><EUR OnDate="29.08.2023"><curs>103.2434</curs></EUR><CNY OnDate="29.08.2023"><curs>13.0550</curs></CNY></Currency><Metall Title="Драгоценные металлы" LUpd="" OnDate="29.08.2023"><Золото val="5879.60" old_val="5837.5100"></Золото><Серебро val="74.24" old_val="73.6400"></Серебро><Платина val="2912.94" old_val="2841.0300"></Платина><Палладий val="3784.67" old_val="3788.0400"></Палладий></Metall><Inflation Title="Инфляция" LUpd="" OnDate="01.07.2023" val="4.30"></Inflation><InflationTarget Title="Цель по инфляции" LUpd="" OnDate="01.01.2017" val="4.0"></InflationTarget><MBK Title="Ставки межбанковского кредитного рынка" LUpd=""><MIBID OnDate="30.12.2016"><D1 val="9.79" old_val="9.79"></D1><D2_7 val="10.00" old_val="10.00"></D2_7><D8_30 val="9.93" old_val="9.93"></D8_30></MIBID><MIBOR OnDate="30.12.2016"><D1 val="10.54" old_val="10.54"></D1><D2_7 val="10.67" old_val="10.67"></D2_7><D8_30 val="11.06" old_val="11.06"></D8_30></MIBOR><MIACR OnDate="25.08.2023"><D1 val="11.91" old_val="11.91"></D1><D2_7 val="12.39" old_val="10.67"></D2_7><D8_30 val="" old_val=""></D8_30></MIACR><MIACR-IG OnDate="25.08.2023"><D1 val="11.95" old_val="11.95"></D1><D2_7 val="12.39" old_val="12.39"></D2_7><D8_30 val="" old_val=""></D8_30></MIACR-IG></MBK><MosPrime Title="MosPrime Rate" LUpd="" OnDate="01.03.2022"><D1 val="" old_val="20.39"></D1><M1 val="" old_val="20.96"></M1><M3 val="" old_val="20.96"></M3></MosPrime></MainIndicatorsVR><KEY_RATE Title="Действующая ключевая ставка" val="12.00" date="15.08.2023"></KEY_RATE><KEY_RATE_FUTURE Title="Новое значение ключевой ставки (справочно)" val="12.00" newdate="15.08.2023"></KEY_RATE_FUTURE><REF_RATE Title="Ставка рефинансирования (Значение соответствует значению ключевой ставки Банка России)" val="12.00"></REF_RATE><MBRStavki Title="Параметры операций Банка России"><Overnight_rate Title="Ставка по кредиту overnight" LUpd="15.08.2023 11:14:15"><Val1 Date="15.08.2023" val="13.0"></Val1><Val2 Date="" val="8"></Val2></Overnight_rate><FixedLomb Title="Фиксированные cтавки по ломбардным кредитам" LUpd=""><D30 Date="28.04.2014" val="8.50"></D30><D7 Date="28.04.2014" val="8.50"></D7><D1 Date="15.08.2023" val="13.00"></D1></FixedLomb><DepoRates Title="Ставки по депозитным операциям" LUpd="29.08.2023 1:01:09" OnDate="29.08.2023"><TomNext val="" old_val=""></TomNext><SpotNext val="" old_val=""></SpotNext><W1 val="MIACR_B" old_val=""></W1><W1_SPOT val="" old_val=""></W1_SPOT><CallDeposit val="" old_val=""></CallDeposit></DepoRates><SWAP Title="Своп-разница по валютному свопу"><USD_RUB LUpd="" val="" old_val="0.0748"></USD_RUB><EUR_RUB LUpd="" val="" old_val="0.0882"></EUR_RUB></SWAP><FixedRepoRate Title="Фиксированные cтавки по операциям прямого РЕПО"><D1 val="13"></D1><D7 val="13"></D7></FixedRepoRate><MinimalRepoRates Title="Параметры аукционов прямого РЕПО - Минимальные процентные ставки" LUpd="" OnDate="15.08.2023"><D1 val="12"></D1><D7 val="12"></D7></MinimalRepoRates><MaxVolRepoOnAuction Title="Максимальный объем средств, предоставляемых на первом аукционе прямого РЕПО" LUpd="" OnDate="28.09.2015" val="230"></MaxVolRepoOnAuction><MaxVolSwap Title="Максимальный объем средств, предоставляемых по операциям &#39;валютный своп" LUpd="" OnDate="20.09.2016" val="620"></MaxVolSwap></MBRStavki><Ko Title="Требования Банка России к кредитным организациям"><OnOvernightCredit Title="По кредитам overnight" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="0.0" old_val="0.0"></OnOvernightCredit><OnLombardCredit Title="По ломбардным кредитам" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="14348.7" old_val="15348.7"></OnLombardCredit><OnOtherCredit Title="По другим кредитам" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="1744136.5" old_val="874720.8"></OnOtherCredit><OnDirectRepo Title="По операциям прямого РЕПО" OnDate="29.08.2023"><OnAuction Title="на аукционной основе" val="1307685"></OnAuction><OnFixed Title="по фиксированной ставке" val="601"></OnFixed></OnDirectRepo><UnsecLoans Title="По кредитам без обеспечения" LUpd="" OnDate="31.12.2010" val="0" old_val="0"></UnsecLoans></Ko><BankLikvid Title="Показатели банковской ликвидности"><OstatKO Title="Сведения об остатках средств на корреспондентских счетах кредитных организаций" OnDate="29.08.2023" LUpd="29.08.2023 9:04:24"><Russ val="4769.8000" old_val="4356.7000"></Russ><Msk val="4530.5000" old_val="4123.9000"></Msk></OstatKO><InDCredit Title="Объем предоставленных внутридневных кредитов" LUpd="29.08.2023 9:18:46" OnDate="28.08.2023" val="1486.62" old_val="334.55"></InDCredit><DepoBR Title="Депозиты банков в Банке России" LUpd="29.08.2023 9:20:51" OnDate="29.08.2023" val="2368.1896" old_val="2362.4110"></DepoBR><Saldo Title="Сальдо операций Банка России по предоставлению /абсорбированию ликвидности" LUpd="29.08.2023 9:56:14" OnDate="29.08.2023" val="-167.2" old_val="591.7"></Saldo><VolOBR Title="Объем рынка ОБР" val="0"></VolOBR><VolDepo Title="Объем средств федерального бюджета, размещенных на депозитах коммерческих банков" OnDate="05.03.2018" val="0"></VolDepo></BankLikvid><Nor date="28.06.2023" Title="Нормативы обязательных резервов"><Ob_1 Title="по обязательствам перед юридическими лицами – нерезидентами"><Ob_1_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_1><Ob_1_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_2><Ob_1_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_3></Ob_1><Ob_2 Title=""><Ob_2_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_1><Ob_2_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_2><Ob_2_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_3></Ob_2><Ob_3 Title=""><Ob_3_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_1><Ob_3_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_2><Ob_3_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_3></Ob_3><Kor Title="Коэффициент усреднения обязательных резервов"><Ku_1 Title="для банков с универсальной лицензией, банков с базовой лицензией" val="0.9"></Ku_1><Ku_2 Title="для небанковских кредитных организаций" val="1.0"></Ku_2></Kor></Nor><Macro Title="Макроэкономические индикаторы"><DB Title="Денежная база" val="11084.8"></DB><DM Title="Денежная масса (M2)" val="36917.8"></DM><M_rez Title="Международные резервы" val="579.5" date="18.08.2023"></M_rez><Vol_GKO_OFZ Title="Объем рынка ГКО-ОФЗ" val="6741.11"></Vol_GKO_OFZ></Macro></AllData></AllDataInfoXMLResult></MainInfoXMLResponse></soap:Body></soap:Envelope>`), nil
	default:
		return nil, errors.New("SoapRequestSenderMock: unsupported action")
	}
}
