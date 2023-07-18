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

func (config *ConfigMock) GetDateTimeResponseLayout() string {
	return "2006-01-02"
}

func (config *ConfigMock) GetDateTimeRequestLayout() string {
	return "2006-01-02"
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

func (srsm *SoapRequestSenderMock) SoapCall(_ context.Context, action string, input interface{}) ([]byte, error) {
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
	default:
		return nil, errors.New("SoapRequestSenderMock: unsupported action")
	}
}
