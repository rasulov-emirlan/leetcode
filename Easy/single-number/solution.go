package singlenumber

// solution from neetcode

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}

	return res
}
