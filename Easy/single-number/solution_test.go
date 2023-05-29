package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, c := range cases {
		got := singleNumber(c.nums)

		if got != c.want {
			t.Errorf("singleNumber(%v) == %d, want %d", c.nums, got, c.want)
		}
	}
}
