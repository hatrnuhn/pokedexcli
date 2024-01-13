package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey   string
		inputVal   []byte
		expectedOk bool
	}{
		{
			inputKey:   "key1",
			inputVal:   []byte("val1"),
			expectedOk: true,
		}, {
			inputKey:   "key2",
			inputVal:   []byte("val2"),
			expectedOk: true,
		},
	}

	for _, cs := range cases {

		cache.Add(cs.inputKey, cs.inputVal)

		actual, ok := cache.Get(cs.inputKey)

		if !ok {
			t.Error("could not find ", cs.inputKey)
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("value does not match: %v actual vs %v case", actual, cs.inputVal)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	// add 2ms so reap will loop
	time.Sleep(interval + time.Millisecond*2)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s shoulda been reaped", keyOne)
	}
}
