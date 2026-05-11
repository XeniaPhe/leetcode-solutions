package spiral_matrix_II

func generateMatrix(n int) [][]int {
    result := make([][]int, n)
    for i := 0; i < n; i += 1 {
        result[i] = make([]int, n)
    }

    for x, y, layer, count, size := -1, -1, 0, 1, n * n; count <= size; layer += 1 {
        for x, y = x + 1, y + 1; count <= size && x < n - layer; count, x = count + 1, x + 1 {
            result[y][x] = count
        }

        for x, y = x - 1, y + 1; count <= size && y < n - layer; count, y = count + 1, y + 1 {
            result[y][x] = count
        }

        for x, y = x - 1, y - 1; count <= size && x >= layer; count, x = count + 1, x - 1 {
            result[y][x] = count
        }

        for x, y = x + 1, y - 1; count <= size && y > layer; count, y = count + 1, y - 1 {
            result[y][x] = count
        }
    }

    return result
}