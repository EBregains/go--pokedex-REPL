package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com/path",
			val: []byte("testData"),
		},
		{
			key: "https://example.com/",
			val: []byte("moreData"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			entrie, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key: %s", c.key)
				return
			}
			if string(entrie) != string(c.val) {
				t.Errorf("Expected to find value: %s", string(c.val))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com/path", []byte("testData"))

	_, ok := cache.Get("https://example.com/path")
	if !ok {
		t.Errorf("Expected to find key: %s", "https://example.com/path")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com/path")
	if ok {
		t.Errorf("Expected to not find key: %s", "https://example.com/path")
		return
	}

}
