package app_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	app "github.com/skolzkyi/cbrwsdltojson/internal/app"
	"github.com/skolzkyi/cbrwsdltojson/internal/customsoap"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
	memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
	mocks "github.com/skolzkyi/cbrwsdltojson/internal/mocks"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Field1 string
	Field2 string
	Field3 int
}

type AllCasesTable struct {
	CasesByMethod []AppTestTable
}

type AppTestTable struct {
	MethodName string
	Method     func(*app.App, context.Context, interface{}, string) (interface{}, error)
	MethodWP   func(*app.App, context.Context) (interface{}, error)
	IsMethodWP bool
	TestCases  []AppTestCase
}

type AppTestCase struct {
	Input       interface{}
	Output      interface{}
	Name        string
	Error       error
	IsCacheTest bool
	IsCacheData bool
}

func initTestApp(t *testing.T) *app.App {
	t.Helper()
	var testApp *app.App
	loggerMock, err := mocks.NewLoggerMock(false)
	require.NoError(t, err)
	configMock := mocks.ConfigMock{}
	senderMock := mocks.SoapRequestSenderMock{}
	appMemcache := memcache.New()
	appMemcache.Init()
	testApp = app.New(loggerMock, &configMock, &senderMock, appMemcache, nil)
	return testApp
}

func TestPermittedReqSyncMap(t *testing.T) {
	testPermittedReqSyncMap := app.PermittedReqSyncMap{}
	testPermittedReqSyncMap.Init(nil)
	t.Parallel()
	t.Run("TestPermittedReqSyncMap: AddPermittedRequest_And_IsPermittedRequestInMap", func(t *testing.T) {
		t.Parallel()
		testPermittedReqSyncMap := app.PermittedReqSyncMap{}
		testPermittedReqSyncMap.Init(nil)
		testData := testPermittedReqSyncMap.IsPermittedRequestInMap("test1")
		require.Equal(t, false, testData)
		testPermittedReqSyncMap.AddPermittedRequest("test1")
		testData = testPermittedReqSyncMap.IsPermittedRequestInMap("test1")
		require.Equal(t, true, testData)
	})
	t.Run("TestPermittedReqSyncMap: AddPermittedRequest_And_PermittedRequestMapLength", func(t *testing.T) {
		t.Parallel()
		testPermittedReqSyncMap := app.PermittedReqSyncMap{}
		testPermittedReqSyncMap.Init(nil)
		testPermittedReqSyncMap.AddPermittedRequest("test1")
		testPermittedReqSyncMap.AddPermittedRequest("test2")
		testData := testPermittedReqSyncMap.PermittedRequestMapLength()
		require.Equal(t, 2, testData)
	})
}

func TestGenerateTagForMemCacheLogic(t *testing.T) {
	testApp := initTestApp(t)
	testStruct1 := testStruct{
		Field1: "abc",
		Field2: "def",
		Field3: 0,
	}
	testStruct2 := testStruct{
		Field1: "123",
		Field2: "456",
		Field3: 1,
	}
	err := testApp.AddOrUpdateDataInCache("ts1", testStruct1, testStruct1.Field3)
	require.NoError(t, err)
	err = testApp.AddOrUpdateDataInCache("ts2", testStruct2, testStruct2.Field3)
	require.NoError(t, err)
	rawBody, err := json.Marshal(testStruct1)
	require.NoError(t, err)
	payload1, ok := testApp.GetDataInCacheIfExisting("ts1", string(rawBody))
	require.Equal(t, true, ok)
	data1, ok := payload1.(int)
	require.Equal(t, true, ok)
	require.Equal(t, testStruct1.Field3, data1)
	rawBody, err = json.Marshal(testStruct2)
	require.NoError(t, err)
	payload2, ok := testApp.GetDataInCacheIfExisting("ts2", string(rawBody))
	require.Equal(t, true, ok)
	data2, ok := payload2.(int)
	require.Equal(t, true, ok)
	require.Equal(t, testStruct2.Field3, data2)
}

func createStandartTestCacheCases(t *testing.T, input interface{}, output interface{}) []AppTestCase {
	t.Helper()
	standartTestCacheCases := make([]AppTestCase, 2)

	standartTestCacheCases[0] = AppTestCase{
		Name:        "InCacheTest",
		IsCacheTest: true,
		IsCacheData: true,
		Input:       input,
		Output:      output,
	}

	standartTestCacheCases[1] = AppTestCase{
		Name:        "NotInCacheTest",
		IsCacheTest: true,
		IsCacheData: false,
		Input:       input,
		Output:      output,
	}
	return standartTestCacheCases
}

