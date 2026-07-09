package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
    left, right, last := 0, len(nums), nums[len(nums) - 1]
    if nums[0] < last {
        return nums[0]
    }

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] > last {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[left]
}