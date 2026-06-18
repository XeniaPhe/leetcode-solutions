package single_number_II

func singleNumber(nums []int) int {
    result, oneCounts := 0, [64]int{}
    for i := 0; i < len(nums); i += 1 {
        for num, j := uint(nums[i]), 0; num != 0; num, j = num >> 1, j + 1 {
            if num & 1 == 1 {
                oneCounts[j] += 1
            }
        }
    }

    for i := 0; i < len(oneCounts); i += 1 {
        if oneCounts[i] % 3 == 1 {
            result += 1 << i
        }
    }

    return result
}