package find_peak_element

func findPeakElement(nums []int) int {
    if len(nums) > 1 && nums[0] > nums[1] {
        return 0
    }

    for lo, hi := 1, len(nums) - 2; lo <= hi; {
        if i := (lo + hi) / 2; nums[i - 1] < nums[i] && nums[i + 1] < nums[i] {
            return i
        } else if nums[i + 1] >= nums[i - 1] {
            lo = i + 1
        } else {
            hi = i - 1
        }
    }

    return len(nums) - 1
}