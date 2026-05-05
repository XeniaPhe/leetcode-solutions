package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
    lower := 0
    for right := len(nums); lower < right; {
        mid := lower + (right - lower) / 2

        if nums[mid] < target {
            lower = mid + 1
        } else {
            right = mid
        }
    }

    if len(nums) <= lower || nums[lower] != target {
        return []int {-1, -1}
    }

    upper := lower
    for right := len(nums); upper < right; {
        mid := upper + (right - upper) / 2

        if nums[mid] <= target {
            upper = mid + 1
        } else {
            right = mid
        }
    }

    return []int {lower, upper - 1}
}