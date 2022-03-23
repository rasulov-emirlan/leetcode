package removeduplicatesfromsortedarray

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		in []int
		outA int
		outB []int
		
	}{
		{
			in: []int{0,0,1,1,1,2,2,3,3,4},
			outB: []int{0,1,2,3,4},
		},
		{
			in: []int{1,1,2},
			outB: []int{1,2},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			removeDuplicates(tC.in)
			fmt.Println(tC.in)
		})
	}
}