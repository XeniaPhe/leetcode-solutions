package longest_palindromic_substring

/*
Complexity:
    Time Complexity: O(n^2)
    Space Complexity: O(n)

    n: Size of the input
*/

func longestPalindrome(s string) string {
    // Since strings are byte slices, this cast unfortunately makes a copy.
    // But on the bright side, the solution inherently supports unicode.
    runeSlice := []rune(s)
    lastIdx := len(s) - 1
    center := lastIdx / 2
    longestBegin, longestEnd := 0, 1

    for i := 0; i <= center; i += 1 {
        centerLeft, centerRight := center - i, center + i
        centerEvenLeft, centerEvenRight := centerLeft + 1, centerRight + 1

        oddLeftMinDist := mint2(centerLeft, lastIdx - centerLeft)
        oddRightMinDist := mint2(centerRight, lastIdx - centerRight)
        evenLeftMinDist := mint2(centerLeft, lastIdx - centerEvenLeft)
        evenRightMinDist := mint2(centerRight, lastIdx - centerEvenRight)

        longestBegin, longestEnd = longestOddPalindrome(runeSlice, centerLeft, oddLeftMinDist,
                                                        longestBegin, longestEnd)
        longestBegin, longestEnd = longestOddPalindrome(runeSlice, centerRight, oddRightMinDist,
                                                        longestBegin, longestEnd)
        longestBegin, longestEnd = longestEvenPalindrome(runeSlice, centerLeft, centerEvenLeft, evenLeftMinDist,
                                                        longestBegin, longestEnd)
        longestBegin, longestEnd = longestEvenPalindrome(runeSlice, centerRight, centerEvenRight, evenRightMinDist,
                                                        longestBegin, longestEnd)
    }

    return string(runeSlice[longestBegin:longestEnd])
}

func longestOddPalindrome(s []rune, center int, maxDist int, longestBegin int, longestEnd int) (int, int) {
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

func longestEvenPalindrome(s []rune, centerLeft int, centerRight int,
                            maxDist int, longestBegin int, longestEnd int) (int, int) {

    longest := longestEnd - longestBegin
    if 2 * maxDist + 2 <= longest {
        return longestBegin, longestEnd
    }

    ln := 0
    for i := 0; i <= maxDist; ln, i = ln + 2, i + 1 {
        if s[centerLeft - i] != s[centerRight + i] {
            break
        }
    }

    if ln <= longest {
        return longestBegin, longestEnd
    }

    width := (ln - 2) / 2
    return centerLeft - width, centerRight + width + 1
}

func mint2(a int, b int) int {
    if a <= b {
        return a
    }

    return b
}