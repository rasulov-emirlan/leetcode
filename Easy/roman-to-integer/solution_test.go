package romantointeger

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{
			in:  "III",
			out: 3,
		},
		{
			in:  "LVIII",
			out: 58,
		},
		{
			in:  "MCMXCIV",
			out: 1994,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := romanToInt(tC.in); got != tC.out {
				t.Errorf("wanted: %v, got: %v", tC.out, got)
			}
		})
	}

	romanToInt("MCMXCIV")
}
