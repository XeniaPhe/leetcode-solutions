package zigzag_conversion

import "unsafe"

/*
Complexity:
    Time Complexity           : O(n)
    Total Space Complexity    : O(n)
    Auxilary Space Complexity : O(1)

    n: Size of the input
*/

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    ln := len(s)
    result := make([]byte, ln)
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

        result[resultIdx] = s[prevIdx]
        resultIdx += 1

        for {
            prevIdx += offSwitchBuf[currOffset]
            if prevIdx >= ln {
                break
            }

            result[resultIdx] = s[prevIdx]
            currOffset = (currOffset + 1) % 2
            resultIdx += 1
        }
    }

    return unsafe.String(unsafe.SliceData(result), ln)
}