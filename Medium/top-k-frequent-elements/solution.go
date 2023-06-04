package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		freq[v] = append(freq[v], k)
	}

	res := make([]int, 0)
	for i := len(freq) - 1; i >= 0; i-- {
		for j := len(freq[i]) - 1; j >= 0; j-- {
			res = append(res, freq[i][j])
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
