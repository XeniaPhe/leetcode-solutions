package bitwise_and_of_numbers_range

func rangeBitwiseAnd(left int, right int) int {
    sh := 0
    for ; left != right; sh, left, right = sh + 1, left >> 1, right >> 1 { }
    return left << sh
}