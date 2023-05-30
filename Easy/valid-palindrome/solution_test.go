package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"race a car", false},
	}

	for _, c := range cases {
		got := isPalindrome(c.s)
		if got != c.want {
			t.Errorf("isPalindrome(%q) == %v, want %v", c.s, got, c.want)
		}
	}
}
