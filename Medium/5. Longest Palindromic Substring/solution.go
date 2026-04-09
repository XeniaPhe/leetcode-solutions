package longest_palindromic_substring

func longestPalindrome(s string) string {
    lastIdx := len(s) - 1
    center := lastIdx / 2
    begin, end := 0, 1

    for i := 0; i <= center; i += 1 {
        centerLeft, centerRight := center - i, center + i
        centerEvenLeft, centerEvenRight := centerLeft + 1, centerRight + 1

        oddLeftMinDist := mint2(centerLeft, lastIdx - centerLeft)
        oddRightMinDist := mint2(centerRight, lastIdx - centerRight)
        evenLeftMinDist := mint2(centerLeft, lastIdx - centerEvenLeft)
        evenRightMinDist := mint2(centerRight, lastIdx - centerEvenRight)

        begin, end = longestOddPalindrome(s, centerLeft, oddLeftMinDist, begin, end)
        begin, end = longestOddPalindrome(s, centerRight, oddRightMinDist, begin, end)
        begin, end = longestEvenPalindrome(s, centerLeft, centerEvenLeft, evenLeftMinDist, begin, end)
        begin, end = longestEvenPalindrome(s, centerRight, centerEvenRight, evenRightMinDist, begin, end)
    }

    return s[begin:end]
}

func longestOddPalindrome(s string, center int, maxDist int, longestBegin int, longestEnd int) (int, int) {
    longest := longestEnd - longestBegin
    if 2 * maxDist < longest {
        return longestBegin, longestEnd
    }

    ln := 1
    for i := 1; i <= maxDist; ln, i = ln + 2, i + 1 {
        if s[center - i] != s[center + i] {
            break
        }
    }

    if ln <= longest {
        return longestBegin, longestEnd
    }

    width := (ln - 1) / 2
    return center - width, center + width + 1
}

func longestEvenPalindrome(s string, cLeft int, cRight int, maxDist int, longestBegin int, longestEnd int) (int, int) {
    longest := longestEnd - longestBegin
    if 2 * maxDist + 2 <= longest {
        return longestBegin, longestEnd
    }

    ln := 0
    for i := 0; i <= maxDist; ln, i = ln + 2, i + 1 {
        if s[cLeft - i] != s[cRight + i] {
            break
        }
    }

    if ln <= longest {
        return longestBegin, longestEnd
    }

    width := (ln - 2) / 2
    return cLeft - width, cRight + width + 1
}

func mint2(a int, b int) int {
    if a <= b {
        return a
    }

    return b
}