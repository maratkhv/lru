package lru

import (
	"testing"
)

type tcase struct {
	key   any
	value any
	name  string
}

var testCases = []tcase{
	tcase{
		key:   "first",
		value: 1,
		name:  "1: string to int",
	},
	tcase{
		key:   2,
		value: "second",
		name:  "2: int to string",
	},
	tcase{
		key: struct {
			i int
			s string
			f float32
		}{1, "one", 1.337},
		value: nil,
		name:  "3: srtuct to nil",
	},
}

func TestLru(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cache := NewAuto(5, tc.key, tc.value)
			if cache == nil {
				t.Error("returned cache is nil")
			}

			cache.Put(tc.key, tc.value)
			if got := cache.Get(tc.key); got != tc.value {
				t.Errorf("want: %v, got: %v", tc.value, got)
			}
		})
	}
}
