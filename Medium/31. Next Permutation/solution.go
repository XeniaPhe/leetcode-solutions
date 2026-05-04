package next_permutation

import "slices"

func nextPermutation(nums []int)  {
    if len(nums) == 1 {
        return
    }

    for min, i := 0, len(nums) - 1; i > 0; i -= 1 {
        if min = i; nums[i - 1] >= nums[i] {
            continue
        }

        for j := i + 1; j < len(nums); j += 1 {
            if nums[j] > nums[i - 1] && nums[j] <= nums[min] {
                min = j
            }
        }

        nums[i - 1], nums[min] = nums[min], nums[i - 1]
        slices.Reverse(nums[i:])
        return
    }

    slices.Reverse(nums)
}