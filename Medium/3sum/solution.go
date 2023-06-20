package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	for idx := 0; idx < len(nums)-1; idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}
		v := nums[idx]

		i := idx + 1
		j := len(nums) - 1

		for i < j {
			sum := v + nums[i] + nums[j]

			if sum == 0 {
				result = append(result, []int{v, nums[i], nums[j]})
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if sum > 0 {
				j--
			} else {
				i++
			}
		}
	}

	return result
}
