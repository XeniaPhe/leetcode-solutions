package valid_sudoku

func isValidSudoku(board [][]byte) bool {
    var seen [9]bool
    valid := true

    resetSeen := func() {
        for i := 0; i < 9; i += 1 {
            seen[i] = false
        }
    }

    checkCell := func(row, col int) {
        if board[row][col] == '.' {
            return
        }

        if seen[board[row][col] - '0' - 1] {
            valid = false
        }

        seen[board[row][col] - '0' - 1] = true
        return
    }

    for i := 0; i < 9; i += 1 {
        resetSeen()
        for j := 0; j < 9; j += 1 {
            checkCell(i, j)
        }

        resetSeen()
        for j := 0; j < 9; j += 1 {
            checkCell(j, i)
        }
    }

    for bi := 0; bi < 3; bi += 1 {
        for bj := 0; bj < 3; bj += 1 {
            resetSeen()
            for ci := bi * 3; ci < (bi + 1) * 3; ci += 1 {
                for cj := bj * 3; cj < (bj + 1) * 3; cj += 1 {
                    checkCell(ci, cj)
                }
            }
        }
    }

    return valid
}