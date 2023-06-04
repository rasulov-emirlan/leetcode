package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		if nums[i] == nums[j] {
			return true
		}
		i++
		j++
	}
	return false
}

// here is a faster solution, but looks like its more expensive in terms of memory

// func containsDuplicate(nums []int) bool {
// 	m := make(map[int]struct{})
// 	for _, v := range nums {
// 		if _, ok := m[v]; ok {
// 			return true
// 		}
// 		m[v] = struct{}{}
// 	}
// 	return false
// }