func getTagForCache(t *testing.T, SOAPMethod string, request interface{}) string { //nolint: gocritic
	t.Helper()
	if SOAPMethod == "Repo_debtXML" {
		SOAPMethod = "RepoDebtXML" // fix changing name real SOAP method and Handlename TODO
	}
	jsonstring, err := json.Marshal(request)
	require.NoError(t, err)
	return SOAPMethod + string(jsonstring)
}

// GetCursOnDate.
func initTestDataGetCursOnDateXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataGetCursOnDate := AppTestTable{
		MethodName: "GetCursOnDateXML",
		Method:     (*app.App).GetCursOnDateXML,
	}
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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.GetCursOnDateXML{
			OnDate: "2023-06-22",
		},
		Output: testGetCursOnDateXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.GetCursOnDateXML{
			OnDate: "023-14-22",
		},
		Output: datastructures.GetCursOnDateXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.GetCursOnDateXML{
		OnDate: "2023-06-22",
	}, testGetCursOnDateXMLResult)
	testDataGetCursOnDate.TestCases = testCases
	testDataGetCursOnDate.TestCases = append(testDataGetCursOnDate.TestCases, standartTestCacheCases...)

	return testDataGetCursOnDate
}

// BiCurBaseXML.
func initTestDataBiCurBaseXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataBiCurBaseXML := AppTestTable{
		MethodName: "BiCurBaseXML",
		Method:     (*app.App).BiCurBaseXML,
	}
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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.BiCurBaseXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testBiCurBaseXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.BiCurBaseXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.BiCurBaseXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.BiCurBaseXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testBiCurBaseXMLResult)
	testDataBiCurBaseXML.TestCases = testCases
	testDataBiCurBaseXML.TestCases = append(testDataBiCurBaseXML.TestCases, standartTestCacheCases...)

	return testDataBiCurBaseXML
}

// BliquidityXML.
func initTestDataBliquidityXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataBliquidityXML := AppTestTable{
		MethodName: "BliquidityXML",
		Method:     (*app.App).BliquidityXML,
	}
	testBliquidityXMLResult := datastructures.BliquidityXMLResult{
		BL: make([]datastructures.BliquidityXMLResultElem, 2),
	}
	testBliquidityXMLResultElem := datastructures.BliquidityXMLResultElem{
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
	testBliquidityXMLResult.BL[0] = testBliquidityXMLResultElem
	testBliquidityXMLResultElem = datastructures.BliquidityXMLResultElem{
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
	testBliquidityXMLResult.BL[1] = testBliquidityXMLResultElem
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.BliquidityXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testBliquidityXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.BliquidityXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.BliquidityXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.BliquidityXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testBliquidityXMLResult)
	testDataBliquidityXML.TestCases = testCases
	testDataBliquidityXML.TestCases = append(testDataBliquidityXML.TestCases, standartTestCacheCases...)

	return testDataBliquidityXML
}

// DepoDynamicXML.
func initTestDataDepoDynamicXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataDepoDynamicXML := AppTestTable{
		MethodName: "DepoDynamicXML",
		Method:     (*app.App).DepoDynamicXML,
	}
	testDepoDynamicXMLResult := datastructures.DepoDynamicXMLResult{
		Depo: make([]datastructures.DepoDynamicXMLResultElem, 2),
	}
	testDepoDynamicXMLResultElem := datastructures.DepoDynamicXMLResultElem{
		DateDepo:  time.Date(2023, time.June, 22, 0, 0, 0, 0, time.UTC),
		Overnight: "6.50",
	}
	testDepoDynamicXMLResult.Depo[0] = testDepoDynamicXMLResultElem
	testDepoDynamicXMLResultElem = datastructures.DepoDynamicXMLResultElem{
		DateDepo:  time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC),
		Overnight: "6.50",
	}
	testDepoDynamicXMLResult.Depo[1] = testDepoDynamicXMLResultElem
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.DepoDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testDepoDynamicXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.DepoDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.DepoDynamicXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.DepoDynamicXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testDepoDynamicXMLResult)
	testDataDepoDynamicXML.TestCases = testCases
	testDataDepoDynamicXML.TestCases = append(testDataDepoDynamicXML.TestCases, standartTestCacheCases...)

	return testDataDepoDynamicXML
}

// DragMetDynamicXML.
func initTestDragMetDynamicXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataDragMetDynamicXML := AppTestTable{
		MethodName: "DragMetDynamicXML",
		Method:     (*app.App).DragMetDynamicXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.DragMetDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testDragMetDynamicXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.DragMetDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.DragMetDynamicXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.DragMetDynamicXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testDragMetDynamicXMLResult)
	testDataDragMetDynamicXML.TestCases = testCases
	testDataDragMetDynamicXML.TestCases = append(testDataDragMetDynamicXML.TestCases, standartTestCacheCases...)

	return testDataDragMetDynamicXML
}

