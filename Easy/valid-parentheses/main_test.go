package validparentheses

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		input string
		output bool
		
	}{
		{
			input: "[",
			output: false,
		},
		{
			input: "()[]{}",
			output: true,
		},
		{
			input: "()",
			output: true,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if res := isValid(tC.input); res != tC.output {
				t.Errorf("wanted: %v, got: %v", tC.output, res)
			}
		})
	}
}