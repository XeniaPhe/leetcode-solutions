package gray_code

import "math/bits"

func grayCode(n int) []int {
    result, visited := make([]int, 1 << n), make([]bool, 1 << n)
    result[len(result) - 1], visited[0], visited[1] = 1, true, true
    _ = greyCode(result, visited, n, 1)
    return result
}

func greyCode(result []int, visited []bool, n int, depth int) bool {
    idxL, idxR := depth, (1 << n) - depth - 1
    lastL, lastR := result[idxL - 1], result[idxR + 1]
    nextDepth, dist := depth + 1, idxR - idxL

    for i := 0; i < n; i += 1 {
        numL := lastL ^ (1 << i)
        if visited[numL] {
            continue
        }

        for j := 0; j < n; j += 1 {
            numR := lastR ^ (1 << j)
            if visited[numR] || numL == numR {
                continue
            }

            if dist <= n {
                if diff := bits.OnesCount(uint(numL ^ numR)); diff > dist {
                    continue
                } else if dist == 1 {
                    result[idxL], result[idxR] = numL, numR
                    return true
                }
            }

            result[idxL], result[idxR] = numL, numR
            visited[numL], visited[numR] = true, true

            if greyCode(result, visited, n, nextDepth) {
                return true
            }

            visited[numL], visited[numR] = false, false
        }
    }

    return false
}