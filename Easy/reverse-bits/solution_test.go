package reversebits

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		n        uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, c := range cases {
		actual := reverseBits(c.n)
		if actual != c.expected {
			t.Errorf("Input: %d. Expected: %d. Actual: %d\n", c.n, c.expected, actual)
		}
	}
}
