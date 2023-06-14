package generateparentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		n        int
		expected []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, c := range cases {
		res := generateParenthesis(c.n)
		if len(res) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, res)
		}
	}
}
