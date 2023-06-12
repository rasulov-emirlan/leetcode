package validsudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	cases := []struct {
		board [][]byte
		want  bool
	}{
		{
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: true,
		},
	}

	for _, c := range cases {
		got := isValidSudoku(c.board)

		if got != c.want {
			t.Errorf("isValidSudoku(%v) == %t, want %t", c.board, got, c.want)
		}
	}
}
