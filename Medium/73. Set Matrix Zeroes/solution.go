package set_matrix_zeroes

func setZeroes(matrix [][]int) {
    zeroRow, zeroCol, m, n := -1, -1, len(matrix), len(matrix[0])
    outer:
    for i := 0; i < m; i += 1 {
        for j := 0; j < n; j += 1 {
            if matrix[i][j] == 0 {
                zeroRow, zeroCol = i, j
                break outer
            }
        }
    }

    if zeroRow == -1 {
        return
    }

    for i := zeroRow + 1; i < m; i += 1 {
        if matrix[i][zeroCol] == 0 {
            matrix[i][zeroCol] = 1
        } else {
            matrix[i][zeroCol] = -1
        }
    }

    for j := 0; j < n; j += 1 {
        if matrix[zeroRow][j] == 0 {
            matrix[zeroRow][j] = 1
        } else {
            matrix[zeroRow][j] = -1
        }
    }

    for i := zeroRow + 1; i < m; i += 1 {
        for j := 0; j < n; j += 1 {
            if matrix[i][j] == 0 {
                matrix[zeroRow][j], matrix[i][zeroCol] = 1, 1
            }
        }
    }

    for i := zeroRow + 1; i < m; i += 1 {
        if matrix[i][zeroCol] == 1 {
            for j := 0; j < n; matrix[i][j], j = 0, j + 1 { }
        }
    }

    for j := 0; j < n; j += 1 {
        if matrix[zeroRow][j] == 1 {
            for i := 0; i < m; matrix[i][j], i = 0, i + 1 { }
        }
    }

    for j := 0; j < n; matrix[zeroRow][j], j = 0, j + 1 { }
}