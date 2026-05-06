package count_and_say

import (
    "bytes"
    "strconv"
)

func countAndSay(n int) string {
    var buf bytes.Buffer
    buf.WriteByte('1')

    for i := 2; i <= n; i += 1 {
        str := buf.String()
        buf.Reset()
        current := str[0]
        repeat := 1

        for j := 1; j < len(str); j += 1 {
            if current == str[j] {
                repeat += 1
                continue
            }

            buf.WriteString(strconv.Itoa(repeat))
            buf.WriteByte(current)
            current = str[j]
            repeat = 1
        }

        buf.WriteString(strconv.Itoa(repeat))
        buf.WriteByte(current)
    }

    return buf.String()
}