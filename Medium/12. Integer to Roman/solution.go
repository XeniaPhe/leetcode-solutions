package integer_to_roman

import "unsafe"

var symbols = [2][4]byte {
    [4]byte {'I', 'X', 'C', 'M'}, // decimal powers
    [4]byte {'0', 'V', 'L', 'D'}, // half decimal powers
}

func intToRoman(num int) string {
    roman := make([]byte, 15) // 15 is the maximum length of a roman numeral less than 4000 (MMMDCCCLXXXVIII)
    begin := 14

    for i := 0; num > 0; i, num = i + 1, num / 10 {
        digit := num % 10
        nrRomanDigits := 0

        switch digit {
        case 5:
            nrRomanDigits = 1
            roman[begin] = symbols[1][i + 1]
        case 6, 7, 8:
            nrRomanDigits = 1
            digit -= 5
            roman[begin - digit] = symbols[1][i + 1]
            fallthrough
        case 1, 2, 3:
            nrRomanDigits += digit
            for j := 0; j < digit; j += 1 {
                roman[begin - j] = symbols[0][i]
            }
        case 4, 9:
            nrRomanDigits = 2
            roman[begin - 1] = symbols[0][i]
            roman[begin] = symbols[1 - digit / 5][i + 1]
        }

        begin -= nrRomanDigits
    }

    return unsafe.String(unsafe.SliceData(roman[begin + 1:]), 14 - begin)
}