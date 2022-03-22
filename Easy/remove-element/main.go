package removeelement

func removeElement(nums []int, val int) int {
	index := 0
	for i := range nums {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}