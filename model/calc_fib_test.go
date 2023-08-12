package model

import (
	"fmt"
	"testing"
)

func TestCalcFib(t *testing.T) {
	test := []struct {
		num      int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{5, "5"},
		{10, "55"},
		{99, "218922995834555169026"},
		{-1, "0"},
	}

	for _, tc := range test {
		t.Run(fmt.Sprintf("num=%d", tc.num), func(t *testing.T) {
			result := CalcFib(tc.num).String()
			if result != tc.expected {
				t.Errorf("CalcFib(%d) = %s; expected = %s", tc.num, result, tc.expected)
			}
		})
	}
}
