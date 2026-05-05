package search_in_rotated_sorted_array

func search(nums []int, target int) int {
    first, last := nums[0], nums[len(nums) - 1]
    firstHalf := last < first && last < target
    secondHalf := last < first && last >= target

    for left, right := 0, len(nums) - 1; left <= right; {
        mid := left + (right - left) / 2

        if firstHalf && nums[mid] < last {
            right = mid - 1
            continue
        } else if secondHalf && nums[mid] > last {
            left = mid + 1
            continue
        }

        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            return mid
        }
    }

    return -1
}