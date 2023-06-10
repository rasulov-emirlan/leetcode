package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))

	prefix := 1
	for _, v := range nums {
		res = append(res, prefix)
		prefix *= v
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
