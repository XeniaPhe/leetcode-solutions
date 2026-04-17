package three_sum

import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    result := make([][]int, 0)

    for num1, i := nums[0] - 1, 0; i < len(nums) - 2; i += 1 {
        if nums[i] == num1 {
            continue
        }

        num1 = nums[i]
        target := -num1
        j, k := i + 1, len(nums) - 1

        if nums[j] + nums[j + 1] > target {
            break
        }
        
        if nums[k] + nums[k - 1] < target {
            continue
        }

        for j < k {
            sum := nums[j] + nums[k]
            if sum == target {
                result = append(result, []int {num1, nums[j], nums[k]})
                for j += 1; j < k && nums[j] == nums[j - 1]; j += 1 { }
                for k -= 1; j < k && nums[k] == nums[k + 1]; k -= 1 { }
            } else if sum < target {
                for j += 1; j < k && nums[j] == nums[j - 1]; j += 1 { }
            } else {
                for k -= 1; j < k && nums[k] == nums[k + 1]; k -= 1 { }
            }
        }
    } 

    return result
}