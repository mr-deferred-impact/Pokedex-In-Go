package pokecache

import (
	"slices"
	"testing"
	"time"
)

func TestAddGetCache(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{key: "Test1", val: []byte("Testing1_Value")},
		{key: "Test2", val: []byte("Testing2_Value")},
		// Overwriting an existing key should return the new value.
		{key: "Test1", val: []byte("Testing5_Value")},
	}

	testCache := NewCache(interval)

	for _, tc := range cases {
		testCache.Add(tc.key, tc.val)

		val, ok := testCache.Get(tc.key)
		if !ok {
			t.Errorf("expected key %q to be found", tc.key)
			continue
		}
		if !slices.Equal(val, tc.val) {
			t.Errorf("key %q: expected %s, got %s", tc.key, tc.val, val)
		}
	}
}

func TestGetMissingKey(t *testing.T) {
	testCache := NewCache(5 * time.Second)

	if _, ok := testCache.Get("does-not-exist"); ok {
		t.Errorf("expected missing key to report ok == false")
	}
}

func TestFreshEntryStays(t *testing.T) {
	// Long interval, so no tick fires during the test.
	const interval = 200 * time.Millisecond

	testCache := NewCache(interval)
	testCache.Add("fresh", []byte("data"))

	time.Sleep(interval / 4)

	if _, ok := testCache.Get("fresh"); !ok {
		t.Errorf("expected fresh entry to still be in the cache")
	}
}

func TestReapLoopRemovesOldEntries(t *testing.T) {
	const interval = 20 * time.Millisecond

	testCache := NewCache(interval)
	testCache.Add("old", []byte("data"))

	// Wait several intervals so at least one tick runs after the entry expires.
	time.Sleep(5 * interval)

	if _, ok := testCache.Get("old"); ok {
		t.Errorf("expected old entry to be reaped")
	}
}
