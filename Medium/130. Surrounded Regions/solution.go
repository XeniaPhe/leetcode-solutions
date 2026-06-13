package surrounded_regions

func solve(board [][]byte)  {
    m, n := len(board) - 1, len(board[0]) - 1
    free, stack := make([][]bool, m + 1), make([][2]int, 0, 4)
    for i := 0; i <= m; free[i], i = make([]bool, n + 1), i + 1 { }

    for i, pair := range [][2]int{{m, n}, {n, m}} {
        for j, r, c, ls := 0, 0, 0, 0; j <= pair[0]; j += max(1, pair[0]) {
			for k := 0; k <= pair[1]; k += 1 {
				if r, c = j * (1 - i) + k * i, j * i + k * (1 - i); board[r][c] == 'X' || free[r][c] {
					continue
				}

				for stack, ls = append(stack, [2]int{r, c}), 1; ls > 0; ls = len(stack) {
					stack, r, c = stack[:ls - 1], stack[ls - 1][0], stack[ls - 1][1]
					free[r][c] = true
					for _, coord := range [4][2]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
						if r, c = coord[0], coord[1]; r < 0 || r > m || c < 0 || c > n {
							continue
						} else if board[r][c] == 'O' && !free[r][c] {
							stack = append(stack, [2]int{r, c})
						}
					}
				}
			}
		}
    }

    for i := 0; i < m * n; i += 1 {
        if r, c := i / n, i % n; !free[r][c] {
            board[r][c] = 'X'
        }
    }
}