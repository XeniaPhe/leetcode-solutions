package spiral_matrix

func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    total := m * n
    result := make([]int, total)

    for x, y, curr, layer := -1, -1, 0, 0; curr < total; layer += 1 {
        for x, y = x + 1, y + 1; curr < total && x < n - layer; curr, x = curr + 1, x + 1 {
            result[curr] = matrix[y][x]
        }

        for x, y = x - 1, y + 1; curr < total && y < m - layer; curr, y = curr + 1, y + 1 {
            result[curr] = matrix[y][x]
        }

        for x, y = x - 1, y - 1; curr < total && x >= layer; curr, x = curr + 1, x - 1 {
            result[curr] = matrix[y][x]
        }

        for x, y = x + 1, y - 1; curr < total && y > layer; curr, y = curr + 1, y - 1 {
            result[curr] = matrix[y][x]
        }
    }

    return result
}