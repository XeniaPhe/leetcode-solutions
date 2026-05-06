package combination_sum

import "slices"

func combinationSum(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    return sumTo([][]int{}, []int{}, candidates, target)
}

func sumTo(combs [][]int, currPath []int, items []int, target int) [][]int {
    if target == 0 {
        return append(combs, slices.Clone(currPath))
    } else if len(items) == 0 {
        return combs
    }

    for i := 0; i < len(items); i += 1 {
        num := items[i]
        if num > target {
            break
        }

        atMost := target / num
        for nxt, j := target - num, 0; j < atMost; nxt, j = nxt - num, j + 1 {
            currPath = append(currPath, num)
            combs = sumTo(combs, currPath, items[i + 1:], nxt)
        }

        currPath = currPath[:len(currPath) - atMost]
    }

    return combs
}