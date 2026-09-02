package number_of_islands

func numIslands(grid [][]byte) (count int) {
    for i, m, n, stack := 0, len(grid) - 1, len(grid[0]) - 1, [][2]int{}; i <= m; i += 1 {
        for j := 0; j <= n; j += 1 {
            if grid[i][j] == '0' {
                continue
            }
            
            for count, grid[i][j], stack = count + 1, '0', append(stack, [2]int{i, j}); len(stack) > 0; {
                ii, jj := stack[len(stack) - 1][0], stack[len(stack) - 1][1]
                stack = stack[:len(stack) - 1]
                iu, id, jl, jr := max(ii - 1, 0), min(ii + 1, m), max(jj - 1, 0), min(jj + 1, n)

                for _, coord := range [4][2]int {{iu, jj}, {id, jj}, {ii, jl}, {ii, jr}} {
                    if iii, jjj := coord[0], coord[1]; grid[iii][jjj] == '1' {
                        grid[iii][jjj], stack = '0', append(stack, [2]int{iii, jjj})
                    }
                }
            }
        }
    }

    return count
}