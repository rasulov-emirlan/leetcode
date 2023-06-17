package dailytemperatures

import "testing"

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		temperatures []int
		want         []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for _, c := range cases {
		got := dailyTemperatures(c.temperatures)
		if len(got) != len(c.want) {
			t.Errorf("dailyTemperatures(%v) == %v, want %v", c.temperatures, got, c.want)
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("dailyTemperatures(%v) == %v, want %v", c.temperatures, got, c.want)
				break
			}
		}
	}
}
