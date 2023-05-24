package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"aau", 2},
		{"dvdf", 3},
	}

	for _, c := range cases {
		got := lengthOfLongestSubstring(c.in)
		if got != c.want {
			t.Errorf("lengthOfLongestSubstring(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
