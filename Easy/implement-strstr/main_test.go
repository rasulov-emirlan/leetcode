package implementstrstr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		inputA       string
		inputB       string
		wantedAnswer int
	}{
		{
			inputA:       "hello",
			inputB:       "ll",
			wantedAnswer: 2,
		},
		{
			inputA:       "aaaaa",
			inputB:       "bba",
			wantedAnswer: -1,
		},
		{
			inputA:       "",
			inputB:       "",
			wantedAnswer: 0,
		},
		{
			inputA:       "mississippi",
			inputB:       "issipi",
			wantedAnswer: -1,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := strStr(tC.inputA, tC.inputB)
			if result != tC.wantedAnswer {
				t.Errorf("wanted %d, but got %d", tC.wantedAnswer, result)
			}
		})
	}
}
