package combinationsum

// copilot solved it
// TODO: learn more about this stuff

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	var dfs func([]int, int, int) [][]int // dfs is a recursive function that returns all possible combinations to reach target

	dfs = func(combination []int, sum int, index int) [][]int {
		if sum == target {
			res = append(res, combination)
			return res
		}

		if sum > target {
			return res
		}

		for i := index; i < len(candidates); i++ {
			res = dfs(append(append([]int{}, combination...), candidates[i]), sum+candidates[i], i)
		}

		return res
	}

	return dfs([]int{}, 0, 0)
}
