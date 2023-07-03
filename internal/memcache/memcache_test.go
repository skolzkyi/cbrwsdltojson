package memcache_test

import (
	"testing"

	memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
	"github.com/stretchr/testify/require"
)

func TestAllMemCacheCases(t *testing.T) {
	memcacheExempl := memcache.New()
	memcacheExempl.Init()
	t.Run("Test_AddOrUpdatePayloadInCache_And_GetPayloadInCache", func(t *testing.T) {
		t.Parallel()
		ok := memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCache")
		require.Equal(t, ok, false)
		testPayload, ok := memcacheExempl.GetPayloadInCache("testTag_TestAddOrUpdatePayloadInCache")
		require.Equal(t, ok, true)
		testPayloadStr, ok := testPayload.(string)
		require.Equal(t, ok, true)
		require.Equal(t, testPayloadStr, "testPayload_TestAddOrUpdatePayloadInCache")
		ok = memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCacheUpd")
		require.Equal(t, ok, true)
	})
	t.Run("Test_AddOrUpdatePayloadInCache_And_RemovePayloadInCache", func(t *testing.T) {
		t.Parallel()
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemovePayloadInCache", "testPayload_RemovePayloadInCache")
		_, ok := memcacheExempl.GetPayloadInCache("testTag_RemovePayloadInCache")
		require.Equal(t, ok, true)
		memcacheExempl.RemovePayloadInCache("testTag_RemovePayloadInCache")
		_, ok = memcacheExempl.GetPayloadInCache("testTag_RemovePayloadInCache")
		require.Equal(t, ok, false)
	})

}
