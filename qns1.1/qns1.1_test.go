package qns1_1

import (
	"testing"
)

var params = map[string]bool{
	"abcdefg": true, "aaaabgcdefggg": false,
}

func TestIsUnique1(t *testing.T) {
	for param, expected := range params {
		actual := IsUnique1(param)
		if actual != expected {
			t.Fatalf(`Expected (%v) but got Actual (%v) for inputs (%v)`,
				expected, actual, param)
		}
	}
}

func TestIsUnique2(t *testing.T) {
	for param, expected := range params {
		actual := IsUnique2(param)
		if actual != expected {
			t.Fatalf(`Expected (%v) but got Actual (%v) for inputs (%v)`,
				expected, actual, param)
		}
	}
}
