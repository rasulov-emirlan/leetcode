package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		digits []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, testCase := range testCases {
		actual := plusOne(testCase.digits)
		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("digits: %v, expect: %v, actual: %v", testCase.digits, testCase.expect, actual)
		}
	}
}
