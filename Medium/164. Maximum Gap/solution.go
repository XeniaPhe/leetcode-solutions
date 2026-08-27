package maximum_gap

func maximumGap(nums []int) (maxGap int) {
    for sh, srt := 0, make([]int, len(nums)); sh < 32; nums, srt, sh = srt, nums, sh + 8 {
        var counts [256]int
        for _, num := range nums {
            counts[(num >> sh) & 0xff] += 1
        }

        for sum, i := 0, 0; i < 256; i += 1 {
            counts[i], sum = sum, sum + counts[i]
        }

        for _, num := range nums {
            idx := (num >> sh) & 0xff
            counts[idx], srt[counts[idx]] = counts[idx] + 1, num
        }
    }

    for prev, i := nums[0], 1; i < len(nums); i += 1 {
        prev, maxGap = nums[i], max(maxGap, nums[i] - prev)
    }

    return maxGap
}