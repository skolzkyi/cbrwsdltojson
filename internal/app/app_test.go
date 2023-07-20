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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.GetCursOnDateXML{
		OnDate: "2023-06-22",
	}, testGetCursOnDateXMLResult)
	testDataGetCursOnDate.TestCases = append(testDataGetCursOnDate.TestCases, standartTestCacheCases...)
	testDataGetCursOnDate.TestCases = testCases
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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.BiCurBaseXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
	}, testBiCurBaseXMLResult)
	testDataBiCurBaseXML.TestCases = append(testDataBiCurBaseXML.TestCases, standartTestCacheCases...)
	testDataBiCurBaseXML.TestCases = testCases
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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.BliquidityXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
	}, testBliquidityXMLResult)
	testDataBliquidityXML.TestCases = append(testDataBliquidityXML.TestCases, standartTestCacheCases...)
	testDataBliquidityXML.TestCases = testCases
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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.DepoDynamicXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
	}, testDepoDynamicXMLResult)
	testDataDepoDynamicXML.TestCases = append(testDataDepoDynamicXML.TestCases, standartTestCacheCases...)
	testDataDepoDynamicXML.TestCases = testCases
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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.DragMetDynamicXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
	}, testDragMetDynamicXMLResult)
	testDataDragMetDynamicXML.TestCases = append(testDataDragMetDynamicXML.TestCases, standartTestCacheCases...)
	testDataDragMetDynamicXML.TestCases = testCases
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
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.DVXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
	}, testDVXMLResult)
	testDataDVXML.TestCases = append(testDataDVXML.TestCases, standartTestCacheCases...)
	testDataDVXML.TestCases = testCases
	return testDataDVXML
}

func TestAllAppCases(t *testing.T) {
	acTable := AllCasesTable{}
	acTable.CasesByMethod = make([]AppTestTable, 6)
	acTable.CasesByMethod[0] = initTestDataGetCursOnDateXML(t)
	acTable.CasesByMethod[1] = initTestDataBiCurBaseXML(t)
	acTable.CasesByMethod[2] = initTestDataBliquidityXML(t)
	acTable.CasesByMethod[3] = initTestDataDepoDynamicXML(t)
	acTable.CasesByMethod[4] = initTestDragMetDynamicXML(t)
	acTable.CasesByMethod[5] = initTestDataDVXML(t)
	t.Parallel()
	for _, curMethodTable := range acTable.CasesByMethod {
		curMethodTable := curMethodTable
		for _, curTestCase := range curMethodTable.TestCases {
			curTestCase := curTestCase
			t.Run(curMethodTable.MethodName+":"+curTestCase.Name, func(t *testing.T) {
				t.Parallel()
				var cachedData memcache.CacheInfo
				var ok bool
				testApp := initTestApp(t)
				rawBody, err := json.Marshal(curTestCase.Input)
				require.NoError(t, err)
				testRes, err := curMethodTable.Method(testApp, context.Background(), curTestCase.Input, string(rawBody))
				if err == nil {
					// testApp.Appmemcache.PrintAllCacheKeys()
					cacheTag := getTagForCache(t, curMethodTable.MethodName, curTestCase.Input)
					cachedData, ok = testApp.Appmemcache.GetCacheDataInCache(cacheTag)
					require.Equal(t, true, ok)
				}
				if !curTestCase.IsCacheTest {
					require.Equal(t, curTestCase.Error, err)
					require.Equal(t, curTestCase.Output, testRes)
				} else {
					if !curTestCase.IsCacheData {
						time.Sleep(3 * time.Second)
					}
					_, err := curMethodTable.Method(testApp, context.Background(), curTestCase.Input, string(rawBody))
					require.Equal(t, nil, err)
					cachedData2, ok := testApp.Appmemcache.GetCacheDataInCache(curMethodTable.MethodName)
					require.Equal(t, true, ok)
					if curTestCase.IsCacheData {
						require.Equal(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
					} else {
						require.NotEqual(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
					}
				}
				testApp.RemoveDataInMemCacheBySOAPAction(curMethodTable.MethodName)
			})
		}
	}
}
