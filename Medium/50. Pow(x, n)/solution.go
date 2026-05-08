package pow_x_n

func myPow(x float64, n int) float64 {
    pow := int64(n)
    if pow < 0 {
        pow *= -1
    }

    result := float64(1)
    for curr := x; pow != 0; curr, pow = curr * curr, pow >> 1 {
        if pow & 1 == 1 {
            result *= curr
        }
    }

    if n < 0 {
        result = 1 / result
    }
    
    return result
}