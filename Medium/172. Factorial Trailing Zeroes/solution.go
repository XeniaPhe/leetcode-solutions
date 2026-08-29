package factorial_trailing_zeroes

func trailingZeroes(n int) int {
    a, b, c, d, e := n / 5, n / 25, n / 125, n / 625, n / 3125
    single, double, triple, quadruple, quintuple := a - b, b - c, c - d, d - e, e
    return 1 * single + 2 * double + 3 * triple + 4 * quadruple + 5 * quintuple
}