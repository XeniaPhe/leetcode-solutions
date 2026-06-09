package triangle

import "slices"

func minimumTotal(triangle [][]int) int {
    for i := 0; i < len(triangle) - 1; i += 1 {
        triangle[i + 1][0] += triangle[i][0]
        triangle[i + 1][i + 1] += triangle[i][i]
        for j := 1; j <= i; j += 1 {
            triangle[i + 1][j] += min(triangle[i][j - 1], triangle[i][j])
        }
    }

    return slices.Min(triangle[len(triangle) - 1])
}