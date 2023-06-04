package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := containsDuplicate(c.nums)
		if actual != c.expected {
			t.Errorf("expected: %v, actual: %v", c.expected, actual)
		}
	}
}
