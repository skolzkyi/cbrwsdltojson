package datastructures_test

import (
	"encoding/xml"
	"testing"
	"time"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	"github.com/stretchr/testify/require"
)

type AllDatastructuresTestTable []DatastructuresTestTable

type DatastructuresTestTable struct {
	MethodName      string
	InputDataCases  []DatastructuresTestCase
	OutputDataCases []DatastructuresTestCase
}

type DatastructuresTestCase struct {
	MarshalXMLTestFunc      func(t *testing.T, Datastructure interface{}, XMLMarshalControl string)
	ValidateControlTestFunc func(t *testing.T, Datastructure interface{}, ValidateControl error)
	Datastructure           interface{}
	Name                    string
	DataStructureType       string
	XMLMarshalControl       string
	ValidateControl         error
	NeedXMLMarshal          bool
	NeedValidate            bool
}

func (dtc *DatastructuresTestCase) MarshalXMLTest(t *testing.T) {
	t.Helper()
	if dtc.NeedXMLMarshal {
		dtc.MarshalXMLTestFunc(t, dtc.Datastructure, dtc.XMLMarshalControl)
	}
}

func (dtc *DatastructuresTestCase) ValidateControlTest(t *testing.T) {
	t.Helper()
	if dtc.NeedXMLMarshal {
		dtc.ValidateControlTestFunc(t, dtc.Datastructure, dtc.ValidateControl)
	}
}

func initAllDatastructuresTestTable(t *testing.T) AllDatastructuresTestTable {
	t.Helper()
	AllDTTable := make(AllDatastructuresTestTable, 32)
	AllDTTable[0] = initTestCasesGetCursOnDateXML(t)
	AllDTTable[1] = initTestCasesBiCurBaseXML(t)
	AllDTTable[2] = initTestCasesBliquidityXML(t)
	AllDTTable[3] = initTestCasesDepoDynamicXML(t)
	AllDTTable[4] = initTestCasesDragMetDynamicXML(t)
	AllDTTable[5] = initTestCasesDVXML(t)
	AllDTTable[6] = initTestCasesEnumReutersValutesXML(t)
	AllDTTable[7] = initTestCasesEnumValutesXML(t)
	AllDTTable[8] = initTestCasesKeyRateXML(t)
	AllDTTable[9] = initTestCasesMainInfoXML(t)
	AllDTTable[10] = initTestCasesMrrf7DXML(t)
	AllDTTable[11] = initTestCasesMrrfXML(t)
	AllDTTable[12] = initTestCasesNewsInfoXML(t)
	AllDTTable[13] = initTestCasesOmodInfoXML(t)
	AllDTTable[14] = initTestCasesOstatDepoNewXML(t)
	AllDTTable[15] = initTestCasesOstatDepoXML(t)
	AllDTTable[16] = initTestCasesOstatDynamicXML(t)
	AllDTTable[17] = initTestCasesOvernightXML(t)
	AllDTTable[18] = initTestCasesRepoDebtXML(t)
	AllDTTable[19] = initTestCasesRepoDebtUSDXML(t)
	AllDTTable[20] = initTestCasesROISfixXML(t)
	AllDTTable[21] = initTestCasesRuoniaSVXML(t)
	AllDTTable[22] = initTestCasesRuoniaXML(t)
	AllDTTable[23] = initTestCasesSaldoXML(t)
	AllDTTable[24] = initTestCasesSwapDayTotalXML(t)
	AllDTTable[25] = initTestCasesSwapDynamicXML(t)
	AllDTTable[26] = initTestCasesSwapInfoSellUSDVolXML(t)
	AllDTTable[27] = initTestCasesSwapInfoSellUSDXML(t)
	AllDTTable[28] = initTestCasesSwapInfoSellVolXML(t)
	AllDTTable[29] = initTestCasesSwapInfoSellXML(t)
	AllDTTable[30] = initTestCasesSwapMonthTotalXML(t)
	AllDTTable[31] = initTestCasesAllDataInfoXML(t)
	return AllDTTable
}

