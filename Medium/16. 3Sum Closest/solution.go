package three_sum_closest

import "slices"

func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
    bestSum := nums[0] + nums[1] + nums[2]
    bestDiff := abs(target - bestSum)

    for i := 0; i < len(nums) - 2; i += 1 {
        j, k := i + 1, len(nums) - 1
        
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            diff := abs(target - sum)
            
            if diff < bestDiff {
                bestDiff = diff
                bestSum = sum
            }

            if sum <= target {
                j += 1
            } else {
                k -= 1
            }
        }
    }

    return bestSum
}

func abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}