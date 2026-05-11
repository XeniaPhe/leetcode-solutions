package unique_paths_II

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    g, m, n := obstacleGrid, len(obstacleGrid), len(obstacleGrid[0])
    for i, d := 0, 1; i < n; i, d, g[0][i] = i + 1, d - d * g[0][i], d - d * g[0][i] { }
    for i, d := 1, g[0][0]; i < m; i, d, g[i][0] = i + 1, d - d * g[i][0], d - d * g[i][0] { }

    for i := 1; i < m; i += 1 {
        for j := 1; j < n; j, g[i][j] = j + 1, (1 - g[i][j]) * (g[i][j - 1] + g[i - 1][j]) { }
    }

    return g[m - 1][n - 1]
}