package groupanagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)

	for _, v := range strs {
		words[sortString(v)] = append(words[sortString(v)], v)
	}

	res := make([][]string, 0, len(words))

	for _, v := range words {
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	return strings.Join(sl, "")
}
