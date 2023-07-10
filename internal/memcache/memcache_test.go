package memcache_test

import (
	"testing"

	memcache "github.com/skolzkyi/cbrwsdltojson/internal/memcache"
	"github.com/stretchr/testify/require"
)

func TestAllMemCacheCases(t *testing.T) {
	memcacheExempl := memcache.New()
	memcacheExempl.Init()
	t.Parallel()
	t.Run("Test_AddOrUpdatePayloadInCache_And_GetPayloadInCache", func(t *testing.T) {
		t.Parallel()
		ok := memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCache")
		require.Equal(t, false, ok)
		testPayload, ok := memcacheExempl.GetPayloadInCache("testTag_TestAddOrUpdatePayloadInCache")
		require.Equal(t, true, ok)
		testPayloadStr, ok := testPayload.(string)
		require.Equal(t, true, ok)
		require.Equal(t, testPayloadStr, "testPayload_TestAddOrUpdatePayloadInCache")
		ok = memcacheExempl.AddOrUpdatePayloadInCache("testTag_TestAddOrUpdatePayloadInCache", "testPayload_TestAddOrUpdatePayloadInCacheUpd")
		require.Equal(t, true, ok)
	})
	t.Run("Test_AddOrUpdatePayloadInCache_And_RemovePayloadInCache", func(t *testing.T) {
		t.Parallel()
		memcacheExempl.AddOrUpdatePayloadInCache("testTag_RemovePayloadInCache", "testPayload_RemovePayloadInCache")
		_, ok := memcacheExempl.GetPayloadInCache("testTag_RemovePayloadInCache")
		require.Equal(t, true, ok)
		memcacheExempl.RemovePayloadInCache("testTag_RemovePayloadInCache")
		_, ok = memcacheExempl.GetPayloadInCache("testTag_RemovePayloadInCache")
		require.Equal(t, false, ok)
	})
}
