package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums) == 2 && nums[0] == nums[1] {
		return 1
	}
    m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] >= 1 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		if m[nums[i]] == 0 {
			m[nums[i]]++
		}
	}
	return len(nums)
}