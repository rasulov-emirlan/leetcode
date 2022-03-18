package longestcommonprefix

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		input        []string
		wantedAnswer string
	}{
		{
			input:        []string{"flower", "flow", "flight"},
			wantedAnswer: "fl",
		},
		{
			input:        []string{"dog", "racecar", "car"},
			wantedAnswer: "",
		},
		{
			input:        []string{"ab", "a"},
			wantedAnswer: "a",
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			answer := longestCommonPrefix(tC.input)
			if answer != tC.wantedAnswer {
				t.Errorf("wanted to see: %s, but got: %s", tC.wantedAnswer, answer)
			}
		})
	}
}
