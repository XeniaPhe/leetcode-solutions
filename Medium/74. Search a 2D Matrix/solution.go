package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    for low, high := 0, (m * n) - 1; low <= high; {
        mid := low + (high - low) / 2
        i, j := mid / n, mid % n

        if matrix[i][j] < target {
            low = mid + 1
        } else if matrix[i][j] > target {
            high = mid - 1
        } else {
            return true
        }
    }

    return false
}