package decode_ways

func numDecodings(s string) int {
    groupSize, count, fib := 0, 1, []int {1, 1}
    updateFib := func() {
        for groupSize >= len(fib) {
            fib = append(fib, fib[len(fib) - 2] + fib[len(fib) - 1])
        }
    }

    for i, prev, digit := 0, byte('-'), byte('-'); i < len(s); prev, i = digit, i + 1 {
        if digit = s[i]; digit == '1' || digit == '2' {
            groupSize += 1
            continue
        } else if prev == '1' || (prev == '2' && digit < '7') {
            if groupSize += 1; digit == '0' {
                groupSize -= 2
            }
        } else if digit != '0' && groupSize == 0 {
            continue
        } else if digit == '0' {
            return 0
        }
        
        updateFib()
        count, groupSize = count * fib[groupSize], 0
    }

    updateFib();
    return count * fib[groupSize]
}