func initTestCasesGetCursOnDateXML(t *testing.T) DatastructuresTestTable {
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "GetCursOnDateXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 3)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "GetCursOnDateXML",
		Datastructure: datastructures.GetCursOnDateXML{
			OnDate: "2023-06-22",
			XMLNs:  "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<GetCursOnDateXML xmlns="http://web.cbr.ru/"><On_date>2023-06-22</On_date></GetCursOnDateXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.GetCursOnDateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:GetCursOnDateXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegative",
		DataStructureType: "GetCursOnDateXML",
		Datastructure: datastructures.GetCursOnDateXML{
			OnDate: "022-14-22",
			XMLNs:  "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.GetCursOnDateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:GetCursOnDateXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "GetCursOnDateXML",
		Datastructure: datastructures.GetCursOnDateXML{
			OnDate: "2023-06-22",
			XMLNs:  "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.GetCursOnDateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:GetCursOnDateXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	testGetCursOnDateXMLResult := datastructures.GetCursOnDateXMLResult{
		OnDate:           "20230622",
		ValuteCursOnDate: make([]datastructures.GetCursOnDateXMLResultElem, 2),
	}
	testGetCursOnDateXMLResultElem := datastructures.GetCursOnDateXMLResultElem{
		Vname:   "Австралийский доллар",
		Vnom:    1,
		Vcurs:   "57.1445",
		Vcode:   "36",
		VchCode: "AUD",
	}
	testGetCursOnDateXMLResult.ValuteCursOnDate[0] = testGetCursOnDateXMLResultElem
	testGetCursOnDateXMLResultElem = datastructures.GetCursOnDateXMLResultElem{
		Vname:   "Азербайджанский манат",
		Vnom:    1,
		Vcurs:   "49.5569",
		Vcode:   "944",
		VchCode: "AZN",
	}
	testGetCursOnDateXMLResult.ValuteCursOnDate[1] = testGetCursOnDateXMLResultElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "GetCursOnDateXMLResult",
		Datastructure:     testGetCursOnDateXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<GetCursOnDateXMLResult OnDate="20230622"><ValuteCursOnDate><Vname>Австралийский доллар</Vname><Vnom>1</Vnom><Vcurs>57.1445</Vcurs><Vcode>36</Vcode><VchCode>AUD</VchCode></ValuteCursOnDate><ValuteCursOnDate><Vname>Азербайджанский манат</Vname><Vnom>1</Vnom><Vcurs>49.5569</Vcurs><Vcode>944</Vcode><VchCode>AZN</VchCode></ValuteCursOnDate></GetCursOnDateXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.GetCursOnDateXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:GetCursOnDateXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesBiCurBaseXML(t *testing.T) DatastructuresTestTable {
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "BiCurBaseXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "BiCurBaseXML",
		Datastructure: datastructures.BiCurBaseXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<BiCurBaseXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></BiCurBaseXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BiCurBaseXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "BiCurBaseXML",
		Datastructure: datastructures.BiCurBaseXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BiCurBaseXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "BiCurBaseXML",
		Datastructure: datastructures.BiCurBaseXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BiCurBaseXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "BiCurBaseXML",
		Datastructure: datastructures.BiCurBaseXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BiCurBaseXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testBiCurBaseXMLResult := datastructures.BiCurBaseXMLResult{
		BCB: make([]datastructures.BiCurBaseXMLResultElem, 2),
	}
	testBiCurBaseXMLResultElem := datastructures.BiCurBaseXMLResultElem{
		D0:  time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		VAL: "87.736315",
	}
	testBiCurBaseXMLResult.BCB[0] = testBiCurBaseXMLResultElem
	testBiCurBaseXMLResultElem = datastructures.BiCurBaseXMLResultElem{
		D0:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		VAL: "87.358585",
	}
	testBiCurBaseXMLResult.BCB[1] = testBiCurBaseXMLResultElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "BiCurBaseXMLResult",
		Datastructure:     testBiCurBaseXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<BiCurBaseXMLResult><BCB><D0>2023-06-22T00:00:00Z</D0><VAL>87.736315</VAL></BCB><BCB><D0>2023-06-23T00:00:00Z</D0><VAL>87.358585</VAL></BCB></BiCurBaseXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BiCurBaseXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesBliquidityXML(t *testing.T) DatastructuresTestTable {
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "BliquidityXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "BliquidityXML",
		Datastructure: datastructures.BliquidityXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<BliquidityXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></BliquidityXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BliquidityXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BliquidityXML1")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "BliquidityXML",
		Datastructure: datastructures.BliquidityXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BliquidityXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BliquidityXML2")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "BliquidityXML",
		Datastructure: datastructures.BliquidityXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BliquidityXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BliquidityXML3")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "BliquidityXML",
		Datastructure: datastructures.BliquidityXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BliquidityXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BliquidityXML4")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testBliquidityXML := datastructures.BliquidityXMLResult{
		BL: make([]datastructures.BliquidityXMLResultElem, 2),
	}
	testBliquidityXMLElem := datastructures.BliquidityXMLResultElem{
		DT:                            time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		StrLiDef:                      "-1022.50",
		Claims:                        "1533.70",
		ActionBasedRepoFX:             "1378.40",
		ActionBasedSecureLoans:        "0.00",
		StandingFacilitiesRepoFX:      "0.00",
		StandingFacilitiesSecureLoans: "155.30",
		Liabilities:                   "-2890.20",
		DepositAuctionBased:           "-1828.30",
		DepositStandingFacilities:     "-1061.90",
		CBRbonds:                      "0.00",
		NetCBRclaims:                  "334.10",
	}
	testBliquidityXML.BL[0] = testBliquidityXMLElem
	testBliquidityXMLElem = datastructures.BliquidityXMLResultElem{
		DT:                            time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		StrLiDef:                      "-980.70",
		Claims:                        "1558.80",
		ActionBasedRepoFX:             "1378.40",
		ActionBasedSecureLoans:        "0.00",
		StandingFacilitiesRepoFX:      "0.00",
		StandingFacilitiesSecureLoans: "180.40",
		Liabilities:                   "-2873.00",
		DepositAuctionBased:           "-1828.30",
		DepositStandingFacilities:     "-1044.60",
		CBRbonds:                      "0.00",
		NetCBRclaims:                  "333.40",
	}
	testBliquidityXML.BL[1] = testBliquidityXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "BliquidityXML",
		Datastructure:     testBliquidityXML,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<BliquidityXMLResult><BL><DT>2023-06-22T00:00:00Z</DT><StrLiDef>-1022.50</StrLiDef><claims>1533.70</claims><actionBasedRepoFX>1378.40</actionBasedRepoFX><actionBasedSecureLoans>0.00</actionBasedSecureLoans><standingFacilitiesRepoFX>0.00</standingFacilitiesRepoFX><standingFacilitiesSecureLoans>155.30</standingFacilitiesSecureLoans><liabilities>-2890.20</liabilities><depositAuctionBased>-1828.30</depositAuctionBased><depositStandingFacilities>-1061.90</depositStandingFacilities><CBRbonds>0.00</CBRbonds><netCBRclaims>334.10</netCBRclaims></BL><BL><DT>2023-06-23T00:00:00Z</DT><StrLiDef>-980.70</StrLiDef><claims>1558.80</claims><actionBasedRepoFX>1378.40</actionBasedRepoFX><actionBasedSecureLoans>0.00</actionBasedSecureLoans><standingFacilitiesRepoFX>0.00</standingFacilitiesRepoFX><standingFacilitiesSecureLoans>180.40</standingFacilitiesSecureLoans><liabilities>-2873.00</liabilities><depositAuctionBased>-1828.30</depositAuctionBased><depositStandingFacilities>-1044.60</depositStandingFacilities><CBRbonds>0.00</CBRbonds><netCBRclaims>333.40</netCBRclaims></BL></BliquidityXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.BliquidityXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BliquidityXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesDepoDynamicXML(t *testing.T) DatastructuresTestTable {
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "DepoDynamicXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "DepoDynamicXML",
		Datastructure: datastructures.DepoDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DepoDynamicXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></DepoDynamicXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DepoDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DepoDynamicXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "DepoDynamicXML",
		Datastructure: datastructures.DepoDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DepoDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DepoDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "DepoDynamicXML",
		Datastructure: datastructures.DepoDynamicXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DepoDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DepoDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "DepoDynamicXML",
		Datastructure: datastructures.DepoDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DepoDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DepoDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testDepoDynamicXMLResult := datastructures.DepoDynamicXMLResult{
		Depo: make([]datastructures.DepoDynamicXMLResultElem, 2),
	}
	testDepoDynamicXMLElem := datastructures.DepoDynamicXMLResultElem{
		DateDepo:  time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Overnight: "6.50",
	}
	testDepoDynamicXMLResult.Depo[0] = testDepoDynamicXMLElem
	testDepoDynamicXMLElem = datastructures.DepoDynamicXMLResultElem{
		DateDepo:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Overnight: "6.50",
	}
	testDepoDynamicXMLResult.Depo[1] = testDepoDynamicXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "DepoDynamicXML",
		Datastructure:     testDepoDynamicXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DepoDynamicXMLResult><Depo><DateDepo>2023-06-22T00:00:00Z</DateDepo><Overnight>6.50</Overnight></Depo><Depo><DateDepo>2023-06-23T00:00:00Z</DateDepo><Overnight>6.50</Overnight></Depo></DepoDynamicXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DepoDynamicXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DepoDynamicXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesDragMetDynamicXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "DragMetDynamicXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "DragMetDynamicXML",
		Datastructure: datastructures.DragMetDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DragMetDynamicXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></DragMetDynamicXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DragMetDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DragMetDynamicXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "DepoDynamicXML",
		Datastructure: datastructures.DragMetDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DragMetDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DragMetDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "DragMetDynamicXML",
		Datastructure: datastructures.DragMetDynamicXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DragMetDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DragMetDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "DragMetDynamicXML",
		Datastructure: datastructures.DragMetDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DragMetDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DragMetDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testDragMetDynamicXMLResult := datastructures.DragMetDynamicXMLResult{
		DrgMet: make([]datastructures.DragMetDynamicXMLResultElem, 8),
	}
	testDragMetDynamicXMLElem := datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		CodMet:  "1",
		Price:   "5228.8000",
	}
	testDragMetDynamicXMLResult.DrgMet[0] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		CodMet:  "2",
		Price:   "64.3800",
	}
	testDragMetDynamicXMLResult.DrgMet[1] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		CodMet:  "3",
		Price:   "2611.0800",
	}
	testDragMetDynamicXMLResult.DrgMet[2] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		CodMet:  "4",
		Price:   "3786.6100",
	}
	testDragMetDynamicXMLResult.DrgMet[3] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		CodMet:  "1",
		Price:   "5176.2400",
	}
	testDragMetDynamicXMLResult.DrgMet[4] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		CodMet:  "2",
		Price:   "62.0300",
	}
	testDragMetDynamicXMLResult.DrgMet[5] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		CodMet:  "3",
		Price:   "2550.9600",
	}
	testDragMetDynamicXMLResult.DrgMet[6] = testDragMetDynamicXMLElem
	testDragMetDynamicXMLElem = datastructures.DragMetDynamicXMLResultElem{
		DateMet: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		CodMet:  "4",
		Price:   "3610.0500",
	}
	testDragMetDynamicXMLResult.DrgMet[7] = testDragMetDynamicXMLElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "DragMetDynamicXML",
		Datastructure:     testDragMetDynamicXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DragMetDynamicXMLResult><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>1</CodMet><price>5228.8000</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>2</CodMet><price>64.3800</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>3</CodMet><price>2611.0800</price></DrgMet><DrgMet><DateMet>2023-06-22T00:00:00Z</DateMet><CodMet>4</CodMet><price>3786.6100</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>1</CodMet><price>5176.2400</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>2</CodMet><price>62.0300</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>3</CodMet><price>2550.9600</price></DrgMet><DrgMet><DateMet>2023-06-23T00:00:00Z</DateMet><CodMet>4</CodMet><price>3610.0500</price></DrgMet></DragMetDynamicXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DragMetDynamicXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DragMetDynamicXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesDVXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "DVXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "DVXML",
		Datastructure: datastructures.DVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DVXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></DVXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DVXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "DVXML",
		Datastructure: datastructures.DVXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "DVXML",
		Datastructure: datastructures.DVXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "DVXML",
		Datastructure: datastructures.DVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testDVXMLResult := datastructures.DVXMLResult{
		DV: make([]datastructures.DVXMLResultElem, 2),
	}
	testDVXMLElem := datastructures.DVXMLResultElem{
		Date:     time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		VOvern:   "0.0000",
		VLomb:    "9051.4000",
		VIDay:    "281.3800",
		VOther:   "504831.8300",
		Vol_Gold: "0.0000",
		VIDate:   time.Date(2023, time.June, 21, 0, 0, 0, 0, time.UTC),
	}
	testDVXMLResult.DV[0] = testDVXMLElem
	testDVXMLElem = datastructures.DVXMLResultElem{
		Date:     time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		VOvern:   "0.0000",
		VLomb:    "8851.4000",
		VIDay:    "118.5300",
		VOther:   "480499.1600",
		Vol_Gold: "0.0000",
		VIDate:   time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
	}
	testDVXMLResult.DV[1] = testDVXMLElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "DVXML",
		Datastructure:     testDVXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<DVXMLResult><DV><Date>2023-06-22T00:00:00Z</Date><VOvern>0.0000</VOvern><VLomb>9051.4000</VLomb><VIDay>281.3800</VIDay><VOther>504831.8300</VOther><Vol_Gold>0.0000</Vol_Gold><VIDate>2023-06-21T00:00:00Z</VIDate></DV><DV><Date>2023-06-23T00:00:00Z</Date><VOvern>0.0000</VOvern><VLomb>8851.4000</VLomb><VIDay>118.5300</VIDay><VOther>480499.1600</VOther><Vol_Gold>0.0000</Vol_Gold><VIDate>2023-06-22T00:00:00Z</VIDate></DV></DVXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.DVXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:DVXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesEnumReutersValutesXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "EnumReutersValutesXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 0)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	testEnumReutersValutesXML := datastructures.EnumReutersValutesXMLResult{
		EnumRValutes: make([]datastructures.EnumReutersValutesXMLResultElem, 2),
	}
	testEnumReutersValutesXMLElem := datastructures.EnumReutersValutesXMLResultElem{
		Num_code:  8,
		Char_code: "ALL",
		Title_ru:  "Албанский лек",
		Title_en:  "Albanian Lek",
	}
	testEnumReutersValutesXML.EnumRValutes[0] = testEnumReutersValutesXMLElem
	testEnumReutersValutesXMLElem = datastructures.EnumReutersValutesXMLResultElem{
		Num_code:  12,
		Char_code: "DZD",
		Title_ru:  "Алжирский динар",
		Title_en:  "Algerian Dinar",
	}
	testEnumReutersValutesXML.EnumRValutes[1] = testEnumReutersValutesXMLElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "EnumReutersValutesXML",
		Datastructure:     testEnumReutersValutesXML,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<EnumReutersValutesXMLResult><EnumRValutes><num_code>8</num_code><char_code>ALL</char_code><Title_ru>Албанский лек</Title_ru><Title_en>Albanian Lek</Title_en></EnumRValutes><EnumRValutes><num_code>12</num_code><char_code>DZD</char_code><Title_ru>Алжирский динар</Title_ru><Title_en>Algerian Dinar</Title_en></EnumRValutes></EnumReutersValutesXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.EnumReutersValutesXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:EnumReutersValutesXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesEnumValutesXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "EnumValutesXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 1)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "EnumValutesXML",
		Datastructure: datastructures.EnumValutesXML{
			Seld:  false,
			XMLNs: "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<EnumValutesXML xmlns="http://web.cbr.ru/"><Seld>false</Seld></EnumValutesXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.EnumValutesXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:EnumValutesXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase

	testEnumValutesXMLResult := datastructures.EnumValutesXMLResult{
		EnumValutes: make([]datastructures.EnumValutesXMLResultElem, 2),
	}
	testEnumValutesXMLElem := datastructures.EnumValutesXMLResultElem{
		Vcode:       "R01010",
		Vname:       "Австралийский доллар",
		VEngname:    "Australian Dollar",
		Vnom:        1,
		VcommonCode: "R01010",
		VnumCode:    36,
		VcharCode:   "AUD",
	}
	testEnumValutesXMLResult.EnumValutes[0] = testEnumValutesXMLElem
	testEnumValutesXMLElem = datastructures.EnumValutesXMLResultElem{
		Vcode:       "R01015",
		Vname:       "Австрийский шиллинг",
		VEngname:    "Austrian Shilling",
		Vnom:        1000,
		VcommonCode: "R01015",
		VnumCode:    40,
		VcharCode:   "ATS",
	}
	testEnumValutesXMLResult.EnumValutes[1] = testEnumValutesXMLElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "EnumValutesXML",
		Datastructure:     testEnumValutesXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<EnumValutesXMLResult><EnumValutes><Vcode>R01010</Vcode><Vname>Австралийский доллар</Vname><VEngname>Australian Dollar</VEngname><Vnom>1</Vnom><VcommonCode>R01010</VcommonCode><VnumCode>36</VnumCode><VcharCode>AUD</VcharCode></EnumValutes><EnumValutes><Vcode>R01015</Vcode><Vname>Австрийский шиллинг</Vname><VEngname>Austrian Shilling</VEngname><Vnom>1000</Vnom><VcommonCode>R01015</VcommonCode><VnumCode>40</VnumCode><VcharCode>ATS</VcharCode></EnumValutes></EnumValutesXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.EnumValutesXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:EnumValutesXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesKeyRateXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "KeyRateXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "KeyRateXML",
		Datastructure: datastructures.KeyRateXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<KeyRateXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></KeyRateXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.KeyRateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:KeyRateXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "KeyRateXML",
		Datastructure: datastructures.KeyRateXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.KeyRateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:KeyRateXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "KeyRateXML",
		Datastructure: datastructures.KeyRateXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.KeyRateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:KeyRateXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "KeyRateXML",
		Datastructure: datastructures.KeyRateXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.KeyRateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:KeyRateXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testKeyRateXMLResult := datastructures.KeyRateXMLResult{
		KR: make([]datastructures.KeyRateXMLResultElem, 2),
	}
	testKeyRateXMLResultElem := datastructures.KeyRateXMLResultElem{
		DT:   time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Rate: "7.50",
	}
	testKeyRateXMLResult.KR[0] = testKeyRateXMLResultElem
	testKeyRateXMLResultElem = datastructures.KeyRateXMLResultElem{
		DT:   time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Rate: "7.50",
	}
	testKeyRateXMLResult.KR[1] = testKeyRateXMLResultElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "KeyRateXML",
		Datastructure:     testKeyRateXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<KeyRateXMLResult><KR><DT>2023-06-22T00:00:00Z</DT><Rate>7.50</Rate></KR><KR><DT>2023-06-23T00:00:00Z</DT><Rate>7.50</Rate></KR></KeyRateXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.KeyRateXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:KeyRateXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesMainInfoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "MainInfoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 0)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	testMainInfoXMLResult := datastructures.MainInfoXMLResult{
		KeyRate: datastructures.KeyRateElem{
			Title:   "Ключевая ставка",
			Date:    "24.07.2023",
			KeyRate: "8.50",
		},
		Inflation: datastructures.InflationElem{
			Title:     "Инфляция",
			Date:      "01.06.2023",
			Inflation: "3.25",
		},
		Stavka_ref: datastructures.Stavka_refElem{
			Title:      "Ставка рефинансирования",
			Date:       "24.07.2023",
			Stavka_ref: "8.50",
		},
		GoldBaks: datastructures.GoldBaksElem{
			Title:    "Международные резервы",
			Date:     "28.07.2023",
			GoldBaks: "594",
		},
	}

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "MainInfoXML",
		Datastructure:     testMainInfoXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<MainInfoXMLResult><keyRate Title="Ключевая ставка" Date="24.07.2023">8.50</keyRate><Inflation Title="Инфляция" Date="01.06.2023">3.25</Inflation><stavka_ref Title="Ставка рефинансирования" Date="24.07.2023">8.50</stavka_ref><GoldBaks Title="Международные резервы" Date="28.07.2023">594</GoldBaks></MainInfoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MainInfoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MainInfoXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesMrrf7DXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "mrrf7DXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "Mrrf7DXML",
		Datastructure: datastructures.Mrrf7DXML{
			FromDate: "2023-06-15",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<mrrf7DXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-15</fromDate><ToDate>2023-06-23</ToDate></mrrf7DXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Mrrf7DXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Mrrf7DXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "Mrrf7DXML",
		Datastructure: datastructures.Mrrf7DXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Mrrf7DXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Mrrf7DXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "Mrrf7DXML",
		Datastructure: datastructures.Mrrf7DXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Mrrf7DXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Mrrf7DXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "Mrrf7DXML",
		Datastructure: datastructures.Mrrf7DXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Mrrf7DXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Mrrf7DXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testMrrf7DXMLResult := datastructures.Mrrf7DXMLResult{
		Mr: make([]datastructures.Mrrf7DXMLResultElem, 2),
	}
	testMrrf7DXMLResultElem := datastructures.Mrrf7DXMLResultElem{
		D0:  time.Date(2023, time.June, 16, 0, 0, 0, 0, time.UTC),
		Val: "587.50",
	}
	testMrrf7DXMLResult.Mr[0] = testMrrf7DXMLResultElem
	testMrrf7DXMLResultElem = datastructures.Mrrf7DXMLResultElem{
		D0:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Val: "586.90",
	}
	testMrrf7DXMLResult.Mr[1] = testMrrf7DXMLResultElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "Mrrf7DXML",
		Datastructure:     testMrrf7DXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<Mrrf7DXMLResult><mr><D0>2023-06-16T00:00:00Z</D0><val>587.50</val></mr><mr><D0>2023-06-23T00:00:00Z</D0><val>586.90</val></mr></Mrrf7DXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Mrrf7DXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Mrrf7DXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesMrrfXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "mrrfXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "MrrfXML",
		Datastructure: datastructures.MrrfXML{
			FromDate: "2023-05-01",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<mrrfXML xmlns="http://web.cbr.ru/"><fromDate>2023-05-01</fromDate><ToDate>2023-06-23</ToDate></mrrfXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MrrfXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MrrfXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "MrrfXML",
		Datastructure: datastructures.MrrfXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MrrfXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MrrfXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "MrrfXML",
		Datastructure: datastructures.MrrfXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MrrfXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MrrfXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "MrrfXML",
		Datastructure: datastructures.MrrfXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MrrfXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MrrfXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testMrrfXMLResult := datastructures.MrrfXMLResult{
		Mr: make([]datastructures.MrrfXMLResultElem, 2),
	}
	testMrrfXMLResultElem := datastructures.MrrfXMLResultElem{
		D0: time.Date(2023, time.May, 0o1, 0, 0, 0, 0, time.UTC),
		P1: "595787.00",
		P2: "447187.00",
		P3: "418628.00",
		P4: "23559.00",
		P5: "5000.00",
		P6: "148599.00",
	}
	testMrrfXMLResult.Mr[0] = testMrrfXMLResultElem
	testMrrfXMLResultElem = datastructures.MrrfXMLResultElem{
		D0: time.Date(2023, time.June, 0o1, 0, 0, 0, 0, time.UTC),
		P1: "584175.00",
		P2: "438344.00",
		P3: "410313.00",
		P4: "23127.00",
		P5: "4903.00",
		P6: "145832.00",
	}
	testMrrfXMLResult.Mr[1] = testMrrfXMLResultElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "MrrfXML",
		Datastructure:     testMrrfXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<MrrfXMLResult><mr><D0>2023-05-01T00:00:00Z</D0><p1>595787.00</p1><p2>447187.00</p2><p3>418628.00</p3><p4>23559.00</p4><p5>5000.00</p5><p6>148599.00</p6></mr><mr><D0>2023-06-01T00:00:00Z</D0><p1>584175.00</p1><p2>438344.00</p2><p3>410313.00</p3><p4>23127.00</p4><p5>4903.00</p5><p6>145832.00</p6></mr></MrrfXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.MrrfXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:MrrfXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesNewsInfoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "NewsInfoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "NewsInfoXML",
		Datastructure: datastructures.NewsInfoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<NewsInfoXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></NewsInfoXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.NewsInfoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:NewsInfoXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "NewsInfoXML",
		Datastructure: datastructures.NewsInfoXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.NewsInfoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:NewsInfoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "NewsInfoXML",
		Datastructure: datastructures.NewsInfoXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.NewsInfoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:NewsInfoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "NewsInfoXML",
		Datastructure: datastructures.NewsInfoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.NewsInfoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:NewsInfoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testNewsInfoXMLResult := datastructures.NewsInfoXMLResult{
		News: make([]datastructures.NewsInfoXMLResultElem, 2),
	}
	testNewsInfoXMLResultElem := datastructures.NewsInfoXMLResultElem{
		Doc_id:  35498,
		DocDate: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Title:   "О развитии банковского сектора Российской Федерации в мае 2023 года",
		Url:     "/analytics/bank_sector/develop/#a_48876",
	}
	testNewsInfoXMLResult.News[0] = testNewsInfoXMLResultElem
	testNewsInfoXMLResultElem = datastructures.NewsInfoXMLResultElem{
		Doc_id:  35495,
		DocDate: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Title:   "Указание Банка России от 10.01.2023 № 6356-У",
		Url:     "/Queries/UniDbQuery/File/90134/2803",
	}
	testNewsInfoXMLResult.News[1] = testNewsInfoXMLResultElem

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "NewsInfoXML",
		Datastructure:     testNewsInfoXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<NewsInfoXMLResult><News><Doc_id>35498</Doc_id><DocDate>2023-06-22T00:00:00Z</DocDate><Title>О развитии банковского сектора Российской Федерации в мае 2023 года</Title><Url>/analytics/bank_sector/develop/#a_48876</Url></News><News><Doc_id>35495</Doc_id><DocDate>2023-06-22T00:00:00Z</DocDate><Title>Указание Банка России от 10.01.2023 № 6356-У</Title><Url>/Queries/UniDbQuery/File/90134/2803</Url></News></NewsInfoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.NewsInfoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:NewsInfoXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesOmodInfoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "OmodInfoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 0)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	testOmodInfoXML := datastructures.OmodInfoXMLResult{
		Date: "05.03.2018",
		DirectRepo: datastructures.DirectRepoElem{
			Time:      "10:00",
			Debt:      "0",
			Rate:      "0",
			Minrate1D: "7.5",
			Minrate7D: "7.5",
		},
		RevRepo: datastructures.RevRepoElem{
			Time:     "10:00",
			Debt:     "0",
			Rate:     "4.97",
			Sum_debt: "0",
		},
		OBR: datastructures.OBRElem{
			Time: "10:00",
			Debt: "0",
			Rate: "3.55",
		},
		Deposit:         "0",
		Credit:          "0",
		VolNom:          "6741.11",
		TotalFixRepoVol: "3132.2",
		FixRepoDate:     "02.03.2018",
		FixRepo1D: datastructures.FixRepo1DElem{
			Debt: "3130.1",
			Rate: "8.5",
		},
		FixRepo7D: datastructures.FixRepo7DElem{
			Debt: "0",
			Rate: "8.5",
		},
		FixRepo1Y: datastructures.FixRepo1YElem{
			Rate: "8.5",
		},
	}

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "OmodInfoXML",
		Datastructure:     testOmodInfoXML,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OmodInfoXMLResult Date="05.03.2018"><DirectRepo Time="10:00"><debt>0</debt><rate>0</rate><minrate1D>7.5</minrate1D><minrate7D>7.5</minrate7D></DirectRepo><RevRepo Time="10:00"><debt>0</debt><rate>4.97</rate><sum_debt>0</sum_debt></RevRepo><OBR Time="10:00"><debt>0</debt><rate>3.55</rate></OBR><Deposit>0</Deposit><Credit>0</Credit><VolNom>6741.11</VolNom><TotalFixRepoVol>3132.2</TotalFixRepoVol><FixRepoDate>02.03.2018</FixRepoDate><FixRepo1D><debt>3130.1</debt><rate>8.5</rate></FixRepo1D><FixRepo7D><debt>0</debt><rate>8.5</rate></FixRepo7D><FixRepo1Y><rate>8.5</rate></FixRepo1Y></OmodInfoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OmodInfoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OmodInfoXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesOstatDepoNewXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "OstatDepoNewXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "OstatDepoNewXML",
		Datastructure: datastructures.OstatDepoNewXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDepoNewXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></OstatDepoNewXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoNewXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoNewXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "OstatDepoNewXML",
		Datastructure: datastructures.OstatDepoNewXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoNewXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoNewXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "OstatDepoNewXML",
		Datastructure: datastructures.OstatDepoNewXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoNewXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoNewXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "OstatDepoNewXML",
		Datastructure: datastructures.OstatDepoNewXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoNewXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoNewXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testOstatDepoNewXMLResult := datastructures.OstatDepoNewXMLResult{
		Odn: make([]datastructures.OstatDepoNewXMLResultElem, 2),
	}
	testOstatDepoNewXMLElem := datastructures.OstatDepoNewXMLResultElem{
		DT:     time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		TOTAL:  "2872966.59",
		AUC_1W: "1828340.00",
		OV_P:   "1044626.59",
	}
	testOstatDepoNewXMLResult.Odn[0] = testOstatDepoNewXMLElem
	testOstatDepoNewXMLElem = datastructures.OstatDepoNewXMLResultElem{
		DT:     time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		TOTAL:  "2890199.16",
		AUC_1W: "1828340.00",
		OV_P:   "1061859.16",
	}
	testOstatDepoNewXMLResult.Odn[1] = testOstatDepoNewXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "OstatDepoNewXML",
		Datastructure:     testOstatDepoNewXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDepoNewXMLResult><odn><DT>2023-06-22T00:00:00Z</DT><TOTAL>2872966.59</TOTAL><AUC_1W>1828340.00</AUC_1W><OV_P>1044626.59</OV_P></odn><odn><DT>2023-06-23T00:00:00Z</DT><TOTAL>2890199.16</TOTAL><AUC_1W>1828340.00</AUC_1W><OV_P>1061859.16</OV_P></odn></OstatDepoNewXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoNewXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoNewXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesOstatDepoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "OstatDepoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "OstatDepoXML",
		Datastructure: datastructures.OstatDepoXML{
			FromDate: "2022-12-29",
			ToDate:   "2022-12-30",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDepoXML xmlns="http://web.cbr.ru/"><fromDate>2022-12-29</fromDate><ToDate>2022-12-30</ToDate></OstatDepoXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "OstatDepoXML",
		Datastructure: datastructures.OstatDepoXML{
			FromDate: "022-14-22",
			ToDate:   "2022-12-30",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "OstatDepoNewXML",
		Datastructure: datastructures.OstatDepoXML{
			FromDate: "2022-12-30",
			ToDate:   "2022-12-29",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "OstatDepoXML",
		Datastructure: datastructures.OstatDepoXML{
			FromDate: "2022-12-29",
			ToDate:   "2022-12-30",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testOstatDepoXMLResult := datastructures.OstatDepoXMLResult{
		Odr: make([]datastructures.OstatDepoXMLResultElem, 2),
	}
	testOstatDepoXMLElem := datastructures.OstatDepoXMLResultElem{
		D0:    time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC),
		D1_7:  "1747362.67",
		D8_30: "2515151.15",
		Total: "4262513.81",
	}
	testOstatDepoXMLResult.Odr[0] = testOstatDepoXMLElem
	testOstatDepoXMLElem = datastructures.OstatDepoXMLResultElem{
		D0:    time.Date(2022, time.December, 30, 0, 0, 0, 0, time.UTC),
		D1_7:  "1387715.38",
		D8_30: "2515151.15",
		Total: "3897866.53",
	}
	testOstatDepoXMLResult.Odr[1] = testOstatDepoXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "OstatDepoXML",
		Datastructure:     testOstatDepoXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDepoXMLResult><odr><D0>2022-12-29T00:00:00Z</D0><D1_7>1747362.67</D1_7><D8_30>2515151.15</D8_30><total>4262513.81</total></odr><odr><D0>2022-12-30T00:00:00Z</D0><D1_7>1387715.38</D1_7><D8_30>2515151.15</D8_30><total>3897866.53</total></odr></OstatDepoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDepoXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesOstatDynamicXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "OstatDynamicXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "OstatDynamicXML",
		Datastructure: datastructures.OstatDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDynamicXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></OstatDynamicXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDynamicXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "OstatDynamicXML",
		Datastructure: datastructures.OstatDynamicXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDepoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "OstatDynamicXML",
		Datastructure: datastructures.OstatDynamicXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "OstatDynamicXML",
		Datastructure: datastructures.OstatDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testOstatDynamicXMLResult := datastructures.OstatDynamicXMLResult{
		Ostat: make([]datastructures.OstatDynamicXMLResultElem, 2),
	}
	testOstatDynamicXMLElem := datastructures.OstatDynamicXMLResultElem{
		DateOst:  time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		InRuss:   "3756300.00",
		InMoscow: "3528600.00",
	}
	testOstatDynamicXMLResult.Ostat[0] = testOstatDynamicXMLElem
	testOstatDynamicXMLElem = datastructures.OstatDynamicXMLResultElem{
		DateOst:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		InRuss:   "3688300.00",
		InMoscow: "3441000.00",
	}
	testOstatDynamicXMLResult.Ostat[1] = testOstatDynamicXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "OstatDepoXML",
		Datastructure:     testOstatDynamicXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OstatDynamicXMLResult><Ostat><DateOst>2023-06-22T00:00:00Z</DateOst><InRuss>3756300.00</InRuss><InMoscow>3528600.00</InMoscow></Ostat><Ostat><DateOst>2023-06-23T00:00:00Z</DateOst><InRuss>3688300.00</InRuss><InMoscow>3441000.00</InMoscow></Ostat></OstatDynamicXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OstatDynamicXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OstatDynamicXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesOvernightXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "OvernightXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "OvernightXML",
		Datastructure: datastructures.OvernightXML{
			FromDate: "2023-07-22",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OvernightXML xmlns="http://web.cbr.ru/"><fromDate>2023-07-22</fromDate><ToDate>2023-08-16</ToDate></OvernightXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OvernightXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OvernightXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "OvernightXML",
		Datastructure: datastructures.OvernightXML{
			FromDate: "022-14-23",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OvernightXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OvernightXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "OvernightXML",
		Datastructure: datastructures.OvernightXML{
			FromDate: "2023-08-16",
			ToDate:   "2023-07-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OvernightXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OvernightXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "OvernightXML",
		Datastructure: datastructures.OvernightXML{
			FromDate: "2023-07-22",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OvernightXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OvernightXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testOvernightXMLResult := datastructures.OvernightXMLResult{
		OB: make([]datastructures.OvernightXMLResultElem, 2),
	}
	testOvernightXMLElem := datastructures.OvernightXMLResultElem{
		Date:   time.Date(2023, time.July, 24, 0, 0, 0, 0, time.UTC),
		Stavka: "9.50",
	}
	testOvernightXMLResult.OB[0] = testOvernightXMLElem
	testOvernightXMLElem = datastructures.OvernightXMLResultElem{
		Date:   time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC),
		Stavka: "13.00",
	}
	testOvernightXMLResult.OB[1] = testOvernightXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "OvernightXML",
		Datastructure:     testOvernightXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<OvernightXMLResult><OB><date>2023-07-24T00:00:00Z</date><stavka>9.50</stavka></OB><OB><date>2023-08-15T00:00:00Z</date><stavka>13.00</stavka></OB></OvernightXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.OvernightXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:OvernightXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesRepoDebtUSDXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "RepoDebtUSDXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "RepoDebtUSDXML",
		Datastructure: datastructures.RepoDebtUSDXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RepoDebtUSDXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></RepoDebtUSDXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RepoDebtUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RepoDebtUSDXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "RepoDebtUSDXML",
		Datastructure: datastructures.RepoDebtUSDXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RepoDebtUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RepoDebtUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "RepoDebtUSDXML",
		Datastructure: datastructures.RepoDebtUSDXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RepoDebtUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RepoDebtUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "RepoDebtUSDXML",
		Datastructure: datastructures.RepoDebtUSDXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RepoDebtUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RepoDebtUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testRepoDebtUSDXMLResult := datastructures.RepoDebtUSDXMLResult{
		Rd: make([]datastructures.RepoDebtUSDXMLResultElem, 4),
	}
	testRepoDebtUSDXMLElem := datastructures.RepoDebtUSDXMLResultElem{
		D0: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		TP: 0,
	}
	testRepoDebtUSDXMLResult.Rd[0] = testRepoDebtUSDXMLElem
	testRepoDebtUSDXMLElem = datastructures.RepoDebtUSDXMLResultElem{
		D0: time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		TP: 1,
	}
	testRepoDebtUSDXMLResult.Rd[1] = testRepoDebtUSDXMLElem
	testRepoDebtUSDXMLElem = datastructures.RepoDebtUSDXMLResultElem{
		D0: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		TP: 0,
	}
	testRepoDebtUSDXMLResult.Rd[2] = testRepoDebtUSDXMLElem
	testRepoDebtUSDXMLElem = datastructures.RepoDebtUSDXMLResultElem{
		D0: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		TP: 1,
	}
	testRepoDebtUSDXMLResult.Rd[3] = testRepoDebtUSDXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "RepoDebtUSDXML",
		Datastructure:     testRepoDebtUSDXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RepoDebtUSDXMLResult><rd><D0>2023-06-22T00:00:00Z</D0><TP>0</TP></rd><rd><D0>2023-06-22T00:00:00Z</D0><TP>1</TP></rd><rd><D0>2023-06-23T00:00:00Z</D0><TP>0</TP></rd><rd><D0>2023-06-23T00:00:00Z</D0><TP>1</TP></rd></RepoDebtUSDXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RepoDebtUSDXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RepoDebtUSDXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesRepoDebtXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "Repo_debtXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "Repo_debtXML",
		Datastructure: datastructures.Repo_debtXML{
			FromDate: "2023-07-22",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<Repo_debtXML xmlns="http://web.cbr.ru/"><fromDate>2023-07-22</fromDate><ToDate>2023-08-16</ToDate></Repo_debtXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Repo_debtXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Repo_debtXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "Repo_debtXML",
		Datastructure: datastructures.Repo_debtXML{
			FromDate: "022-14-23",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Repo_debtXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Repo_debtXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "Repo_debtXML",
		Datastructure: datastructures.Repo_debtXML{
			FromDate: "2023-08-16",
			ToDate:   "2023-07-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Repo_debtXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Repo_debtXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "Repo_debtXML",
		Datastructure: datastructures.Repo_debtXML{
			FromDate: "2023-07-22",
			ToDate:   "2023-08-16",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Repo_debtXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Repo_debtXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testRepo_debtXMLResult := datastructures.Repo_debtXMLResult{ //nolint:revive, stylecheck, nolintlint
		RD: make([]datastructures.Repo_debtXMLResultElem, 2),
	}
	testRepo_debtXMLElem := datastructures.Repo_debtXMLResultElem{ //nolint:revive, stylecheck, nolintlint
		Date:     time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Debt:     "1378387.6",
		Debt_auc: "1378387.6",
		Debt_fix: "0.0",
	}
	testRepo_debtXMLResult.RD[0] = testRepo_debtXMLElem
	testRepo_debtXMLElem = datastructures.Repo_debtXMLResultElem{
		Date:     time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Debt:     "1378379.7",
		Debt_auc: "1378379.7",
		Debt_fix: "0.0",
	}
	testRepo_debtXMLResult.RD[1] = testRepo_debtXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "Repo_debtXML",
		Datastructure:     testRepo_debtXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<Repo_debtXMLResult><RD><Date>2023-06-22T00:00:00Z</Date><debt>1378387.6</debt><debt_auc>1378387.6</debt_auc><debt_fix>0.0</debt_fix></RD><RD><Date>2023-06-23T00:00:00Z</Date><debt>1378379.7</debt><debt_auc>1378379.7</debt_auc><debt_fix>0.0</debt_fix></RD></Repo_debtXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.Repo_debtXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:Repo_debtXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesROISfixXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "ROISfixXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "ROISfixXML",
		Datastructure: datastructures.ROISfixXML{
			FromDate: "2022-02-27",
			ToDate:   "2022-03-02",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<ROISfixXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-27</fromDate><ToDate>2022-03-02</ToDate></ROISfixXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.ROISfixXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:ROISfixXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "ROISfixXML",
		Datastructure: datastructures.ROISfixXML{
			FromDate: "022-14-23",
			ToDate:   "2022-03-02",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.ROISfixXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:ROISfixXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "ROISfixXML",
		Datastructure: datastructures.ROISfixXML{
			FromDate: "2022-03-02",
			ToDate:   "2022-02-27",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.ROISfixXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:ROISfixXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "ROISfixXML",
		Datastructure: datastructures.ROISfixXML{
			FromDate: "2022-02-27",
			ToDate:   "2022-03-02",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.ROISfixXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:ROISfixXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testROISfixXMLResult := datastructures.ROISfixXMLResult{
		Rf: make([]datastructures.ROISfixXMLResultElem, 2),
	}
	testROISfixXMLElem := datastructures.ROISfixXMLResultElem{
		D0:  time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		R1W: "17.83",
		R2W: "18.00",
		R1M: "20.65",
		R2M: "21.96",
		R3M: "23.23",
		R6M: "24.52",
	}
	testROISfixXMLResult.Rf[0] = testROISfixXMLElem
	testROISfixXMLElem = datastructures.ROISfixXMLResultElem{
		D0:  time.Date(2022, time.March, 0o1, 0, 0, 0, 0, time.UTC),
		R1W: "19.85",
		R2W: "19.91",
		R1M: "22.63",
		R2M: "23.79",
		R3M: "24.49",
		R6M: "25.71",
	}
	testROISfixXMLResult.Rf[1] = testROISfixXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "ROISfixXML",
		Datastructure:     testROISfixXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<ROISfixXMLResult><rf><D0>2022-02-28T00:00:00Z</D0><R1W>17.83</R1W><R2W>18.00</R2W><R1M>20.65</R1M><R2M>21.96</R2M><R3M>23.23</R3M><R6M>24.52</R6M></rf><rf><D0>2022-03-01T00:00:00Z</D0><R1W>19.85</R1W><R2W>19.91</R2W><R1M>22.63</R1M><R2M>23.79</R2M><R3M>24.49</R3M><R6M>25.71</R6M></rf></ROISfixXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.ROISfixXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:ROISfixXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesRuoniaSVXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "RuoniaSVXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "RuoniaSVXML",
		Datastructure: datastructures.RuoniaSVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RuoniaSVXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></RuoniaSVXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaSVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaSVXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "RuoniaSVXML",
		Datastructure: datastructures.RuoniaSVXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaSVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaSVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "RuoniaSVXML",
		Datastructure: datastructures.RuoniaSVXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaSVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaSVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "RuoniaSVXML",
		Datastructure: datastructures.RuoniaSVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaSVXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaSVXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testRuoniaSVXMLResult := datastructures.RuoniaSVXMLResult{
		Ra: make([]datastructures.RuoniaSVXMLResultElem, 2),
	}
	testRuoniaSVXMLElem := datastructures.RuoniaSVXMLResultElem{
		DT:            time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		RUONIA_Index:  "2.65003371140540",
		RUONIA_AVG_1M: "7.33031817626889",
		RUONIA_AVG_3M: "7.28023580262342",
		RUONIA_AVG_6M: "7.34479164787354",
	}
	testRuoniaSVXMLResult.Ra[0] = testRuoniaSVXMLElem
	testRuoniaSVXMLElem = datastructures.RuoniaSVXMLResultElem{
		DT:            time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		RUONIA_Index:  "2.65055282759819",
		RUONIA_AVG_1M: "7.32512579295002",
		RUONIA_AVG_3M: "7.27890778428907",
		RUONIA_AVG_6M: "7.34359578515310",
	}
	testRuoniaSVXMLResult.Ra[1] = testRuoniaSVXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "RuoniaSVXMLResult",
		Datastructure:     testRuoniaSVXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RuoniaSVXMLResult><ra><DT>2023-06-22T00:00:00Z</DT><RUONIA_Index>2.65003371140540</RUONIA_Index><RUONIA_AVG_1M>7.33031817626889</RUONIA_AVG_1M><RUONIA_AVG_3M>7.28023580262342</RUONIA_AVG_3M><RUONIA_AVG_6M>7.34479164787354</RUONIA_AVG_6M></ra><ra><DT>2023-06-23T00:00:00Z</DT><RUONIA_Index>2.65055282759819</RUONIA_Index><RUONIA_AVG_1M>7.32512579295002</RUONIA_AVG_1M><RUONIA_AVG_3M>7.27890778428907</RUONIA_AVG_3M><RUONIA_AVG_6M>7.34359578515310</RUONIA_AVG_6M></ra></RuoniaSVXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaSVXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaSVXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesRuoniaXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "RuoniaXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "RuoniaXML",
		Datastructure: datastructures.RuoniaXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RuoniaXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></RuoniaXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "RuoniaXML",
		Datastructure: datastructures.RuoniaXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "RuoniaXML",
		Datastructure: datastructures.RuoniaXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "RuoniaXML",
		Datastructure: datastructures.RuoniaXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testRuoniaXMLResult := datastructures.RuoniaXMLResult{
		Ro: make([]datastructures.RuoniaXMLResultElem, 2),
	}
	testRuoniaXMLElem := datastructures.RuoniaXMLResultElem{
		D0:         time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Ruo:        "7.1500",
		Vol:        "367.9500",
		DateUpdate: time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
	}
	testRuoniaXMLResult.Ro[0] = testRuoniaXMLElem
	testRuoniaXMLElem = datastructures.RuoniaXMLResultElem{
		D0:         time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Ruo:        "7.1300",
		Vol:        "388.4500",
		DateUpdate: time.Date(2023, time.June, 26, 0, 0, 0, 0, time.UTC),
	}
	testRuoniaXMLResult.Ro[1] = testRuoniaXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "RuoniaXMLResult",
		Datastructure:     testRuoniaXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<RuoniaXMLResult><ro><D0>2023-06-22T00:00:00Z</D0><ruo>7.1500</ruo><vol>367.9500</vol><DateUpdate>2023-06-23T00:00:00Z</DateUpdate></ro><ro><D0>2023-06-23T00:00:00Z</D0><ruo>7.1300</ruo><vol>388.4500</vol><DateUpdate>2023-06-26T00:00:00Z</DateUpdate></ro></RuoniaXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.RuoniaXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:RuoniaXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSaldoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SaldoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SaldoXML",
		Datastructure: datastructures.SaldoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SaldoXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-22</fromDate><ToDate>2023-06-23</ToDate></SaldoXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SaldoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SaldoXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SaldoXML",
		Datastructure: datastructures.SaldoXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SaldoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SaldoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SaldoXML",
		Datastructure: datastructures.SaldoXML{
			FromDate: "2023-06-23",
			ToDate:   "2023-06-22",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SaldoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SaldoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SaldoXML",
		Datastructure: datastructures.SaldoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SaldoXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SaldoXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSaldoXMLResult := datastructures.SaldoXMLResult{
		So: make([]datastructures.SaldoXMLResultElem, 2),
	}
	testSaldoXMLElem := datastructures.SaldoXMLResultElem{
		Dt:         time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		DEADLINEBS: "1044.60",
	}
	testSaldoXMLResult.So[0] = testSaldoXMLElem
	testSaldoXMLElem = datastructures.SaldoXMLResultElem{
		Dt:         time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		DEADLINEBS: "1061.30",
	}
	testSaldoXMLResult.So[1] = testSaldoXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SaldoXMLResult",
		Datastructure:     testSaldoXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SaldoXMLResult><So><Dt>2023-06-22T00:00:00Z</Dt><DEADLINEBS>1044.60</DEADLINEBS></So><So><Dt>2023-06-23T00:00:00Z</Dt><DEADLINEBS>1061.30</DEADLINEBS></So></SaldoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SaldoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SaldoXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapDayTotalXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapDayTotalXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapDayTotalXML",
		Datastructure: datastructures.SwapDayTotalXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapDayTotalXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-25</fromDate><ToDate>2022-02-28</ToDate></SwapDayTotalXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDayTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDayTotalXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapDayTotalXML",
		Datastructure: datastructures.SwapDayTotalXML{
			FromDate: "022-14-23",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDayTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDayTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapDayTotalXML",
		Datastructure: datastructures.SwapDayTotalXML{
			FromDate: "2022-02-28",
			ToDate:   "2022-02-25",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDayTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDayTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapDayTotalXML",
		Datastructure: datastructures.SwapDayTotalXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDayTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDayTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapDayTotalXMLResult := datastructures.SwapDayTotalXMLResult{
		SDT: make([]datastructures.SwapDayTotalXMLResultElem, 2),
	}
	testSwapDayTotalXMLElem := datastructures.SwapDayTotalXMLResultElem{
		DT:   time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		Swap: "0.0",
	}
	testSwapDayTotalXMLResult.SDT[0] = testSwapDayTotalXMLElem
	testSwapDayTotalXMLElem = datastructures.SwapDayTotalXMLResultElem{
		DT:   time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		Swap: "24120.4",
	}
	testSwapDayTotalXMLResult.SDT[1] = testSwapDayTotalXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapDayTotalXML",
		Datastructure:     testSwapDayTotalXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapDayTotalXMLResult><SDT><DT>2022-02-28T00:00:00Z</DT><Swap>0.0</Swap></SDT><SDT><DT>2022-02-25T00:00:00Z</DT><Swap>24120.4</Swap></SDT></SwapDayTotalXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDayTotalXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDayTotalXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapDynamicXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapDynamicXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapDynamicXML",
		Datastructure: datastructures.SwapDynamicXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapDynamicXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-25</fromDate><ToDate>2022-02-28</ToDate></SwapDynamicXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDynamicXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapDynamicXML",
		Datastructure: datastructures.SwapDynamicXML{
			FromDate: "022-14-23",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapDynamicXML",
		Datastructure: datastructures.SwapDynamicXML{
			FromDate: "2022-02-28",
			ToDate:   "2022-02-25",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapDynamicXML",
		Datastructure: datastructures.SwapDynamicXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDynamicXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDynamicXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapDynamicXMLResult := datastructures.SwapDynamicXMLResult{
		Swap: make([]datastructures.SwapDynamicXMLResultElem, 2),
	}
	testSwapDynamicXMLElem := datastructures.SwapDynamicXMLResultElem{
		DateBuy:  time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		BaseRate: "96.8252",
		SD:       "0.0882",
		TIR:      "10.5000",
		Stavka:   "-0.576000",
		Currency: 1,
	}
	testSwapDynamicXMLResult.Swap[0] = testSwapDynamicXMLElem
	testSwapDynamicXMLElem = datastructures.SwapDynamicXMLResultElem{
		DateBuy:  time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		BaseRate: "87.1154",
		SD:       "0.0748",
		TIR:      "10.5000",
		Stavka:   "0.050000",
		Currency: 0,
	}
	testSwapDynamicXMLResult.Swap[1] = testSwapDynamicXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapDynamicXML",
		Datastructure:     testSwapDynamicXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapDynamicXMLResult><Swap><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><BaseRate>96.8252</BaseRate><SD>0.0882</SD><TIR>10.5000</TIR><Stavka>-0.576000</Stavka><Currency>1</Currency></Swap><Swap><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><BaseRate>87.1154</BaseRate><SD>0.0748</SD><TIR>10.5000</TIR><Stavka>0.050000</Stavka><Currency>0</Currency></Swap></SwapDynamicXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapDynamicXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapDynamicXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapInfoSellUSDVolXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapInfoSellUSDVolXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapInfoSellUSDVolXML",
		Datastructure: datastructures.SwapInfoSellUSDVolXML{
			FromDate: "2022-02-24",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellUSDVolXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-24</fromDate><ToDate>2022-02-28</ToDate></SwapInfoSellUSDVolXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDVolXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapInfoSellUSDVolXML",
		Datastructure: datastructures.SwapInfoSellUSDVolXML{
			FromDate: "022-14-23",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapInfoSellUSDVolXML",
		Datastructure: datastructures.SwapInfoSellUSDVolXML{
			FromDate: "2022-02-28",
			ToDate:   "2022-02-24",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapInfoSellUSDVolXML",
		Datastructure: datastructures.SwapInfoSellUSDVolXML{
			FromDate: "2022-02-24",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapInfoSellUSDVolXMLResult := datastructures.SwapInfoSellUSDVolXMLResult{
		SSUV: make([]datastructures.SwapInfoSellUSDVolXMLResultElem, 2),
	}
	testSwapInfoSellUSDVolXMLElem := datastructures.SwapInfoSellUSDVolXMLResultElem{
		DT:           time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		TODTOMrubvol: "435577.0",
		TODTOMusdvol: "5000.0",
		TOMSPTrubvol: "128974.3",
		TOMSPTusdvol: "1480.5",
	}
	testSwapInfoSellUSDVolXMLResult.SSUV[0] = testSwapInfoSellUSDVolXMLElem
	testSwapInfoSellUSDVolXMLElem = datastructures.SwapInfoSellUSDVolXMLResultElem{
		DT:           time.Date(2022, time.February, 24, 0, 0, 0, 0, time.UTC),
		TODTOMrubvol: "403236.5",
		TODTOMusdvol: "5000.0",
		TOMSPTrubvol: "32299.2",
		TOMSPTusdvol: "400.5",
	}
	testSwapInfoSellUSDVolXMLResult.SSUV[1] = testSwapInfoSellUSDVolXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapInfoSellUSDVolXML",
		Datastructure:     testSwapInfoSellUSDVolXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellUSDVolXMLResult><SSUV><DT>2022-02-25T00:00:00Z</DT><TODTOMrubvol>435577.0</TODTOMrubvol><TODTOMusdvol>5000.0</TODTOMusdvol><TOMSPTrubvol>128974.3</TOMSPTrubvol><TOMSPTusdvol>1480.5</TOMSPTusdvol></SSUV><SSUV><DT>2022-02-24T00:00:00Z</DT><TODTOMrubvol>403236.5</TODTOMrubvol><TODTOMusdvol>5000.0</TODTOMusdvol><TOMSPTrubvol>32299.2</TOMSPTrubvol><TOMSPTusdvol>400.5</TOMSPTusdvol></SSUV></SwapInfoSellUSDVolXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDVolXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDVolXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapInfoSellUSDXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapInfoSellUSDXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure: datastructures.SwapInfoSellUSDXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellUSDXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-25</fromDate><ToDate>2022-02-28</ToDate></SwapInfoSellUSDXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure: datastructures.SwapInfoSellUSDXML{
			FromDate: "022-14-23",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure: datastructures.SwapInfoSellUSDXML{
			FromDate: "2022-02-28",
			ToDate:   "2022-02-25",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure: datastructures.SwapInfoSellUSDXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapInfoSellUSDXMLResult := datastructures.SwapInfoSellUSDXMLResult{
		SSU: make([]datastructures.SwapInfoSellUSDXMLResultElem, 2),
	}
	testSwapInfoSellUSDXMLElem := datastructures.SwapInfoSellUSDXMLResultElem{
		DateBuy:  time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		DateSPOT: time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
		BaseRate: "87.115400",
		SD:       "0.016500",
		TIR:      "8.5000",
		Stavka:   "1.5500",
		Limit:    "2.0000",
		Type:     1,
	}
	testSwapInfoSellUSDXMLResult.SSU[0] = testSwapInfoSellUSDXMLElem
	testSwapInfoSellUSDXMLElem = datastructures.SwapInfoSellUSDXMLResultElem{
		DateBuy:  time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC),
		DateSPOT: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
		BaseRate: "87.115400",
		SD:       "0.049600",
		TIR:      "8.5000",
		Stavka:   "1.5500",
		Limit:    "5.0000",
		Type:     0,
	}
	testSwapInfoSellUSDXMLResult.SSU[1] = testSwapInfoSellUSDXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure:     testSwapInfoSellUSDXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellUSDXMLResult><SSU><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-28T00:00:00Z</DateSell><DateSPOT>2022-03-01T00:00:00Z</DateSPOT><Type>1</Type><BaseRate>87.115400</BaseRate><SD>0.016500</SD><TIR>8.5000</TIR><Stavka>1.5500</Stavka><limit>2.0000</limit></SSU><SSU><DateBuy>2022-02-25T00:00:00Z</DateBuy><DateSell>2022-02-25T00:00:00Z</DateSell><DateSPOT>2022-02-28T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>87.115400</BaseRate><SD>0.049600</SD><TIR>8.5000</TIR><Stavka>1.5500</Stavka><limit>5.0000</limit></SSU></SwapInfoSellUSDXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellUSDXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellUSDXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapInfoSellVolXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapInfoSellVolXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapInfoSellVolXML",
		Datastructure: datastructures.SwapInfoSellVolXML{
			FromDate: "2023-05-05",
			ToDate:   "2023-05-10",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellVolXML xmlns="http://web.cbr.ru/"><fromDate>2023-05-05</fromDate><ToDate>2023-05-10</ToDate></SwapInfoSellVolXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellVolXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapInfoSellVolXML",
		Datastructure: datastructures.SwapInfoSellVolXML{
			FromDate: "022-14-23",
			ToDate:   "2023-05-10",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapInfoSellVolXML",
		Datastructure: datastructures.SwapInfoSellVolXML{
			FromDate: "2023-05-10",
			ToDate:   "2023-05-05",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapInfoSellUSDXML",
		Datastructure: datastructures.SwapInfoSellVolXML{
			FromDate: "2023-05-05",
			ToDate:   "2023-05-10",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellVolXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellVolXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapInfoSellVolXMLResult := datastructures.SwapInfoSellVolXMLResult{
		SSUV: make([]datastructures.SwapInfoSellVolXMLResultElem, 2),
	}
	testSwapInfoSellVolXMLElem := datastructures.SwapInfoSellVolXMLResultElem{
		DT:       time.Date(2023, time.May, 10, 0, 0, 0, 0, time.UTC),
		Currency: 2,
		Type:     0,
		VOL_FC:   "1113.5",
		VOL_RUB:  "12512.6",
	}
	testSwapInfoSellVolXMLResult.SSUV[0] = testSwapInfoSellVolXMLElem
	testSwapInfoSellVolXMLElem = datastructures.SwapInfoSellVolXMLResultElem{
		DT:       time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC),
		Currency: 2,
		Type:     0,
		VOL_FC:   "4583.7",
		VOL_RUB:  "51606.0",
	}
	testSwapInfoSellVolXMLResult.SSUV[1] = testSwapInfoSellVolXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapInfoSellVolXML",
		Datastructure:     testSwapInfoSellVolXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellVolXMLResult><SSUV><DT>2023-05-10T00:00:00Z</DT><Currency>2</Currency><type>0</type><VOL_FC>1113.5</VOL_FC><VOL_RUB>12512.6</VOL_RUB></SSUV><SSUV><DT>2023-05-05T00:00:00Z</DT><Currency>2</Currency><type>0</type><VOL_FC>4583.7</VOL_FC><VOL_RUB>51606.0</VOL_RUB></SSUV></SwapInfoSellVolXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellVolXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellVolXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapInfoSellXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapInfoSellXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapInfoSellXML",
		Datastructure: datastructures.SwapInfoSellXML{
			FromDate: "2023-06-20",
			ToDate:   "2023-06-21",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellXML xmlns="http://web.cbr.ru/"><fromDate>2023-06-20</fromDate><ToDate>2023-06-21</ToDate></SwapInfoSellXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapInfoSellXML",
		Datastructure: datastructures.SwapInfoSellXML{
			FromDate: "022-14-23",
			ToDate:   "2023-06-21",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapInfoSellXML",
		Datastructure: datastructures.SwapInfoSellXML{
			FromDate: "2023-06-21",
			ToDate:   "2023-06-20",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapInfoSellXML",
		Datastructure: datastructures.SwapInfoSellXML{
			FromDate: "2023-06-20",
			ToDate:   "2023-06-21",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapInfoSellXMLResult := datastructures.SwapInfoSellXMLResult{
		SSU: make([]datastructures.SwapInfoSellXMLResultElem, 2),
	}
	testSwapInfoSellXMLElem := datastructures.SwapInfoSellXMLResultElem{
		Currency: 2,
		DateBuy:  time.Date(2023, time.June, 21, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2023, time.June, 21, 0, 0, 0, 0, time.UTC),
		DateSPOT: time.Date(2023, time.June, 26, 0, 0, 0, 0, time.UTC),
		Type:     0,
		BaseRate: "11.764246",
		SD:       "0.003375",
		TIR:      "6.5000",
		Stavka:   "4.3440",
		Limit:    "10.0000",
	}
	testSwapInfoSellXMLResult.SSU[0] = testSwapInfoSellXMLElem
	testSwapInfoSellXMLElem = datastructures.SwapInfoSellXMLResultElem{
		Currency: 2,
		DateBuy:  time.Date(2023, time.June, 20, 0, 0, 0, 0, time.UTC),
		DateSell: time.Date(2023, time.June, 20, 0, 0, 0, 0, time.UTC),
		DateSPOT: time.Date(2023, time.June, 21, 0, 0, 0, 0, time.UTC),
		Type:     0,
		BaseRate: "11.730496",
		SD:       "0.000626",
		TIR:      "6.5000",
		Stavka:   "4.4890",
		Limit:    "10.0000",
	}
	testSwapInfoSellXMLResult.SSU[1] = testSwapInfoSellXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapInfoSellXML",
		Datastructure:     testSwapInfoSellXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapInfoSellXMLResult><SSU><Currency>2</Currency><DateBuy>2023-06-21T00:00:00Z</DateBuy><DateSell>2023-06-21T00:00:00Z</DateSell><DateSPOT>2023-06-26T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>11.764246</BaseRate><SD>0.003375</SD><TIR>6.5000</TIR><Stavka>4.3440</Stavka><limit>10.0000</limit></SSU><SSU><Currency>2</Currency><DateBuy>2023-06-20T00:00:00Z</DateBuy><DateSell>2023-06-20T00:00:00Z</DateSell><DateSPOT>2023-06-21T00:00:00Z</DateSPOT><Type>0</Type><BaseRate>11.730496</BaseRate><SD>0.000626</SD><TIR>6.5000</TIR><Stavka>4.4890</Stavka><limit>10.0000</limit></SSU></SwapInfoSellXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapInfoSellXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapInfoSellXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesSwapMonthTotalXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "SwapMonthTotalXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 4)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlIn",
		DataStructureType: "SwapMonthTotalXML",
		Datastructure: datastructures.SwapMonthTotalXML{
			FromDate: "2022-02-11",
			ToDate:   "2022-02-24",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapMonthTotalXML xmlns="http://web.cbr.ru/"><fromDate>2022-02-11</fromDate><ToDate>2022-02-24</ToDate></SwapMonthTotalXML>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapMonthTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapMonthTotalXML")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.InputDataCases[0] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeBadRawData",
		DataStructureType: "SwapMonthTotalXML",
		Datastructure: datastructures.SwapMonthTotalXML{
			FromDate: "022-14-23",
			ToDate:   "2022-02-24",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadRawData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapMonthTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapMonthTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[1] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlNegativeFromDateAfterToDate",
		DataStructureType: "SwapMonthTotalXML",
		Datastructure: datastructures.SwapMonthTotalXML{
			FromDate: "2022-02-24",
			ToDate:   "2022-02-11",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: datastructures.ErrBadInputDateData,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapMonthTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapMonthTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[2] = newCase
	newCase = DatastructuresTestCase{
		Name:              "ValidateControlPositive",
		DataStructureType: "SwapMonthTotalXML",
		Datastructure: datastructures.SwapMonthTotalXML{
			FromDate: "2022-02-11",
			ToDate:   "2022-02-24",
			XMLNs:    "http://web.cbr.ru/",
		},
		NeedValidate:    true,
		ValidateControl: nil,
	}
	newCase.MarshalXMLTestFunc = func(_ *testing.T, _ interface{}, _ string) {}
	newCase.ValidateControlTestFunc = func(t *testing.T, Datastructure interface{}, ValidateControl error) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapMonthTotalXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapMonthTotalXML")
		}
		err := DSAssert.Validate()
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testSwapMonthTotalXMLResult := datastructures.SwapMonthTotalXMLResult{
		SMT: make([]datastructures.SwapMonthTotalXMLResultElem, 2),
	}
	testSwapMonthTotalXMLElem := datastructures.SwapMonthTotalXMLResultElem{
		D0:  time.Date(2022, time.February, 11, 0, 0, 0, 0, time.UTC),
		RUB: "41208.1",
		USD: "553.3",
	}
	testSwapMonthTotalXMLResult.SMT[0] = testSwapMonthTotalXMLElem
	testSwapMonthTotalXMLElem = datastructures.SwapMonthTotalXMLResultElem{
		D0:  time.Date(2022, time.February, 24, 0, 0, 0, 0, time.UTC),
		RUB: "24113.5",
		USD: "299.0",
	}
	testSwapMonthTotalXMLResult.SMT[1] = testSwapMonthTotalXMLElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "SwapMonthTotalXML",
		Datastructure:     testSwapMonthTotalXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<SwapMonthTotalXMLResult><SMT><D0>2022-02-11T00:00:00Z</D0><RUB>41208.1</RUB><USD>553.3</USD></SMT><SMT><D0>2022-02-24T00:00:00Z</D0><RUB>24113.5</RUB><USD>299.0</USD></SMT></SwapMonthTotalXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.SwapMonthTotalXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:SwapMonthTotalXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func initTestCasesAllDataInfoXML(t *testing.T) DatastructuresTestTable { // nolint:funlen, nolintlint
	t.Helper()
	DatastructuresTest := DatastructuresTestTable{}
	DatastructuresTest.MethodName = "AllDataInfoXML"
	DatastructuresTest.InputDataCases = make([]DatastructuresTestCase, 0)
	DatastructuresTest.OutputDataCases = make([]DatastructuresTestCase, 1)
	var newCase DatastructuresTestCase
	testAllDataInfoXMLResult := datastructures.AllDataInfoXMLResult{
		MainIndicatorsVR: datastructures.MainIndicatorsVRElem{
			Title: "Основные индикаторы финансового рынка",
			Currency: datastructures.CurrencyElem{
				Title: "Курсы валют",
				LUpd:  "",
				USD: datastructures.USDElem{
					OnDate: "29.08.2023",
					Curs:   "95.4717",
				},
				EUR: datastructures.EURElem{
					OnDate: "29.08.2023",
					Curs:   "103.2434",
				},
				CNY: datastructures.CNYElem{
					OnDate: "29.08.2023",
					Curs:   "13.0550",
				},
			},
			Metall: datastructures.MetallElem{
				Title:  "Драгоценные металлы",
				LUpd:   "",
				OnDate: "29.08.2023",
				Gold: datastructures.VoVStElem{
					Val:     "5879.60",
					Old_val: "5837.5100",
				},
				Silver: datastructures.VoVStElem{
					Val:     "74.24",
					Old_val: "73.6400",
				},
				Platinum: datastructures.VoVStElem{
					Val:     "2912.94",
					Old_val: "2841.0300",
				},
				Palladium: datastructures.VoVStElem{
					Val:     "3784.67",
					Old_val: "3788.0400",
				},
			},
			Inflation: datastructures.InflationElemADI{
				Title:  "Инфляция",
				LUpd:   "",
				OnDate: "01.07.2023",
				Val:    "4.30",
			},
			InflationTarget: datastructures.InflationTargetElem{
				Title:  "Цель по инфляции",
				LUpd:   "",
				OnDate: "01.01.2017",
				Val:    "4.0",
			},
			MBK: datastructures.MBKElem{
				Title: "Ставки межбанковского кредитного рынка",
				LUpd:  "",
				MIBID: datastructures.MBKStructElem{
					OnDate: "30.12.2016",
					D1: datastructures.VoVStElem{
						Val:     "9.79",
						Old_val: "9.79",
					},
					D2_7: datastructures.VoVStElem{
						Val:     "10.00",
						Old_val: "10.00",
					},
					D8_30: datastructures.VoVStElem{
						Val:     "9.93",
						Old_val: "9.93",
					},
				},
				MIBOR: datastructures.MBKStructElem{
					OnDate: "30.12.2016",
					D1: datastructures.VoVStElem{
						Val:     "10.54",
						Old_val: "10.54",
					},
					D2_7: datastructures.VoVStElem{
						Val:     "10.67",
						Old_val: "10.67",
					},
					D8_30: datastructures.VoVStElem{
						Val:     "11.06",
						Old_val: "11.06",
					},
				},
				MIACR: datastructures.MBKStructElem{
					OnDate: "25.08.2023",
					D1: datastructures.VoVStElem{
						Val:     "11.91",
						Old_val: "11.91",
					},
					D2_7: datastructures.VoVStElem{
						Val:     "12.39",
						Old_val: "10.67",
					},
					D8_30: datastructures.VoVStElem{
						Val:     "",
						Old_val: "",
					},
				},
				MIACRIG: datastructures.MBKStructElem{
					OnDate: "25.08.2023",
					D1: datastructures.VoVStElem{
						Val:     "11.95",
						Old_val: "11.95",
					},
					D2_7: datastructures.VoVStElem{
						Val:     "12.39",
						Old_val: "12.39",
					},
					D8_30: datastructures.VoVStElem{
						Val:     "",
						Old_val: "",
					},
				},
			},
			MosPrime: datastructures.MosPrimeElem{
				Title:  "MosPrime Rate",
				LUpd:   "",
				OnDate: "01.03.2022",
				D1: datastructures.VoVStElem{
					Val:     "",
					Old_val: "20.39",
				},
				M1: datastructures.VoVStElem{
					Val:     "",
					Old_val: "20.96",
				},
				M3: datastructures.VoVStElem{
					Val:     "",
					Old_val: "20.96",
				},
			},
		},
		KEY_RATE: datastructures.KEY_RATEElem{
			Title: "Действующая ключевая ставка",
			Val:   "12.00",
			Date:  "15.08.2023",
		},
		KEY_RATE_FUTURE: datastructures.KEY_RATE_FUTUREElem{
			Title:   "Новое значение ключевой ставки (справочно)",
			Val:     "12.00",
			NewDate: "15.08.2023",
		},
		REF_RATE: datastructures.TVStElem{
			Title: "Ставка рефинансирования (Значение соответствует значению ключевой ставки Банка России)",
			Val:   "12.00",
		},
		MBRStavki: datastructures.MBRStavkiElem{
			Title: "Параметры операций Банка России",
			Overnight_rate: datastructures.Overnight_rateElem{
				Title: "Ставка по кредиту overnight",
				LUpd:  "15.08.2023 11:14:15",
				Val1: datastructures.ValORElem{
					Date: "15.08.2023",
					Val:  "13.0",
				},
				Val2: datastructures.ValORElem{
					Date: "",
					Val:  "8",
				},
			},
			FixedLomb: datastructures.FixedLombElem{
				Title: "Фиксированные cтавки по ломбардным кредитам",
				LUpd:  "",
				D30: datastructures.FLElem{
					Date: "28.04.2014",
					Val:  "8.50",
				},
				D7: datastructures.FLElem{
					Date: "28.04.2014",
					Val:  "8.50",
				},
				D1: datastructures.FLElem{
					Date: "15.08.2023",
					Val:  "13.00",
				},
			},
			DepoRates: datastructures.DepoRatesElem{
				Title:  "Ставки по депозитным операциям",
				LUpd:   "29.08.2023 1:01:09",
				OnDate: "29.08.2023",
				TomNext: datastructures.VoVStElem{
					Val:     "",
					Old_val: "",
				},
				SpotNext: datastructures.VoVStElem{
					Val:     "",
					Old_val: "",
				},
				W1: datastructures.VoVStElem{
					Val:     "MIACR_B",
					Old_val: "",
				},
				W1_SPOT: datastructures.VoVStElem{
					Val:     "",
					Old_val: "",
				},
				CallDeposit: datastructures.VoVStElem{
					Val:     "",
					Old_val: "",
				},
			},
			SWAP: datastructures.SWAPElem{
				Title: "Своп-разница по валютному свопу",
				USD_RUB: datastructures.SWAPCurElem{
					LUpd:    "",
					Val:     "",
					Old_val: "0.0748",
				},
				EUR_RUB: datastructures.SWAPCurElem{
					LUpd:    "",
					Val:     "",
					Old_val: "0.0882",
				},
			},
			FixedRepoRate: datastructures.FixedRepoRateElem{
				Title: "Фиксированные cтавки по операциям прямого РЕПО",
				D1: datastructures.VStElem{
					Val: "13",
				},
				D7: datastructures.VStElem{
					Val: "13",
				},
			},
			MinimalRepoRates: datastructures.MinimalRepoRatesElem{
				Title:  "Параметры аукционов прямого РЕПО - Минимальные процентные ставки",
				LUpd:   "",
				OnDate: "15.08.2023",
				D1: datastructures.VStElem{
					Val: "12",
				},
				D7: datastructures.VStElem{
					Val: "12",
				},
			},
			MaxVolRepoOnAuction: datastructures.MaxVolMBRelem{
				Title:  "Максимальный объем средств, предоставляемых на первом аукционе прямого РЕПО",
				LUpd:   "",
				OnDate: "28.09.2015",
				Val:    "230",
			},
			MaxVolSwap: datastructures.MaxVolMBRelem{
				Title:  "Максимальный объем средств, предоставляемых по операциям 'валютный своп",
				LUpd:   "",
				OnDate: "20.09.2016",
				Val:    "620",
			},
		},
		Ko: datastructures.KoElem{
			Title: "Требования Банка России к кредитным организациям",
			OnOvernightCredit: datastructures.TLOVOStElem{
				Title:   "По кредитам overnight",
				LUpd:    "29.08.2023 9:18:46",
				OnDate:  "29.08.2023",
				Val:     "0.0",
				Old_val: "0.0",
			},
			OnLombardCredit: datastructures.TLOVOStElem{
				Title:   "По ломбардным кредитам",
				LUpd:    "29.08.2023 9:18:46",
				OnDate:  "29.08.2023",
				Val:     "14348.7",
				Old_val: "15348.7",
			},
			OnOtherCredit: datastructures.TLOVOStElem{
				Title:   "По другим кредитам",
				LUpd:    "29.08.2023 9:18:46",
				OnDate:  "29.08.2023",
				Val:     "1744136.5",
				Old_val: "874720.8",
			},
			OnDirectRepo: datastructures.OnDirectRepoElem{
				Title:  "По операциям прямого РЕПО",
				OnDate: "29.08.2023",
				OnAuction: datastructures.TVStElem{
					Title: "на аукционной основе",
					Val:   "1307685",
				},
				OnFixed: datastructures.TVStElem{
					Title: "по фиксированной ставке",
					Val:   "601",
				},
			},
			UnsecLoans: datastructures.TLOVOStElem{
				Title:   "По кредитам без обеспечения",
				LUpd:    "",
				OnDate:  "31.12.2010",
				Val:     "0",
				Old_val: "0",
			},
		},
		BankLikvid: datastructures.BankLikvidElem{
			Title: "Показатели банковской ликвидности",
			OstatKO: datastructures.OstatKOElem{
				Title:  "Сведения об остатках средств на корреспондентских счетах кредитных организаций",
				LUpd:   "29.08.2023 9:04:24",
				OnDate: "29.08.2023",
				Russ: datastructures.VoVStElem{
					Val:     "4769.8000",
					Old_val: "4356.7000",
				},
				Msk: datastructures.VoVStElem{
					Val:     "4530.5000",
					Old_val: "4123.9000",
				},
			},
			InDCredit: datastructures.TLOVOStElem{
				Title:   "Объем предоставленных внутридневных кредитов",
				LUpd:    "29.08.2023 9:18:46",
				OnDate:  "28.08.2023",
				Val:     "1486.62",
				Old_val: "334.55",
			},
			DepoBR: datastructures.TLOVOStElem{
				Title:   "Депозиты банков в Банке России",
				LUpd:    "29.08.2023 9:20:51",
				OnDate:  "29.08.2023",
				Val:     "2368.1896",
				Old_val: "2362.4110",
			},
			Saldo: datastructures.TLOVOStElem{
				Title:   "Сальдо операций Банка России по предоставлению /абсорбированию ликвидности",
				LUpd:    "29.08.2023 9:56:14",
				OnDate:  "29.08.2023",
				Val:     "-167.2",
				Old_val: "591.7",
			},
			VolOBR: datastructures.TVStElem{
				Title: "Объем рынка ОБР",
				Val:   "0",
			},
			VolDepo: datastructures.VolDepoElem{
				Title:  "Объем средств федерального бюджета, размещенных на депозитах коммерческих банков",
				OnDate: "05.03.2018",
				Val:    "0",
			},
		},
		Nor: datastructures.NorElem{
			Date:  "28.06.2023",
			Title: "Нормативы обязательных резервов",
			Ob_1: datastructures.Ob_1Elem{
				Title: "по обязательствам перед юридическими лицами – нерезидентами",
				Ob_1_1: datastructures.NorTLevelelem{
					Title:            "для банков с универсальной лицензией",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_1_2: datastructures.NorTLevelelem{
					Title:            "для небанковских кредитных организаций",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_1_3: datastructures.NorTLevelelem{
					Title:            "для банков с базовой лицензией",
					Val_rub:          "1.00",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
			},
			Ob_2: datastructures.Ob_2Elem{
				Title: "",
				Ob_2_1: datastructures.NorTLevelelem{
					Title:            "для банков с универсальной лицензией",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_2_2: datastructures.NorTLevelelem{
					Title:            "для небанковских кредитных организаций",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_2_3: datastructures.NorTLevelelem{
					Title:            "для банков с базовой лицензией",
					Val_rub:          "1.00",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
			},
			Ob_3: datastructures.Ob_3Elem{
				Title: "",
				Ob_3_1: datastructures.NorTLevelelem{
					Title:            "для банков с универсальной лицензией",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_3_2: datastructures.NorTLevelelem{
					Title:            "для небанковских кредитных организаций",
					Val_rub:          "4.50",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
				Ob_3_3: datastructures.NorTLevelelem{
					Title:            "для банков с базовой лицензией",
					Val_rub:          "1.00",
					Val_usd:          "8.50",
					Val_usd_excludUC: "6.00",
				},
			},
			Kor: datastructures.KorElem{
				Title: "Коэффициент усреднения обязательных резервов",
				Ku_1: datastructures.TVStElem{
					Title: "для банков с универсальной лицензией, банков с базовой лицензией",
					Val:   "0.9",
				},
				Ku_2: datastructures.TVStElem{
					Title: "для небанковских кредитных организаций",
					Val:   "1.0",
				},
			},
		},
		Macro: datastructures.MacroElem{
			Title: "Макроэкономические индикаторы",
			DB: datastructures.TVStElem{
				Title: "Денежная база",
				Val:   "11084.8",
			},
			DM: datastructures.TVStElem{
				Title: "Денежная масса (M2)",
				Val:   "36917.8",
			},
			M_rez: datastructures.M_rezElem{
				Title: "Международные резервы",
				Val:   "579.5",
				Date:  "18.08.2023",
			},
			Vol_GKO_OFZ: datastructures.TVStElem{
				Title: "Объем рынка ГКО-ОФЗ",
				Val:   "6741.11",
			},
		},
	}

	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControlOut",
		DataStructureType: "AllDataInfoXML",
		Datastructure:     testAllDataInfoXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<AllDataInfoXMLResult><MainIndicatorsVR Title="Основные индикаторы финансового рынка"><Currency Title="Курсы валют" LUpd=""><USD OnDate="29.08.2023"><curs>95.4717</curs></USD><EUR OnDate="29.08.2023"><curs>103.2434</curs></EUR><CNY OnDate="29.08.2023"><curs>13.0550</curs></CNY></Currency><Metall Title="Драгоценные металлы" LUpd="" OnDate="29.08.2023"><Золото val="5879.60" old_val="5837.5100"></Золото><Серебро val="74.24" old_val="73.6400"></Серебро><Платина val="2912.94" old_val="2841.0300"></Платина><Палладий val="3784.67" old_val="3788.0400"></Палладий></Metall><Inflation Title="Инфляция" LUpd="" OnDate="01.07.2023" val="4.30"></Inflation><InflationTarget Title="Цель по инфляции" LUpd="" OnDate="01.01.2017" val="4.0"></InflationTarget><MBK Title="Ставки межбанковского кредитного рынка" LUpd=""><MIBID OnDate="30.12.2016"><D1 val="9.79" old_val="9.79"></D1><D2_7 val="10.00" old_val="10.00"></D2_7><D8_30 val="9.93" old_val="9.93"></D8_30></MIBID><MIBOR OnDate="30.12.2016"><D1 val="10.54" old_val="10.54"></D1><D2_7 val="10.67" old_val="10.67"></D2_7><D8_30 val="11.06" old_val="11.06"></D8_30></MIBOR><MIACR OnDate="25.08.2023"><D1 val="11.91" old_val="11.91"></D1><D2_7 val="12.39" old_val="10.67"></D2_7><D8_30 val="" old_val=""></D8_30></MIACR><MIACR-IG OnDate="25.08.2023"><D1 val="11.95" old_val="11.95"></D1><D2_7 val="12.39" old_val="12.39"></D2_7><D8_30 val="" old_val=""></D8_30></MIACR-IG></MBK><MosPrime Title="MosPrime Rate" LUpd="" OnDate="01.03.2022"><D1 val="" old_val="20.39"></D1><M1 val="" old_val="20.96"></M1><M3 val="" old_val="20.96"></M3></MosPrime></MainIndicatorsVR><KEY_RATE Title="Действующая ключевая ставка" val="12.00" date="15.08.2023"></KEY_RATE><KEY_RATE_FUTURE Title="Новое значение ключевой ставки (справочно)" val="12.00" newdate="15.08.2023"></KEY_RATE_FUTURE><REF_RATE Title="Ставка рефинансирования (Значение соответствует значению ключевой ставки Банка России)" val="12.00"></REF_RATE><MBRStavki Title="Параметры операций Банка России"><Overnight_rate Title="Ставка по кредиту overnight" LUpd="15.08.2023 11:14:15"><Val1 Date="15.08.2023" val="13.0"></Val1><Val2 Date="" val="8"></Val2></Overnight_rate><FixedLomb Title="Фиксированные cтавки по ломбардным кредитам" LUpd=""><D30 Date="28.04.2014" val="8.50"></D30><D7 Date="28.04.2014" val="8.50"></D7><D1 Date="15.08.2023" val="13.00"></D1></FixedLomb><DepoRates Title="Ставки по депозитным операциям" LUpd="29.08.2023 1:01:09" OnDate="29.08.2023"><TomNext val="" old_val=""></TomNext><SpotNext val="" old_val=""></SpotNext><W1 val="MIACR_B" old_val=""></W1><W1_SPOT val="" old_val=""></W1_SPOT><CallDeposit val="" old_val=""></CallDeposit></DepoRates><SWAP Title="Своп-разница по валютному свопу"><USD_RUB LUpd="" val="" old_val="0.0748"></USD_RUB><EUR_RUB LUpd="" val="" old_val="0.0882"></EUR_RUB></SWAP><FixedRepoRate Title="Фиксированные cтавки по операциям прямого РЕПО"><D1 val="13"></D1><D7 val="13"></D7></FixedRepoRate><MinimalRepoRates Title="Параметры аукционов прямого РЕПО - Минимальные процентные ставки" LUpd="" OnDate="15.08.2023"><D1 val="12"></D1><D7 val="12"></D7></MinimalRepoRates><MaxVolRepoOnAuction Title="Максимальный объем средств, предоставляемых на первом аукционе прямого РЕПО" LUpd="" OnDate="28.09.2015" val="230"></MaxVolRepoOnAuction><MaxVolSwap Title="Максимальный объем средств, предоставляемых по операциям &#39;валютный своп" LUpd="" OnDate="20.09.2016" val="620"></MaxVolSwap></MBRStavki><Ko Title="Требования Банка России к кредитным организациям"><OnOvernightCredit Title="По кредитам overnight" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="0.0" old_val="0.0"></OnOvernightCredit><OnLombardCredit Title="По ломбардным кредитам" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="14348.7" old_val="15348.7"></OnLombardCredit><OnOtherCredit Title="По другим кредитам" LUpd="29.08.2023 9:18:46" OnDate="29.08.2023" val="1744136.5" old_val="874720.8"></OnOtherCredit><OnDirectRepo Title="По операциям прямого РЕПО" OnDate="29.08.2023"><OnAuction Title="на аукционной основе" val="1307685"></OnAuction><OnFixed Title="по фиксированной ставке" val="601"></OnFixed></OnDirectRepo><UnsecLoans Title="По кредитам без обеспечения" LUpd="" OnDate="31.12.2010" val="0" old_val="0"></UnsecLoans></Ko><BankLikvid Title="Показатели банковской ликвидности"><OstatKO Title="Сведения об остатках средств на корреспондентских счетах кредитных организаций" OnDate="29.08.2023" LUpd="29.08.2023 9:04:24"><Russ val="4769.8000" old_val="4356.7000"></Russ><Msk val="4530.5000" old_val="4123.9000"></Msk></OstatKO><InDCredit Title="Объем предоставленных внутридневных кредитов" LUpd="29.08.2023 9:18:46" OnDate="28.08.2023" val="1486.62" old_val="334.55"></InDCredit><DepoBR Title="Депозиты банков в Банке России" LUpd="29.08.2023 9:20:51" OnDate="29.08.2023" val="2368.1896" old_val="2362.4110"></DepoBR><Saldo Title="Сальдо операций Банка России по предоставлению /абсорбированию ликвидности" LUpd="29.08.2023 9:56:14" OnDate="29.08.2023" val="-167.2" old_val="591.7"></Saldo><VolOBR Title="Объем рынка ОБР" val="0"></VolOBR><VolDepo Title="Объем средств федерального бюджета, размещенных на депозитах коммерческих банков" OnDate="05.03.2018" val="0"></VolDepo></BankLikvid><Nor date="28.06.2023" Title="Нормативы обязательных резервов"><Ob_1 Title="по обязательствам перед юридическими лицами – нерезидентами"><Ob_1_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_1><Ob_1_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_2><Ob_1_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_1_3></Ob_1><Ob_2 Title=""><Ob_2_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_1><Ob_2_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_2><Ob_2_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_2_3></Ob_2><Ob_3 Title=""><Ob_3_1 Title="для банков с универсальной лицензией" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_1><Ob_3_2 Title="для небанковских кредитных организаций" val_rub="4.50" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_2><Ob_3_3 Title="для банков с базовой лицензией" val_rub="1.00" val_usd="8.50" val_usd_excludUC="6.00"></Ob_3_3></Ob_3><Kor Title="Коэффициент усреднения обязательных резервов"><Ku_1 Title="для банков с универсальной лицензией, банков с базовой лицензией" val="0.9"></Ku_1><Ku_2 Title="для небанковских кредитных организаций" val="1.0"></Ku_2></Kor></Nor><Macro Title="Макроэкономические индикаторы"><DB Title="Денежная база" val="11084.8"></DB><DM Title="Денежная масса (M2)" val="36917.8"></DM><M_rez Title="Международные резервы" val="579.5" date="18.08.2023"></M_rez><Vol_GKO_OFZ Title="Объем рынка ГКО-ОФЗ" val="6741.11"></Vol_GKO_OFZ></Macro></AllDataInfoXMLResult>`,
	}
	newCase.MarshalXMLTestFunc = func(t *testing.T, Datastructure interface{}, XMLMarshalControl string) {
		t.Helper()
		DSAssert, ok := Datastructure.(datastructures.AllDataInfoXMLResult)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:AllDataInfoXMLResult")
		}
		marshXMLres, err := xml.Marshal(DSAssert)
		require.NoError(t, err)
		require.Equal(t, XMLMarshalControl, string(marshXMLres))
	}
	newCase.ValidateControlTestFunc = func(_ *testing.T, _ interface{}, _ error) {}
	DatastructuresTest.OutputDataCases[0] = newCase
	return DatastructuresTest
}

func TestAllDatastructuresTableCases(t *testing.T) {
	AllDTTable := initAllDatastructuresTestTable(t)
	t.Parallel()
	for _, curTestTable := range AllDTTable {
		for _, curInputDataCase := range curTestTable.InputDataCases {
			curInputDataCase := curInputDataCase
			t.Run(curTestTable.MethodName+":"+curInputDataCase.Name, func(t *testing.T) {
				t.Parallel()
				curInputDataCase.MarshalXMLTest(t)
				curInputDataCase.ValidateControlTest(t)
			})
		}
		for _, curOutputDataCase := range curTestTable.OutputDataCases {
			curOutputDataCase := curOutputDataCase
			t.Run(curTestTable.MethodName+":"+curOutputDataCase.Name, func(t *testing.T) {
				t.Parallel()
				curOutputDataCase.MarshalXMLTest(t)
				curOutputDataCase.ValidateControlTest(t)
			})
		}
	}
}
