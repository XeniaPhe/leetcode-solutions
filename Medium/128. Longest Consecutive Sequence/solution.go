package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
    seqArena := make([][3]int, 0, len(nums))
    seqMap := make(map[int]int, len(nums))
    seqArena = append(seqArena, [3]int{})
    longestSeq := min(1, len(nums))
    var prev, curr, next int
    
    for _, num := range nums {
        if curr = seqMap[num]; curr != 0 {
            continue
        }

        if prev, next = seqMap[num - 1], seqMap[num + 1]; prev == 0 && next == 0 {
            seqArena, seqMap[num] = append(seqArena, [3]int{1, num, num}), len(seqArena)
            continue
        } else if prev != 0 && next != 0 {
            seqMap[num], seqMap[seqArena[next][2]] = prev, prev
            curr = seqArena[prev][0] + seqArena[next][0] + 1
            seqArena[prev][0], seqArena[prev][2] = curr, seqArena[next][2]
        } else if prev != 0 {
            seqMap[num] = prev
            curr = seqArena[prev][0] + 1
            seqArena[prev][0], seqArena[prev][2] = curr, num
        } else {
            seqMap[num] = next
            curr = seqArena[next][0] + 1
            seqArena[next][0], seqArena[next][1] = curr, num
        }

        longestSeq = max(longestSeq, curr)
    }

    return longestSeq
}