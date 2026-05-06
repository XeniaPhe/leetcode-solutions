package combination_sum_II

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    return sumTo([][]int{}, []int{}, candidates, target)
}

func sumTo(combs [][]int, currPath []int, items []int, target int) [][]int {
    if target == 0 {
        return append(combs, slices.Clone(currPath))
    } else if len(items) == 0 {
        return combs
    }

    for prev, i := items[0] - 1, 0; i < len(items); i += 1 {
        if items[i] > target {
            break
        } else if items[i] == prev {
            continue
        }

        prev = items[i]
        combs = sumTo(combs, append(currPath, prev), items[i + 1:], target - prev)
    }

    return combs
}