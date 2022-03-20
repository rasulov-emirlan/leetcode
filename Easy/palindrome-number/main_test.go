package palindromenumber

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		input  int
		output bool
	}{
		{
			input:  121,
			output: true,
		},
		{
			input:  10,
			output: false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := isPalindrome(tC.input)
			if result != tC.output {
				t.Errorf("wated to see: %v, but got: %v", tC.output, result)
			}
		})
	}
}
