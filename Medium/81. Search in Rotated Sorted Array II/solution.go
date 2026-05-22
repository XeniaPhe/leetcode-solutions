package search_in_rotated_sorted_array_II

func search(nums []int, target int) bool {
    lastIdx := len(nums) - 1
    first, last := nums[0], nums[lastIdx]
    secondHalf := last <= first && target <= last
    firstHalf := !secondHalf && last <= first && target >= first

    for currHalf, left, right := false, 0, lastIdx; left <= right; {
        mid := (left + right) / 2
        if nums[mid] == last && first == last {
            i, end := mid + 1, lastIdx
            if mid < lastIdx - mid {
                i, end = 0, mid - 1
            }

            for ; i <= end; i += 1 {
                if nums[i] != nums[mid] {
                    break
                }
            }

            currHalf = (i > end) == (end == lastIdx)
        } else {
            currHalf = nums[mid] <= last
        }
        
        if firstHalf && currHalf {
            right = mid - 1
            continue
        } else if secondHalf && !currHalf {
            left = mid + 1
            continue
        }

        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            return true
        }
    }
    
    return false
}