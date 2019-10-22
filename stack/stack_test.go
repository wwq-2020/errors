package stack

import (
	"testing"
)

func getTest1() string {
	return Get(0)
}

func getTest2() string {
	return Get(0)
}

func getTest3() string {
	return Get(0)
}

func TestGet(t *testing.T) {
	tests := []struct {
		fn       func() string
		expected string
	}{
		{fn: getTest1, expected: "stack.getTest1:8"},
		{fn: getTest2, expected: "stack.getTest2:12"},
		{fn: getTest3, expected: "stack.getTest3:16"},
	}
	for _, test := range tests {
		got := test.fn()
		if got != test.expected {
			t.Fatalf("expected:%+v,got:%+v", test.expected, got)
		}
	}
}