// DVXML.
func initTestDataDVXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataDVXML := AppTestTable{
		MethodName: "DVXML",
		Method:     (*app.App).DVXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.DVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testDVXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.DVXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.DVXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.DVXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testDVXMLResult)
	testDataDVXML.TestCases = testCases
	testDataDVXML.TestCases = append(testDataDVXML.TestCases, standartTestCacheCases...)

	return testDataDVXML
}

// EnumReutersValutesXML.
func initTestDataEnumReutersValutesXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataEnumReutersValutesXML := AppTestTable{
		MethodName: "EnumReutersValutesXML",
		MethodWP:   (*app.App).EnumReutersValutesXML,
		IsMethodWP: true,
	}
	testEnumReutersValutesXMLResult := datastructures.EnumReutersValutesXMLResult{
		EnumRValutes: make([]datastructures.EnumReutersValutesXMLResultElem, 2),
	}
	testEnumReutersValutesXMLElem := datastructures.EnumReutersValutesXMLResultElem{
		Num_code:  8,
		Char_code: "ALL",
		Title_ru:  "Албанский лек",
		Title_en:  "Albanian Lek",
	}
	testEnumReutersValutesXMLResult.EnumRValutes[0] = testEnumReutersValutesXMLElem
	testEnumReutersValutesXMLElem = datastructures.EnumReutersValutesXMLResultElem{
		Num_code:  12,
		Char_code: "DZD",
		Title_ru:  "Алжирский динар",
		Title_en:  "Algerian Dinar",
	}
	testEnumReutersValutesXMLResult.EnumRValutes[1] = testEnumReutersValutesXMLElem
	testCases := make([]AppTestCase, 1)
	testCases[0] = AppTestCase{
		Name:   "Positive",
		Input:  &datastructures.EnumReutersValutesXML{},
		Output: testEnumReutersValutesXMLResult,
		Error:  nil,
	}

	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.EnumReutersValutesXML{}, testEnumReutersValutesXMLResult)
	testDataEnumReutersValutesXML.TestCases = testCases
	testDataEnumReutersValutesXML.TestCases = append(testDataEnumReutersValutesXML.TestCases, standartTestCacheCases...)

	return testDataEnumReutersValutesXML
}

// EnumValutesXML.
func initTestDataEnumValutesXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataEnumValutesXML := AppTestTable{
		MethodName: "EnumValutesXML",
		Method:     (*app.App).EnumValutesXML,
	}
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

	testCases := make([]AppTestCase, 1)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.EnumValutesXML{
			Seld: false,
		},
		Output: testEnumValutesXMLResult,
		Error:  nil,
	}

	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.EnumValutesXML{}, testEnumValutesXMLResult)
	testDataEnumValutesXML.TestCases = testCases
	testDataEnumValutesXML.TestCases = append(testDataEnumValutesXML.TestCases, standartTestCacheCases...)

	return testDataEnumValutesXML
}

// KeyRateXML.
func initTestDataKeyRateXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataDVXML := AppTestTable{
		MethodName: "KeyRateXML",
		Method:     (*app.App).KeyRateXML,
	}
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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.KeyRateXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testKeyRateXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.KeyRateXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.KeyRateXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.KeyRateXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testKeyRateXMLResult)
	testDataDVXML.TestCases = testCases
	testDataDVXML.TestCases = append(testDataDVXML.TestCases, standartTestCacheCases...)

	return testDataDVXML
}

// MainInfoXML.
func initTestDataMainInfoXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataMainInfoXML := AppTestTable{
		MethodName: "MainInfoXML",
		MethodWP:   (*app.App).MainInfoXML,
		IsMethodWP: true,
	}
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
	testCases := make([]AppTestCase, 1)
	testCases[0] = AppTestCase{
		Name:   "Positive",
		Input:  &datastructures.MainInfoXML{},
		Output: testMainInfoXMLResult,
		Error:  nil,
	}

	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.MainInfoXML{}, testMainInfoXMLResult)
	testDataMainInfoXML.TestCases = testCases
	testDataMainInfoXML.TestCases = append(testDataMainInfoXML.TestCases, standartTestCacheCases...)

	return testDataMainInfoXML
}

