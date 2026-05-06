package multiply_strings

import "unsafe"

func multiply(num1 string, num2 string) string {
    if len(num1) < len(num2) {
        num1, num2 = num2, num1
    }

    inter := make([]byte, len(num1) + 1)
    res := make([]byte, len(num1) + len(num2) + 1)

    for i := 0; i < len(res); i += 1 {
        res[i] = '0'
    }

    for carry, i := byte(0), len(num2) - 1; i >= 0; carry, i = 0, i - 1 {
        for sd, j := num2[i] - '0', len(num1) - 1; j >= 0; j -= 1 {
            mult := carry + sd * (num1[j] - '0')
            carry = mult / 10
            inter[j + 1] = '0' + (mult % 10)
        }

        inter[0] = '0' + carry
        carry = 0

        for j, k := i + len(res) - len(num2), len(inter) - 1; k >= 0; j, k = j - 1, k - 1 {
            sum := carry + res[j] + inter[k] - '0' - '0'
            carry = sum / 10
            res[j] = '0' + (sum % 10)
        }

        res[i + len(res) - len(num2) - len(inter)] = '0' + carry
    }

    for i := 0; i < len(res); i += 1 {
        if res[i] != '0' {
            return unsafe.String(unsafe.SliceData(res[i:]), len(res) - i)
        }
    }

    return "0"
}