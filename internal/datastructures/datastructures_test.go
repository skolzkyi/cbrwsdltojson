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
		err := DSAssert.Validate("2006-01-02")
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
		err := DSAssert.Validate("2006-01-02")
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
		err := DSAssert.Validate("2006-01-02")
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
		err := DSAssert.Validate("2006-01-02")
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
		DSAssert, ok := Datastructure.(datastructures.GetCursOnDateXML)
		if !ok {
			require.Fail(t, "fail type assertion in MarshalXMLTestFunc:BiCurBaseXML")
		}
		err := DSAssert.Validate("2006-01-02")
		require.Equal(t, ValidateControl, err)
	}
	DatastructuresTest.InputDataCases[3] = newCase
	testBiCurBaseXMLResult := datastructures.BiCurBaseXMLResult{
		BCB: make([]datastructures.BiCurBaseXMLResultElem, 2),
	}
	testBiCurBaseXMLResultElem := datastructures.BiCurBaseXMLResultElem{
		D0:  time.Date(2023, time.June, 22, 0, 0, 0, 0, time.Local),
		VAL: "87.736315",
	}
	testBiCurBaseXMLResult.BCB[0] = testBiCurBaseXMLResultElem
	testBiCurBaseXMLResultElem = datastructures.BiCurBaseXMLResultElem{
		D0:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.Local),
		VAL: "87.358585",
	}
	testBiCurBaseXMLResult.BCB[1] = testBiCurBaseXMLResultElem
	newCase = DatastructuresTestCase{
		Name:              "XMLMarshalControl",
		DataStructureType: "BiCurBaseXMLResult",
		Datastructure:     testBiCurBaseXMLResult,
		NeedXMLMarshal:    true,
		XMLMarshalControl: `<BiCurBaseXMLResult><BCB><D0>2023-06-22T00:00:00+03:00</D0><VAL>87.736315</VAL></BCB><BCB><D0>2023-06-23T00:00:00+03:00</D0><VAL>87.358585</VAL></BCB></BiCurBaseXMLResult>`,
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
