package floodfill

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	testCases := []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			image: [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
			sr:    1,
			sc:    1,
			color: 1,
			want: [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.image), func(t *testing.T) {
			got := floodFill(tc.image, tc.sr, tc.sc, tc.color)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
