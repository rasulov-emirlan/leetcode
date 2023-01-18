package besttimetobuyandsellstock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.prices), func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
