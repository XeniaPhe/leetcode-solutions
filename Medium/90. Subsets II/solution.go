package subsets_II

import "slices"

func subsetsWithDup(nums []int) [][]int {
    subsets := make([][]int, 1 << len(nums))
    slices.Sort(nums)
    return subsets[:calculateSubsets(subsets, nums, make([]int, len(nums)), 0, 0)]
}

func calculateSubsets(subsets [][]int, nums []int, curr []int, idx int, depth int) int {
    if subsets[idx], idx = slices.Clone(curr[:depth]), idx + 1; len(nums) == 0 {
        return idx
    }

    for i, nextDepth, prev := 0, depth + 1, nums[0] ^ 1; i < len(nums); i += 1 {
        if prev != nums[i] {
            prev, curr[depth] = nums[i], nums[i]
            idx = calculateSubsets(subsets, nums[i + 1:], curr, idx, nextDepth)
        }
    }

    return idx
}