// mrrf7DXML.
func initTestDataMrrf7DXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataMrrf7DXML := AppTestTable{
		MethodName: "mrrf7DXML",
		Method:     (*app.App).Mrrf7DXML,
	}
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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.Mrrf7DXML{
			FromDate: "2023-06-15",
			ToDate:   "2023-06-23",
		},
		Output: testMrrf7DXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.Mrrf7DXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.Mrrf7DXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.Mrrf7DXML{
		FromDate: "2023-06-15",
		ToDate:   "2023-06-23",
	}, testMrrf7DXMLResult)
	testDataMrrf7DXML.TestCases = testCases
	testDataMrrf7DXML.TestCases = append(testDataMrrf7DXML.TestCases, standartTestCacheCases...)

	return testDataMrrf7DXML
}

// mrrfXML.
func initTestDataMrrfXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataMrrfXML := AppTestTable{
		MethodName: "mrrfXML",
		Method:     (*app.App).MrrfXML,
	}

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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.MrrfXML{
			FromDate: "2023-05-01",
			ToDate:   "2023-06-23",
		},
		Output: testMrrfXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.MrrfXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.MrrfXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.MrrfXML{
		FromDate: "2023-05-01",
		ToDate:   "2023-06-23",
	}, testMrrfXMLResult)
	testDataMrrfXML.TestCases = testCases
	testDataMrrfXML.TestCases = append(testDataMrrfXML.TestCases, standartTestCacheCases...)

	return testDataMrrfXML
}

// NewsInfoXML.
func initTestDataNewsInfoXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataNewsInfoXML := AppTestTable{
		MethodName: "NewsInfoXML",
		Method:     (*app.App).NewsInfoXML,
	}

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

	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.NewsInfoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testNewsInfoXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.NewsInfoXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.NewsInfoXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.NewsInfoXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testNewsInfoXMLResult)
	testDataNewsInfoXML.TestCases = testCases
	testDataNewsInfoXML.TestCases = append(testDataNewsInfoXML.TestCases, standartTestCacheCases...)

	return testDataNewsInfoXML
}

// OmodInfoXML.
func initTestDataOmodInfoXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataOmodInfoXML := AppTestTable{
		MethodName: "OmodInfoXML",
		MethodWP:   (*app.App).OmodInfoXML,
		IsMethodWP: true,
	}
	testOmodInfoXMLResult := datastructures.OmodInfoXMLResult{
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
	testCases := make([]AppTestCase, 1)
	testCases[0] = AppTestCase{
		Name:   "Positive",
		Input:  &datastructures.OmodInfoXML{},
		Output: testOmodInfoXMLResult,
		Error:  nil,
	}

	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.OmodInfoXML{}, testOmodInfoXMLResult)
	testDataOmodInfoXML.TestCases = testCases
	testDataOmodInfoXML.TestCases = append(testDataOmodInfoXML.TestCases, standartTestCacheCases...)

	return testDataOmodInfoXML
}

// OstatDepoNewXML.
func initTestDataOstatDepoNewXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataOstatDepoNewXML := AppTestTable{
		MethodName: "OstatDepoNewXML",
		Method:     (*app.App).OstatDepoNewXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.OstatDepoNewXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testOstatDepoNewXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.OstatDepoNewXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.OstatDepoNewXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.OstatDepoNewXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testOstatDepoNewXMLResult)

	testDataOstatDepoNewXML.TestCases = testCases
	testDataOstatDepoNewXML.TestCases = append(testDataOstatDepoNewXML.TestCases, standartTestCacheCases...)

	return testDataOstatDepoNewXML
}

// OstatDepoXML.
func initTestDataOstatDepoXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataOstatDepoXML := AppTestTable{
		MethodName: "OstatDepoXML",
		Method:     (*app.App).OstatDepoXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.OstatDepoXML{
			FromDate: "2022-12-29",
			ToDate:   "2022-12-30",
		},
		Output: testOstatDepoXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.OstatDepoXML{
			FromDate: "022-14-22",
			ToDate:   "2022-12-30",
		},
		Output: datastructures.OstatDepoXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.OstatDepoXML{
		FromDate: "2022-12-29",
		ToDate:   "2022-12-30",
	}, testOstatDepoXMLResult)

	testDataOstatDepoXML.TestCases = testCases
	testDataOstatDepoXML.TestCases = append(testDataOstatDepoXML.TestCases, standartTestCacheCases...)

	return testDataOstatDepoXML
}

// OstatDynamicXML.
func initTestDataOstatDynamicXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataOstatDynamicXML := AppTestTable{
		MethodName: "OstatDynamicXML",
		Method:     (*app.App).OstatDynamicXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.OstatDynamicXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testOstatDynamicXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.OstatDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.OstatDynamicXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.OstatDynamicXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testOstatDynamicXMLResult)

	testDataOstatDynamicXML.TestCases = testCases
	testDataOstatDynamicXML.TestCases = append(testDataOstatDynamicXML.TestCases, standartTestCacheCases...)

	return testDataOstatDynamicXML
}

