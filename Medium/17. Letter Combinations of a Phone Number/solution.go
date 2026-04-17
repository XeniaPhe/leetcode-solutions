package letter_combinations_of_a_phone_number

import "unsafe"

var digitToLetters = [10]string {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
    combCount := 1
    for _, digit := range digits {
        combCount *= len(digitToLetters[digit - '0'])
    }

    combinations := make([][4]byte, combCount)
    result := make([]string, combCount)
    firstLetters := digitToLetters[digits[0] - '0']
    combCount = len(firstLetters)

    for i := 0; i < combCount; i += 1 {
        combinations[i][0] = firstLetters[i]
    }

    for i := 1; i < len(digits); i += 1 {
        letters := digitToLetters[digits[i] - '0']
        letterCount := len(letters)
        
        if letterCount == 3 {
            copy(combinations[1 * combCount:], combinations[:1 * combCount])
            copy(combinations[2 * combCount:], combinations[:1 * combCount])
            copy(combinations[3 * combCount:], combinations[:1 * combCount])
        } else {
            copy(combinations[1 * combCount:], combinations[:1 * combCount])
            copy(combinations[2 * combCount:], combinations[:2 * combCount])
        }

        for j := 0; j < letterCount; j += 1 {
            for k := 0; k < combCount; k += 1 {
                combinations[j * combCount + k][i] = letters[j]
            }
        }

        combCount *= letterCount
    }
    
    for i := 0; i < combCount; i += 1 {
        result[i] = unsafe.String(&combinations[i][0], len(digits))
    }

    return result
}