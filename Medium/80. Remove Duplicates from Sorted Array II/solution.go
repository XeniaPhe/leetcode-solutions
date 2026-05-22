package remove_duplicates_from_sorted_array_II

func removeDuplicates(nums []int) int {
    write := 1
    for i, count := 1, 1; i < len(nums); i += 1 {
        if nums[i] == nums[i - 1] {
            if count += 1; count > 2 {
                continue
            }
        } else {
            count = 1
        }
        
        nums[write], write = nums[i], write + 1
    }

    return write
}