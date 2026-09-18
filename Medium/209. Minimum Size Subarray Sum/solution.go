package minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, res := 0, 0, nums[0], len(nums) - 1
	for left <= right && right < len(nums) {
		if sum < target {
			if right += 1; right < len(nums) {
				sum += nums[right]
			}
		} else {
			left, sum, res = left + 1, sum - nums[left], min(res, right - left)
		}
	}

    if left == 0 && sum < target {
        return 0
    } else {
	    return res + 1
    }
}

func minSubArrayLenFollowup(target int, nums []int) int {
    ln, res, sums := len(nums), len(nums) - 1, make([]int, len(nums) + 1)
    for s, i := 0, 0; i < ln; sums[i], s, i = s, s + nums[i], i + 1 { }

    if sums[ln] = sums[ln - 1] + nums[ln - 1]; sums[ln] < target {
        return 0
    }

    for i := ln - 1; i >= 0 && res > 0; i -= 1 {
        for l, r := i, len(nums); l < r; {
            if mid, sum := (l + r) >> 1, sums[((l + r) >> 1) + 1] - sums[i]; sum >= target {
				r, res = mid, min(res, mid - i)
            } else {
                l = mid + 1
            }
        }
    }

    return res + 1
}