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
	var curDatastructuresTestTable DatastructuresTestTable
	AllDTTable := make(AllDatastructuresTestTable, 0)
	curDatastructuresTestTable = initTestCasesGetCursOnDateXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesBiCurBaseXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesBliquidityXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesDepoDynamicXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesDragMetDynamicXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesDVXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesEnumReutersValutesXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesEnumValutesXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesKeyRateXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesMainInfoXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesMrrf7DXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesMrrfXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesNewsInfoXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesOmodInfoXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesOstatDepoNewXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesOstatDepoXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesOstatDynamicXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
	curDatastructuresTestTable = initTestCasesOvernightXML(t)
	AllDTTable = append(AllDTTable, curDatastructuresTestTable)
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
			GoldBaks: 594,
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
		DataStructureType: "vernightXML",
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
