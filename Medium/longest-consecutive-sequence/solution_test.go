package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
	}

	for _, c := range cases {
		got := longestConsecutive(c.nums)

		if got != c.want {
			t.Errorf("longestConsecutive(%v) == %d, want %d", c.nums, got, c.want)
		}
	}
}
