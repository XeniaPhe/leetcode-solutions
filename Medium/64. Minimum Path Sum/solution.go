package minimum_path_sum

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    for lastRow, i := m - 1, n - 2; i >= 0; i -= 1 {
        grid[lastRow][i] += grid[lastRow][i + 1]
    }

    for lastCol, i := n - 1, m - 2; i >= 0; i -= 1 {
        grid[i][lastCol] += grid[i + 1][lastCol]
    }

    for i := m - 2; i >= 0; i -= 1 {
        for j := n - 2; j >= 0; j -= 1 {
            grid[i][j] += min(grid[i][j + 1], grid[i + 1][j])
        }
    }

    return grid[0][0]
}