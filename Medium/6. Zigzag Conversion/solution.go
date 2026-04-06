package zigzag_conversion

/*
Complexity:
    Time Complexity: O(n)
    Space Complexity: O(n)

    n: Size of the input string
*/

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

	// Since this is a copy, the space complexity is O(n) even when excluding the result
    runeSlice := []rune(s)
    ln := len(s)
    result := make([]rune, ln)
    resultIdx, lastRowIdx := 0, numRows - 1

    for i := 0; i < numRows; i += 1 {
        index1 := i % lastRowIdx
        index2 := (lastRowIdx - i) % lastRowIdx
        offset1 := 2 * (numRows - index1) - 2
        offset2 := 2 * (numRows - index2) - 2
        offSwitchBuf := [2]int {offset1, offset2}
        currOffset, prevIdx := 0, i

        if prevIdx >= ln {
            break
        }

        result[resultIdx] = runeSlice[prevIdx]
        resultIdx += 1

        for {
            prevIdx += offSwitchBuf[currOffset]
            if prevIdx >= ln {
                break
            }

            result[resultIdx] = runeSlice[prevIdx]
            currOffset = (currOffset + 1) % 2
            resultIdx += 1
        }
    }

    return string(result)
}