package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s      string
		t      string
		expect bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
		{"ab", "a", false},
		{"a", "a", true},
		{"a", "b", false},
		{"", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.s+"+"+tc.t, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.expect {
				t.Errorf("got %v, expect %v", got, tc.expect)
			}
		})
	}
}
