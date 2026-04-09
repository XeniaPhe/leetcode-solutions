package string_to_integer

import (
    "strings"
    "math"
)

func myAtoi(s string) int {
    s = strings.Trim(s, " ")
    if len(s) == 0 {
        return 0
    }

    var sign int64 = 1
    var max int64 = int64(math.MaxInt32)

    // Some really unnecessary but beautiful ascii magic to get away with just one branch
    if s[0] == '-' || s[0] == '+' {
        diff := int64(s[0] - '+')
        sign -= diff
        max += diff / 2
        s = s[1:]
    }

    var number int64 = 0
    for i := 0; i < len(s); i++ {
        chr := s[i]
        if chr < '0' || chr > '9' {
            break
        }

        number = 10 * number + int64(chr - '0')
        if number >= max {
            number = max
            break
        }
    }

    return int(sign * number)
}