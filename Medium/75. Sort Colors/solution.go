package sort_colors

// Counting Sort
func sortColors(nums []int)  {
    var colorCounts [3]int
    for _, num := range nums {
        colorCounts[num] += 1
    }

    for i, j := 0, 0; i < 3; i += 1 {
        for lim := j + colorCounts[i]; j < lim; j += 1 {
            nums[j] = i
        }
    }
}

// Follow up: One-pass algorithm using constant space
func sortColorsOnePass(nums []int)  {
    positions := [4]int {-1, 0, 0, 0}
    for i := 0; i < len(nums); i += 1 {
        for j := nums[i] + 1; j < 4; j += 1 {
            if positions[j - 1] <= positions[j] {
                nums[positions[j]] = j - 1
            }

            positions[j] += 1
        }
    }
}