package jump_game_II

// Greedy approach (O(n) time, O(1) space)
func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }

    for target, jmp, max, currMax, i := len(nums) - 1, 1, 0, 0, 0; i < target; i += 1 {
        if maxi := i + nums[i]; maxi > max {
            if max = maxi; max >= target {
                return jmp
            }
        }

        if i == currMax {
            jmp, currMax = jmp + 1, max
        }
    }

    return -1
}

// BFS approach (O(n^2) time, O(n) space)
// Somehow whenever I thought of a greedy algorithm, what I imagined was the bruteforce approach instead
// So this was my best solution (13-25 ms) until I looked up other solutions and saw the actual implementation of the greedy algorithm
func jumpBFS(nums []int) int {
    visited := make([]bool, len(nums) - 1)
    writeBuf := make([]int, 0, len(nums) / 2)
    readBuf := make([]int, 1, len(nums) / 2)
    readBuf[0] = 0

    for target, jmp := len(nums) - 1, 0; true; jmp += 1 {
        for i := 0; i < len(readBuf); i += 1 {
            idx := readBuf[i]
            max := idx + nums[idx]

            if max >= target {
                return jmp
            }

            for j := idx + 1; j <= max; j += 1 {
                if !visited[j] {
                    visited[j] = true
                    writeBuf = append(writeBuf, j)
                }
            }
        }

        readBuf, writeBuf = writeBuf, readBuf
        writeBuf = writeBuf[:0]
    }

    return -1
}