package merge_intervals

import "slices"

func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
    curr := 0
    
    for i, j := 0, 1; i < len(intervals); curr, i, j = curr + 1, j, j + 1 {
        begin, end := intervals[i][0], intervals[i][1]
        for ; j < len(intervals) && intervals[j][0] <= end; j += 1 {
            end = max(end, intervals[j][1])
        }

        intervals[curr][0], intervals[curr][1] = begin, end
    }

    return intervals[:curr]
}