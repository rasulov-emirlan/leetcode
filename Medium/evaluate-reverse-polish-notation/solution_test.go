package evaluatereversepolishnotation

import "testing"

func TestEvalRPN(t *testing.T) {
	cases := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
	}

	for _, c := range cases {
		got := evalRPN(c.tokens)
		if got != c.want {
			t.Errorf("evalRPN(%v) == %d, want %d", c.tokens, got, c.want)
		}
	}
}
