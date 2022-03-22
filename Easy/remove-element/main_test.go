package removeelement

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		inA []int
		inB int
		out []int	
	}{
		{
			inA: []int{3,2,2,3},
			inB: 3,
			out: []int{2,2},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			_ = removeElement(tC.inA, tC.inB)
			for i := 0; i < len(tC.out); i++ {
				if tC.out[i] != tC.inA[i] {
					t.Errorf("wanted: %v, got: %v", tC.out, tC.inA)
				}
			}
		})
	}
}