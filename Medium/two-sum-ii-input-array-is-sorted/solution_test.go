package twosumiiinputarrayissorted

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
	}

	for _, c := range cases {
		got := twoSum(c.numbers, c.target)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"\n input: %v, %d \n got: %v \n want: %v",
				c.numbers,
				c.target,
				got,
				c.want,
			)
		}
	}
}
