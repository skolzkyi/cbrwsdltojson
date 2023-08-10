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
	return 3 * time.Second
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

func (srsm *SoapRequestSenderMock) SoapCall(_ context.Context, action string, input interface{}) ([]byte, error) { // nolint:gocognit, nolintlint, gocyclo
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
			return []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><mrrfXMLResponse xmlns="http://web.cbr.ru/"><mrrfXMLResult><mr><D0>2023-05-01T00:00:00Z</D0><p1>595787.00</p1><p2>447187.00</p2><p3>418628.00</p3><p4>23559.00</p4><p5>5000.00</p5><p6>148599.00</p6></mr><mr><D0>2023-06-01T00:00:00Z</D0><p1>584175.00</p1><p2>438344.00</p2><p3>410313.00</p3><p4>23127.00</p4><p5>4903.00</p5><p6>145832.00</p6></mr></mrrfXMLResult></mrrfXMLResponse></soap:Body></soap:Envelope>`), nil
		}
		return nil, customsoap.ErrContextWSReqExpired
	default:
		return nil, errors.New("SoapRequestSenderMock: unsupported action")
	}
}
