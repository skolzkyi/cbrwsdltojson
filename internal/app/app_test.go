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
	field1 string
	field2 string
	field3 int
}

type AppTestTable struct {
	MethodName string
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
		field1: "abc",
		field2: "def",
		field3: 0,
	}
	testStruct2 := testStruct{
		field1: "123",
		field2: "456",
		field3: 1,
	}
	err := testApp.AddOrUpdateDataInCache("ts1", testStruct1, testStruct1.field3)
	require.NoError(t, err)
	err = testApp.AddOrUpdateDataInCache("ts2", testStruct2, testStruct2.field3)
	require.NoError(t, err)
	payload1, ok, err := testApp.GetDataInCacheIfExisting("ts1", testStruct1)
	require.Equal(t, true, ok)
	require.NoError(t, err)
	data1, ok := payload1.(int)
	require.Equal(t, true, ok)
	require.Equal(t, testStruct1.field3, data1)
	payload2, ok, err := testApp.GetDataInCacheIfExisting("ts2", testStruct2)
	require.Equal(t, true, ok)
	require.NoError(t, err)
	data2, ok := payload2.(int)
	require.Equal(t, true, ok)
	require.Equal(t, testStruct2.field3, data2)
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

func getTagForCache(t *testing.T, SOAPMethod string, request interface{}) string {
	t.Helper()
	jsonstring, err := json.Marshal(request)
	require.NoError(t, err)
	return SOAPMethod + string(jsonstring)
}

// GetCursOnDate.
func initTestDataGetCursOnDateXML(t *testing.T) *AppTestTable {
	t.Helper()
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
	testDataGetCursOnDate := AppTestTable{
		MethodName: "GetCursOnDateXML",
	}
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: datastructures.GetCursOnDateXML{
			OnDate: "2023-06-22",
		},
		Output: testGetCursOnDateXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: datastructures.GetCursOnDateXML{
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
	return &testDataGetCursOnDate
}

// BiCurBaseXML.
func initTestDataBiCurBaseXML(t *testing.T) *AppTestTable {
	t.Helper()
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
	testDataBiCurBaseXML := AppTestTable{
		MethodName: "BiCurBaseXML",
	}
	testCases := make([]AppTestCase, 2)
	testCases[0] = AppTestCase{
		Name: "Positive",
		Input: datastructures.BiCurBaseXML{
			FromDate: "2023-06-22",
			ToDate:   "2023-06-23",
			//	XMLNs:    "http://web.cbr.ru/",
		},
		Output: testBiCurBaseXMLResult,
		Error:  nil,
	}

	testCases[1] = AppTestCase{
		Name: "Negative",
		Input: datastructures.BiCurBaseXML{
			FromDate: "022-14-22",
			ToDate:   "2023-06-23",
			//	XMLNs:    "http://web.cbr.ru/",
		},
		Output: datastructures.BiCurBaseXMLResult{},
		Error:  customsoap.ErrContextWSReqExpired,
	}
	standartTestCacheCases := createStandartTestCacheCases(t, datastructures.BiCurBaseXML{
		FromDate: "022-14-22",
		ToDate:   "2023-06-23",
		//XMLNs:    "http://web.cbr.ru/",
	}, testBiCurBaseXMLResult)
	testDataBiCurBaseXML.TestCases = append(testDataBiCurBaseXML.TestCases, standartTestCacheCases...)
	testDataBiCurBaseXML.TestCases = testCases
	return &testDataBiCurBaseXML
}

func TestCasesGetCursOnDateXML(t *testing.T) {
	t.Parallel()
	testCasesByMethod := initTestDataGetCursOnDateXML(t)
	for _, curTestCase := range testCasesByMethod.TestCases {
		curTestCase := curTestCase
		t.Run(testCasesByMethod.MethodName+":"+curTestCase.Name, func(t *testing.T) {
			t.Parallel()
			var cachedData memcache.CacheInfo
			testApp := initTestApp(t)
			inputAssert, ok := curTestCase.Input.(datastructures.GetCursOnDateXML)
			require.Equal(t, true, ok)
			testRes, err := testApp.GetCursOnDateXML(context.Background(), inputAssert)
			if err == nil {
				testApp.Appmemcache.PrintAllCacheKeys()
				cacheTag := getTagForCache(t, testCasesByMethod.MethodName, inputAssert)
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
				_, err := testApp.GetCursOnDateXML(context.Background(), inputAssert)
				require.Equal(t, nil, err)
				cachedData2, ok := testApp.Appmemcache.GetCacheDataInCache(testCasesByMethod.MethodName)
				require.Equal(t, true, ok)
				if curTestCase.IsCacheData {
					require.Equal(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
				} else {
					require.NotEqual(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
				}
			}
			testApp.RemoveDataInMemCacheBySOAPAction(testCasesByMethod.MethodName)
		})
	}
}

func TestCasesBiCurBaseXML(t *testing.T) {
	t.Parallel()
	testCasesByMethod := initTestDataBiCurBaseXML(t)
	for _, curTestCase := range testCasesByMethod.TestCases {
		curTestCase := curTestCase
		t.Run(testCasesByMethod.MethodName+":"+curTestCase.Name, func(t *testing.T) {
			t.Parallel()
			var cachedData memcache.CacheInfo
			testApp := initTestApp(t)
			inputAssert, ok := curTestCase.Input.(datastructures.BiCurBaseXML)
			require.Equal(t, true, ok)
			testRes, err := testApp.BiCurBaseXML(context.Background(), inputAssert)
			if err == nil {
				testApp.Appmemcache.PrintAllCacheKeys()
				cacheTag := getTagForCache(t, testCasesByMethod.MethodName, inputAssert)
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
				_, err := testApp.BiCurBaseXML(context.Background(), inputAssert)
				require.Equal(t, nil, err)
				cachedData2, ok := testApp.Appmemcache.GetCacheDataInCache(testCasesByMethod.MethodName)
				require.Equal(t, true, ok)
				if curTestCase.IsCacheData {
					require.Equal(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
				} else {
					require.NotEqual(t, cachedData.InfoDTStamp, cachedData2.InfoDTStamp)
				}
			}
			testApp.RemoveDataInMemCacheBySOAPAction(testCasesByMethod.MethodName)
		})
	}
}
