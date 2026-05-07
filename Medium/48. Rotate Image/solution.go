package rotate_image

func rotate(matrix [][]int)  {
    for n, layer := len(matrix), 0; layer < n - 1; layer += 1 {
        for row1, col1 := layer, layer; col1 < n - layer - 1; col1 += 1 {
            row2, col2 := col1, n - row1 - 1
            row3, col3 := col2, n - col1 - 1
            row4, col4 := col3, row1

            matrix[row1][col1], matrix[row2][col2], matrix[row3][col3], matrix[row4][col4] =
            matrix[row4][col4], matrix[row1][col1], matrix[row2][col2], matrix[row3][col3]
        }
    }
}