// OvernightXML.
func initTestDataOvernightXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataOvernightXML := AppTestTable{
		MethodName: "OvernightXML",
		Method:     (*app.App).OvernightXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.OvernightXML{
			FromDate: "2023-07-22",
			ToDate:   "2023-08-16",
		},
		Output: testOvernightXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.OvernightXML{
			FromDate: "022-14-22",
			ToDate:   "2023-08-16",
		},
		Output: datastructures.OvernightXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.OvernightXML{
		FromDate: "2023-07-22",
		ToDate:   "2023-08-16",
	}, testOvernightXMLResult)

	testDataOvernightXML.TestCases = testCases
	testDataOvernightXML.TestCases = append(testDataOvernightXML.TestCases, standartTestCacheCases...)

	return testDataOvernightXML
}

// Repo_debtXML.
func initTestDataRepo_debtXML(t *testing.T) AppTestTable { //nolint:revive, stylecheck, nolintlint
	t.Helper()
	testDataORepo_debtXML := AppTestTable{ //nolint:revive, stylecheck, nolintlint
		MethodName: "Repo_debtXML",
		Method:     (*app.App).RepoDebtXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.Repo_debtXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testRepo_debtXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.Repo_debtXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.Repo_debtXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.Repo_debtXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testRepo_debtXMLResult)

	testDataORepo_debtXML.TestCases = testCases
	testDataORepo_debtXML.TestCases = append(testDataORepo_debtXML.TestCases, standartTestCacheCases...)

	return testDataORepo_debtXML
}

// RepoDebtUSDXML.
func initTestDataRepoDebtUSDXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataRepoDebtUSDXML := AppTestTable{
		MethodName: "RepoDebtUSDXML",
		Method:     (*app.App).RepoDebtUSDXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.RepoDebtUSDXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testRepoDebtUSDXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.RepoDebtUSDXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.RepoDebtUSDXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.RepoDebtUSDXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testRepoDebtUSDXMLResult)

	testDataRepoDebtUSDXML.TestCases = testCases
	testDataRepoDebtUSDXML.TestCases = append(testDataRepoDebtUSDXML.TestCases, standartTestCacheCases...)

	return testDataRepoDebtUSDXML
}

// ROISfixXML.
func initTestDataROISfixXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataROISfixXML := AppTestTable{
		MethodName: "ROISfixXML",
		Method:     (*app.App).ROISfixXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.ROISfixXML{
			FromDate: "2022-02-27",
			ToDate:   "2022-03-02",
		},
		Output: testROISfixXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.ROISfixXML{
			FromDate: "022-14-22",
			ToDate:   "2022-03-02",
		},
		Output: datastructures.ROISfixXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.ROISfixXML{
		FromDate: "2022-02-27",
		ToDate:   "2022-03-02",
	}, testROISfixXMLResult)

	testDataROISfixXML.TestCases = testCases
	testDataROISfixXML.TestCases = append(testDataROISfixXML.TestCases, standartTestCacheCases...)

	return testDataROISfixXML
}

// RuoniaSVXML.
func initTestDataRuoniaSVXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataRuoniaSVXML := AppTestTable{
		MethodName: "RuoniaSVXML",
		Method:     (*app.App).RuoniaSVXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.RuoniaSVXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testRuoniaSVXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.RuoniaSVXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.RuoniaSVXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.RuoniaSVXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testRuoniaSVXMLResult)

	testDataRuoniaSVXML.TestCases = testCases
	testDataRuoniaSVXML.TestCases = append(testDataRuoniaSVXML.TestCases, standartTestCacheCases...)

	return testDataRuoniaSVXML
}

// RuoniaXML.
func initTestDataRuoniaXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataRuoniaXML := AppTestTable{
		MethodName: "RuoniaXML",
		Method:     (*app.App).RuoniaXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.RuoniaXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testRuoniaXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.RuoniaXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.RuoniaXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.RuoniaXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testRuoniaXMLResult)

	testDataRuoniaXML.TestCases = testCases
	testDataRuoniaXML.TestCases = append(testDataRuoniaXML.TestCases, standartTestCacheCases...)

	return testDataRuoniaXML
}

