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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
		Name:              "XMLMarshalControl",
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
