package house_robber_II

func rob(nums []int) int {
	last, matNoRob, matRob := len(nums) - 1, [2][2]int{}, [2][2]int {{0, nums[0]}}
	dp := func(mat [2][2]int, curr int) [2][2]int {
		return [2][2]int {{max(mat[0][1], mat[1][1]), curr + max(mat[0][0], mat[1][0])}, {mat[0][0], mat[0][1]}}
	}

	for i := 1; i < last; matNoRob, matRob, i = dp(matNoRob, nums[i]), dp(matRob, nums[i]), i + 1 { }
	return max(dp(matNoRob, nums[last])[0][0], dp(matNoRob, nums[last])[0][1], matRob[0][0], matRob[0][1])
}