// SaldoXML.
func initTestDataSaldoXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSaldoXML := AppTestTable{
		MethodName: "SaldoXML",
		Method:     (*app.App).SaldoXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SaldoXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
		},
		Output: testSaldoXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SaldoXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
		},
		Output: datastructures.SaldoXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SaldoXML{
		FromDate: "2023-06-22",
		ToDate:   "2023-06-23",
	}, testSaldoXMLResult)

	testDataSaldoXML.TestCases = testCases
	testDataSaldoXML.TestCases = append(testDataSaldoXML.TestCases, standartTestCacheCases...)

	return testDataSaldoXML
}

// SwapDayTotalXML.
func initTestDataSwapDayTotalXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapDayTotalXML := AppTestTable{
		MethodName: "SwapDayTotalXML",
		Method:     (*app.App).SwapDayTotalXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapDayTotalXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
		},
		Output: testSwapDayTotalXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapDayTotalXML{
			FromDate: "022-14-22",
			ToDate:   "2022-02-28",
		},
		Output: datastructures.SwapDayTotalXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapDayTotalXML{
		FromDate: "2022-02-25",
		ToDate:   "2022-02-28",
	}, testSwapDayTotalXMLResult)

	testDataSwapDayTotalXML.TestCases = testCases
	testDataSwapDayTotalXML.TestCases = append(testDataSwapDayTotalXML.TestCases, standartTestCacheCases...)

	return testDataSwapDayTotalXML
}

// SwapDynamicXML.
func initTestDataSwapDynamicXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapDynamicXML := AppTestTable{
		MethodName: "SwapDynamicXML",
		Method:     (*app.App).SwapDynamicXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapDynamicXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
		},
		Output: testSwapDynamicXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapDynamicXML{
			FromDate: "022-14-22",
			ToDate:   "2022-02-28",
		},
		Output: datastructures.SwapDynamicXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapDynamicXML{
		FromDate: "2022-02-25",
		ToDate:   "2022-02-28",
	}, testSwapDynamicXMLResult)

	testDataSwapDynamicXML.TestCases = testCases
	testDataSwapDynamicXML.TestCases = append(testDataSwapDynamicXML.TestCases, standartTestCacheCases...)

	return testDataSwapDynamicXML
}

// SwapInfoSellUSDVolXML.
func initTestDataSwapInfoSellUSDVolXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapInfoSellUSDVolXML := AppTestTable{
		MethodName: "SwapInfoSellUSDVolXML",
		Method:     (*app.App).SwapInfoSellUSDVolXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapInfoSellUSDVolXML{
			FromDate: "2022-02-24",
			ToDate:   "2022-02-28",
		},
		Output: testSwapInfoSellUSDVolXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapInfoSellUSDVolXML{
			FromDate: "022-14-22",
			ToDate:   "2022-02-28",
		},
		Output: datastructures.SwapInfoSellUSDVolXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapInfoSellUSDVolXML{
		FromDate: "2022-02-24",
		ToDate:   "2022-02-28",
	}, testSwapInfoSellUSDVolXMLResult)

	testDataSwapInfoSellUSDVolXML.TestCases = testCases
	testDataSwapInfoSellUSDVolXML.TestCases = append(testDataSwapInfoSellUSDVolXML.TestCases, standartTestCacheCases...)

	return testDataSwapInfoSellUSDVolXML
}

// SwapInfoSellUSDXML.
func initTestDataSwapInfoSellUSDXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapInfoSellUSDXML := AppTestTable{
		MethodName: "SwapInfoSellUSDXML",
		Method:     (*app.App).SwapInfoSellUSDXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapInfoSellUSDXML{
			FromDate: "2022-02-25",
			ToDate:   "2022-02-28",
		},
		Output: testSwapInfoSellUSDXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapInfoSellUSDXML{
			FromDate: "022-14-22",
			ToDate:   "2022-02-28",
		},
		Output: datastructures.SwapInfoSellUSDXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapInfoSellUSDXML{
		FromDate: "2022-02-25",
		ToDate:   "2022-02-28",
	}, testSwapInfoSellUSDXMLResult)

	testDataSwapInfoSellUSDXML.TestCases = testCases
	testDataSwapInfoSellUSDXML.TestCases = append(testDataSwapInfoSellUSDXML.TestCases, standartTestCacheCases...)

	return testDataSwapInfoSellUSDXML
}

// SwapInfoSellVolXML.
func initTestDataSwapInfoSellVolXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapInfoSellVolXML := AppTestTable{
		MethodName: "SwapInfoSellVolXML",
		Method:     (*app.App).SwapInfoSellVolXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapInfoSellVolXML{
			FromDate: "2023-05-05",
			ToDate:   "2023-05-10",
		},
		Output: testSwapInfoSellVolXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapInfoSellVolXML{
			FromDate: "022-14-22",
			ToDate:   "2023-05-10",
		},
		Output: datastructures.SwapInfoSellVolXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapInfoSellVolXML{
		FromDate: "2023-05-05",
		ToDate:   "2023-05-10",
	}, testSwapInfoSellVolXMLResult)

	testDataSwapInfoSellVolXML.TestCases = testCases
	testDataSwapInfoSellVolXML.TestCases = append(testDataSwapInfoSellVolXML.TestCases, standartTestCacheCases...)

	return testDataSwapInfoSellVolXML
}

