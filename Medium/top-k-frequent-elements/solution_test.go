package topkfrequentelements

import "testing"

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}

	for _, c := range cases {
		got := topKFrequent(c.nums, c.k)
		if len(got) != len(c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
