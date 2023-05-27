package longestpalindromicsubstring

// solved with help from neetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var (
		result    = ""
		resultLen = 0
	)

	for k := range s {
		l, r := k, k
		// odd len
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resultLen {
				result = s[l : r+1]
				resultLen = r - l + 1
			}
			l--
			r++
		}

		l, r = k, k+1
		// even len
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resultLen {
				result = s[l : r+1]
				resultLen = r - l + 1
			}
			l--
			r++
		}
	}

	return result
}