// SwapInfoSellXML.
func initTestDataSwapInfoSellXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapInfoSellXML := AppTestTable{
		MethodName: "SwapInfoSellXML",
		Method:     (*app.App).SwapInfoSellXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapInfoSellXML{
			FromDate: "2023-06-20",
			ToDate:   "2023-06-21",
		},
		Output: testSwapInfoSellXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapInfoSellXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-21",
		},
		Output: datastructures.SwapInfoSellXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapInfoSellXML{
		FromDate: "2023-06-20",
		ToDate:   "2023-06-21",
	}, testSwapInfoSellXMLResult)

	testDataSwapInfoSellXML.TestCases = testCases
	testDataSwapInfoSellXML.TestCases = append(testDataSwapInfoSellXML.TestCases, standartTestCacheCases...)

	return testDataSwapInfoSellXML
}

// SwapMonthTotalXML.
func initTestDataSwapMonthTotalXML(t *testing.T) AppTestTable {
	t.Helper()
	testDataSwapMonthTotalXML := AppTestTable{
		MethodName: "SwapMonthTotalXML",
		Method:     (*app.App).SwapMonthTotalXML,
	}
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
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: &datastructures.SwapMonthTotalXML{
			FromDate: "2022-02-11",
			ToDate:   "2022-02-24",
		},
		Output: testSwapMonthTotalXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: &datastructures.SwapMonthTotalXML{
			FromDate: "022-14-22",
			ToDate:   "2022-02-24",
		},
		Output: datastructures.SwapMonthTotalXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.SwapMonthTotalXML{
		FromDate: "2022-02-11",
		ToDate:   "2022-02-24",
	}, testSwapMonthTotalXMLResult)

	testDataSwapMonthTotalXML.TestCases = testCases
	testDataSwapMonthTotalXML.TestCases = append(testDataSwapMonthTotalXML.TestCases, standartTestCacheCases...)

	return testDataSwapMonthTotalXML
}

// AllDataInfoXML.
func initTestDataAllDataInfoXML(t *testing.T) AppTestTable { //nolint:gocognit, nolintlint, gocyclo, funlen
	t.Helper()
	testDataAllDataInfoXML := AppTestTable{
		MethodName: "AllDataInfoXML",
		MethodWP:   (*app.App).AllDataInfoXML,
		IsMethodWP: true,
	}
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
	testCases := make([]AppTestCase, 1)
	testCases[0] = AppTestCase{
		Name:   "Positive",
		Input:  &datastructures.AllDataInfoXML{},
		Output: testAllDataInfoXMLResult,
		Error:  nil,
	}

	standartTestCacheCases := createStandartTestCacheCases(t, &datastructures.AllDataInfoXML{}, testAllDataInfoXMLResult)
	testDataAllDataInfoXML.TestCases = testCases
	testDataAllDataInfoXML.TestCases = append(testDataAllDataInfoXML.TestCases, standartTestCacheCases...)

	return testDataAllDataInfoXML
}

