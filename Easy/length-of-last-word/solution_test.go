package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, tc := range testCases {
		actual := lengthOfLastWord(tc.input)
		if actual != tc.expected {
			t.Errorf("lengthOfLastWord(%q) = %d; expected %d", tc.input, actual, tc.expected)
		}
	}
}
