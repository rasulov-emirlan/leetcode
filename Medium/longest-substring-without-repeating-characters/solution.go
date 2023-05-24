package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	max := 0
	curr := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := curr[s[i]]; ok {
			i -= len(curr) - 1
			curr = make(map[byte]struct{})
		}
		curr[s[i]] = struct{}{}
		if max < len(curr) {
			max = len(curr)
		}
	}

	return max
}
