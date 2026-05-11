package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
    index, begin, end := 0, newInterval[0], newInterval[1]
    for right := len(intervals); index < right; {
        mid := index + (right - index) / 2

        if intervals[mid][0] < begin {
            index = mid + 1
        } else {
            right = mid
        }
    }
    
    if index != 0 && intervals[index - 1][1] >= begin {
        return mergeIfNeeded(intervals, index - 1, end)
    }

    if index != len(intervals) {
        if current := intervals[index]; current[0] == begin {
            return mergeIfNeeded(intervals, index, end)
        } else if current[0] <= end {
            current[0] = begin
            return mergeIfNeeded(intervals, index, end)
        }
    } else {
        return append(intervals, newInterval)
    }

    intervals = append(intervals[:index + 1], intervals[index:]...)
    intervals[index] = newInterval
    return intervals
}

func mergeIfNeeded(intervals[][]int, index int, newEnd int) [][]int {
    if intervals[index][1] < newEnd {
        intervals[index][1] = newEnd
        for i, j := index, index + 1; i < len(intervals); index, i, j = index + 1, j, j + 1 {
            begin, end := intervals[i][0], intervals[i][1]
            for ; j < len(intervals) && intervals[j][0] <= end; j += 1 {
                end = max(end, intervals[j][1])
            }

            intervals[index][0], intervals[index][1] = begin, end
        }

        intervals = intervals[:index]
    }
    
    return intervals
}