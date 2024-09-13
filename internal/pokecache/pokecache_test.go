package pokecache

import (
	"testing"
	"time"
)

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	tests := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, tc := range tests {
		cache.Add(tc.inputKey, tc.inputVal)
		actual, ok := cache.Get(tc.inputKey)
		if !ok {
			t.Errorf("%s not found", tc.inputKey)
		}

		if string(actual) != string(tc.inputVal) {
			t.Errorf("%s doesn't match %s", string(actual), string(tc.inputVal))
		}
	}
}

func TestClean(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been cleaned", keyOne)
	}
}

func TestCleanFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been cleaned", keyOne)
	}
}
