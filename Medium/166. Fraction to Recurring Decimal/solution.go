package fraction_to_recurring_decimal

import "math"
import "unsafe"
import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
    buf, num, denom := []byte{}, int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator)))
    if numerator != 0 && (numerator < 0) != (denominator < 0) {
        buf = append(buf, '-')
    }

    if buf, num = strconv.AppendInt(buf, int64(num / denom), 10), num % denom; num != 0 {
        buf = append(buf, '.')
        for zeros, divs := -1, make(map[int]int); num != 0; {
            if pos, exists := divs[num]; exists {
                buf, num = append(append(buf[:pos], append([]byte{'('}, buf[pos:]...)...), ')'), 0
            } else {
                for zeros, divs[num] = -1, len(buf); num < denom; num, zeros = 10 * num, zeros + 1 { }
                for ; zeros > 0 ; buf, zeros = append(buf, '0'), zeros - 1 { }
                buf, num = strconv.AppendInt(buf, int64(num / denom), 10), num % denom
            }
        }
    }

	return unsafe.String(unsafe.SliceData(buf), len(buf))
}