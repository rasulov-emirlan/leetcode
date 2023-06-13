package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	// fill the set
	for _, v := range nums {
		set[v] = struct{}{}
	}

	maxSeqLen := 0

	for _, v := range nums {
		_, ok := set[v-1]

		if !ok {
			v++
			seqLen := 1
			_, ok = set[v]
			for ok {
				_, ok = set[v]
				if ok {
					seqLen++
				}
				v++
			}
			if seqLen > maxSeqLen {
				maxSeqLen = seqLen
			}
			continue
		}
	}

	return maxSeqLen
}