func TestAllAppCases(t *testing.T) { //nolint:gocognit, nolintlint, gocyclo, funlen
	acTable := AllCasesTable{}
	acTable.CasesByMethod = make([]AppTestTable, 32)
	acTable.CasesByMethod[0] = initTestDataGetCursOnDateXML(t)
	acTable.CasesByMethod[1] = initTestDataBiCurBaseXML(t)
	acTable.CasesByMethod[2] = initTestDataBliquidityXML(t)
	acTable.CasesByMethod[3] = initTestDataDepoDynamicXML(t)
	acTable.CasesByMethod[4] = initTestDragMetDynamicXML(t)
	acTable.CasesByMethod[5] = initTestDataDVXML(t)
	acTable.CasesByMethod[6] = initTestDataEnumReutersValutesXML(t)
	acTable.CasesByMethod[7] = initTestDataEnumValutesXML(t)
	acTable.CasesByMethod[8] = initTestDataKeyRateXML(t)
	acTable.CasesByMethod[9] = initTestDataMainInfoXML(t)
	acTable.CasesByMethod[10] = initTestDataMrrf7DXML(t)
	acTable.CasesByMethod[11] = initTestDataMrrfXML(t)
	acTable.CasesByMethod[12] = initTestDataNewsInfoXML(t)
	acTable.CasesByMethod[13] = initTestDataOmodInfoXML(t)
	acTable.CasesByMethod[14] = initTestDataOstatDepoNewXML(t)
	acTable.CasesByMethod[15] = initTestDataOstatDepoXML(t)
	acTable.CasesByMethod[16] = initTestDataOstatDynamicXML(t)
	acTable.CasesByMethod[17] = initTestDataOvernightXML(t)
	acTable.CasesByMethod[18] = initTestDataRepo_debtXML(t)
	acTable.CasesByMethod[19] = initTestDataRepoDebtUSDXML(t)
	acTable.CasesByMethod[20] = initTestDataROISfixXML(t)
	acTable.CasesByMethod[21] = initTestDataRuoniaSVXML(t)
	acTable.CasesByMethod[22] = initTestDataRuoniaXML(t)
	acTable.CasesByMethod[23] = initTestDataSaldoXML(t)
	acTable.CasesByMethod[24] = initTestDataSwapDayTotalXML(t)
	acTable.CasesByMethod[25] = initTestDataSwapDynamicXML(t)
	acTable.CasesByMethod[26] = initTestDataSwapInfoSellUSDVolXML(t)
	acTable.CasesByMethod[27] = initTestDataSwapInfoSellUSDXML(t)
	acTable.CasesByMethod[28] = initTestDataSwapInfoSellVolXML(t)
	acTable.CasesByMethod[29] = initTestDataSwapInfoSellXML(t)
	acTable.CasesByMethod[30] = initTestDataSwapMonthTotalXML(t)
	acTable.CasesByMethod[31] = initTestDataAllDataInfoXML(t)
	t.Parallel()
	for _, curMethodTable := range acTable.CasesByMethod {
		curMethodTable := curMethodTable
		for _, curTestCase := range curMethodTable.TestCases {
			curTestCase := curTestCase
			t.Run(curMethodTable.MethodName+":"+curTestCase.Name, func(t *testing.T) {
				t.Parallel()
				var testRes interface{}
				var cachedData memcache.CacheInfo
				var rawBody []byte
				var err error
				var ok bool
				testApp := initTestApp(t)
				if !curMethodTable.IsMethodWP {
					rawBody, err = json.Marshal(curTestCase.Input)
					require.NoError(t, err)
					testRes, err = curMethodTable.Method(testApp, context.Background(), curTestCase.Input, string(rawBody))
				} else {
					testRes, err = curMethodTable.MethodWP(testApp, context.Background())
					require.NoError(t, err)
				}
				if err == nil {
					// testApp.Appmemcache.PrintAllCacheKeys()
					var cacheTag string
					if curMethodTable.IsMethodWP {
						cacheTag = curMethodTable.MethodName
					} else {
						cacheTag = getTagForCache(t, curMethodTable.MethodName, curTestCase.Input)
					}
					cachedData, ok = testApp.Appmemcache.GetCacheDataInCache(cacheTag)
					require.Equal(t, true, ok)
				}
				if !curTestCase.IsCacheTest {
					require.Equal(t, curTestCase.Error, err)
					require.Equal(t, curTestCase.Output, testRes)
				} else {
					checkCashLogic(t, testApp, &curMethodTable, &curTestCase, cachedData.InfoDTStamp)
				}
				testApp.RemoveDataInMemCacheBySOAPAction(curMethodTable.MethodName)
			})
		}
	}
}

func checkCashLogic(t *testing.T, testApp *app.App, methodTable *AppTestTable, testCase *AppTestCase, prevDataDTStamp time.Time) {
	t.Helper()
	var cacheTag string
	if !testCase.IsCacheData {
		time.Sleep(2 * time.Second)
	}
	rawBody, err := json.Marshal(testCase.Input)
	require.NoError(t, err)
	if methodTable.IsMethodWP {
		_, err := methodTable.MethodWP(testApp, context.Background())
		require.Equal(t, nil, err)
	} else {
		_, err := methodTable.Method(testApp, context.Background(), testCase.Input, string(rawBody))
		require.Equal(t, nil, err)
	}
	if methodTable.IsMethodWP {
		cacheTag = methodTable.MethodName
	} else {
		cacheTag = getTagForCache(t, methodTable.MethodName, testCase.Input)
	}
	cachedData2, ok := testApp.Appmemcache.GetCacheDataInCache(cacheTag)
	require.Equal(t, true, ok)
	if testCase.IsCacheData {
		require.Equal(t, prevDataDTStamp, cachedData2.InfoDTStamp)
	} else {
		require.NotEqual(t, prevDataDTStamp, cachedData2.InfoDTStamp)
	}
}
