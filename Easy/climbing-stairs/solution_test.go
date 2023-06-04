package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, c := range cases {
		actual := climbStairs(c.n)
		if actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}
