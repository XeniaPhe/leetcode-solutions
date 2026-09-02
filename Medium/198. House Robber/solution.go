package house_robber

func rob(nums []int) int {
    var mat [2][2]int
    for i := 0; i < len(nums); i += 1 {
        noRob, rob := max(mat[0][1], mat[1][1]), nums[i] + max(mat[0][0], mat[1][0])
        mat[0][0], mat[0][1], mat[1][0], mat[1][1] = noRob, rob, mat[0][0], mat[0][1]
    }

    return max(mat[0][0], mat[0][1])
}