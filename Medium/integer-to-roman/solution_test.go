package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, c := range cases {
		got := intToRoman(c.num)
		if got != c.want {
			t.Errorf("intToRoman(%d) == %q, want %q", c.num, got, c.want)
		}
	}
}
