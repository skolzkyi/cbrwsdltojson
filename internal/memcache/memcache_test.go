package memcache_test

import (
	"testing"
	"time"

	memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
	"github.com/stretchr/testify/require"
)

func TestAllMemCacheCases(t *testing.T) {
	memcacheExempl := memcache.New()
	memcacheExempl.Init()
	t.Run("Test_AddOrUpdatePayloadInCache_And_GetPayloadInCache", func(t *testing.T) {
		ok := memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCache")
		require.Equal(t, false, ok)
		testCacheData, ok := memcacheExempl.GetCacheDataInCache("testTag_TestAddOrUpdatePayloadInCache")
		require.Equal(t, true, ok)
		testPayloadStr, ok := testCacheData.Payload.(string)
		require.Equal(t, true, ok)
		require.Equal(t, testPayloadStr, "testPayload_TestAddOrUpdatePayloadInCache")
		ok = memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCacheUpd")
		require.Equal(t, true, ok)
		testCacheData2, ok := memcacheExempl.GetCacheDataInCache("testTag_TestAddOrUpdatePayloadInCache")
		require.Equal(t, true, ok)
		require.NotEqual(t, testCacheData.InfoDTStamp, testCacheData2.InfoDTStamp)
	})
	t.Run("Test_AddOrUpdatePayloadInCache_And_RemovePayloadInCache", func(t *testing.T) {
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemovePayloadInCache", "testPayload_RemovePayloadInCache")
		_, ok := memcacheExempl.GetCacheDataInCache("testTag_RemovePayloadInCache")
		require.Equal(t, true, ok)
		memcacheExempl.RemovePayloadInCache("testTag_RemovePayloadInCache")
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemovePayloadInCache")
		require.Equal(t, false, ok)
	})

	t.Run("Test_RemoveAllPayloadInCacheByTimeStamp", func(t *testing.T) {
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_1", "testPayload_RemoveAllPayloadInCacheByTimeStamp_1")
		_, ok := memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_1")
		require.Equal(t, true, ok)
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_2", "testPayload_RemoveAllPayloadInCacheByTimeStamp_2")
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_2")
		require.Equal(t, true, ok)
		time.Sleep(time.Millisecond)
		testStartTime := time.Now()
		time.Sleep(time.Millisecond)
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_3", "testPayload_RemoveAllPayloadInCacheByTimeStamp_3")
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_3")
		require.Equal(t, true, ok)
		memcacheExempl.RemoveAllPayloadInCacheByTimeStamp(testStartTime)
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_1")
		require.Equal(t, false, ok)
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_1")
		require.Equal(t, false, ok)
		_, ok = memcacheExempl.GetCacheDataInCache("testTag_RemoveAllPayloadInCacheByTimeStamp_3")
		require.Equal(t, true, ok)
	})
}
