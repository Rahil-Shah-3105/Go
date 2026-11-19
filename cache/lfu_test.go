package cache_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/cache"
)

func TestLFU(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestLFU(t=%v)", t)
	lfuCache := cache.NewLFU(3)
	t.Run("Test 1: Put number and Get is correct", func(t *testing.T) {
		key, value := "1", 1
		lfuCache.Put(key, value)
		got := lfuCache.Get(key)

		if got != value {
			t.Errorf("expected: %v, got: %v", value, got)
		}
	})

	t.Run("Test 2: Get data not stored in cache should return nil", func(t *testing.T) {
		got := lfuCache.Get("2")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}
	})

	t.Run("Test 3: Put data over capacity and Get should return nil", func(t *testing.T) {
		lfuCache.Put("2", 2)
		lfuCache.Put("3", 3)
		lfuCache.Put("4", 4)

		got := lfuCache.Get("1")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}
	})

	t.Run("test 4: Put key over capacity but recent key exists", func(t *testing.T) {
		lfuCache.Put("4", 4)
		lfuCache.Put("3", 3)
		lfuCache.Put("2", 2)
		lfuCache.Put("1", 1)

		got := lfuCache.Get("4")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}

		expected := 3
		got = lfuCache.Get("3")

		logritelog.Default().Printf("[DEBUG] TestLFU(t=%v): got → %v", t, got)
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}

	})
	logritelog.Default().Printf("[WARN] <<< Exiting TestLFU(t=%v)", t)
}
