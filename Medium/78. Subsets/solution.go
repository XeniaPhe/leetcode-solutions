package subsets

import "slices"

func subsets(nums []int) [][]int {
    subsets := make([][]int, 1 << len(nums))
    calculateSubsets(subsets, nums, make([]int, len(nums)), 0, 0)
    return subsets
}

func calculateSubsets(subsets [][]int, nums []int, curr []int, subsetIdx int, depth int) int {
    subsets[subsetIdx] = slices.Clone(curr[:depth])
    subsetIdx = subsetIdx + 1

    for i, nextDepth := 0, depth + 1; i < len(nums); i += 1 {
        curr[depth] = nums[i]
        subsetIdx = calculateSubsets(subsets, nums[i + 1:], curr, subsetIdx, nextDepth)
    }

    return subsetIdx
}