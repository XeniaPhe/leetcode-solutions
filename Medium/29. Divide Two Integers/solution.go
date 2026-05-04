package divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
    target := int32(dividend)
    div := int32(divisor)
    quot := int32(0)
    neg := true
    
    if target > 0 {
        target = -target
        neg = !neg
    }

    if div > 0 {
        div = -div
        neg = !neg
    }

    divUpperBound := int32(0)
    for left, right := int32(0), int32(31); left <= right; {
        mid := (left + right) >> 1

        if int32(-1) << mid > div {
            left = mid + 1
        } else {
            divUpperBound = mid
            right = mid - 1
        }
    }

    for sum, limit, factor, right := int32(0), int32(0), int32(1), int32(31) - divUpperBound; ; factor = int32(1) {
        for left := int32(0); left <= right; {
            mid := (left + right) >> 1
            val := div << mid

            if val >= target {
                sum = val
                limit = mid
                factor = int32(-1) << mid
                left = mid + 1
			} else {
                right = mid - 1
            }
        }

        if factor == 1 {
            break
        }

        target -= sum
        right = limit - 1
        quot += factor
    }

    if !neg {
        return int(quot)
    } else if quot == math.MinInt32 {
        return math.MaxInt32
    }

    return int(-quot)
}