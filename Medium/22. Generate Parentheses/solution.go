package generate_parentheses

import "unsafe"

var catalanNumbers = [8]int {1, 2, 5, 14, 42, 132, 429, 1430}

func generateParenthesis(n int) []string {
    var empty [16]byte
    parensBytes := make([][16]byte, 0, catalanNumbers[n - 1])
    parensBytes = genParens(parensBytes, empty, 0, 0, n)
    parens := make([]string, len(parensBytes))

    for i := range len(parens) {
        parens[i] = unsafe.String(&parensBytes[i][0], 2 * n)
    }

    return parens
}

func genParens(parens [][16]byte, current [16]byte, open int, close int, n int) [][16]byte {
    currIdx := open + close
    if currIdx == 2 * n {
        return append(parens, current)
    }

    if open < n {
        current[currIdx] = '('
        parens = genParens(parens, current, open + 1, close, n)
    }
    if open > close {
        current[currIdx] = ')'
        parens = genParens(parens, current, open, close + 1, n)
    }

    return parens
}