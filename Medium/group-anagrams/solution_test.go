package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnograms(t *testing.T) {
	cases := []struct {
		in   []string
		want [][]string
	}{
		{
			in:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
	}

	for _, c := range cases {
		got := groupAnagrams(c.in)
		// sort the result for comparison
		for _, v := range got {
			sort.Strings(v)
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("groupAnagrams(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
