package permutations_II

import "slices"

func permuteUnique(nums []int) [][]int {
    permCount := 1
    for i := len(nums); i >= 2; i -= 1 {
        permCount *= i
    }

    slices.Sort(nums)
    curr := make([]int, len(nums))
    indices := make([]bool, len(nums))
    perms := make([][]int, 0, permCount)
    return perm(perms, nums, indices, curr, 0)
}

func perm(perms [][]int, nums []int, indices []bool, curr []int, idx int) [][]int {
    if len(nums) == idx {
        return append(perms, slices.Clone(curr))
    }

    for prev, i := nums[0] - 1, 0; i < len(indices); i += 1 {
        if !indices[i] && nums[i] != prev {
            curr[idx], prev = nums[i], nums[i]
            indices[i] = true
            perms = perm(perms, nums, indices, curr, idx + 1)
            indices[i] = false
        }
    }

    return perms
}