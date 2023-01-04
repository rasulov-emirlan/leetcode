package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		inputA string
		inputB string
		output string
	}{
		{
			inputA: "11",
			inputB: "1",
			output: "100",
		},
		{
			inputA: "1010",
			inputB: "1011",
			output: "10101",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.inputA+"+"+tC.inputB, func(t *testing.T) {
			res := addBinary(tC.inputA, tC.inputB)
			if res != tC.output {
				t.Errorf("expected %q, got %q", tC.output, res)
			}
		})
	}
}
