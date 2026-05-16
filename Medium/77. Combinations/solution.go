package combinations

import "slices"

func combine(n int, k int) [][]int {
    combCount := combCount(n, k)
    nums, res := make([]int, k), make([][]int, combCount)
    for i := 0; i < k; nums[i], i = i + 1, i + 1 { }

    for i, m, maxIdx, shIdx := 0, n - k + 1, k - 1, k - 1; i < combCount; i, shIdx = i + 1, maxIdx {
        for res[i] = slices.Clone(nums); shIdx >= 0; shIdx -= 1 {
            if nums[shIdx] != m + shIdx {
                nums[shIdx] += 1
                for j := shIdx + 1; j < k; j, nums[j] = j + 1, nums[j - 1] + 1 { }
                break
            }
        }
    }

    return res
}

func combCount(n, k int) int {
    count := 1
    for den, denFact, denEnd, num := 2, 1, min(n - k, k), max(n - k, k) + 1; num <= n; num += 1 {
        if denFact == 1 && den <= denEnd {
            den, denFact = den + 1, den * denFact
        }

        if count *= num; denFact > 1 && count % denFact == 0 {
            denFact, count = 1, count / denFact
        }
    }

    return count
}