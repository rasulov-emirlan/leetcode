package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	switch {
	case len(strs) == 0:
		return ""
	case len(strs[0]) == 0:
		return ""
	default:
	}
	prefix := []byte("")
	for i := 0; i < len(strs[0]); i++ {
		check := true
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				check = false
				break
			}
			if strs[j][i] != strs[0][i] {
				check = false
				break
			}
		}
		if check {
			prefix = append(prefix, strs[0][i])
		} else {
			break
		}
	}
	return string(prefix)
}
