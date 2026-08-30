package largest_number

import "math/bits"
import "strconv"
import "slices"
import "unsafe"
import "cmp"

var decimalPows = [11]int {1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}
var bitLenToDecimalDigits = [35]int {1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}

func largestNumber(nums []int) string {
    if slices.SortFunc(nums, compareDesc); nums[0] == 0 {
        return "0"
    }

    var buf []byte
    for i := 0; i < len(nums); buf, i = strconv.AppendInt(buf, int64(nums[i]), 10), i + 1 { }
	return unsafe.String(unsafe.SliceData(buf), len(buf))
}

func compareDesc(a, b int) int {
	res, sm, lr, smdc, lrdc := 0,  min(a, b), max(a, b), decimalDigits(min(a, b)), decimalDigits(max(a, b))
	for i := 1; res == 0 && i <= smdc; i += 1 {
		res = cmp.Compare((sm / decimalPows[smdc - i]) % 10, (lr / decimalPows[lrdc - i]) % 10)
	}

	for diff, i := lrdc - smdc, 1; res == 0 && i <= diff; i += 1 {
		res = cmp.Compare((lr / decimalPows[lrdc - i]) % 10, (lr / decimalPows[diff - i]) % 10)
	}

	for i := 1; res == 0 && i <= smdc; i += 1 {
		res = cmp.Compare((lr / decimalPows[smdc - i]) % 10, (sm / decimalPows[smdc - i]) % 10)
	}

	if a == lr {
		return res
	} else {
		return -res
	}
}

func decimalDigits(num int) int {
	if digits := bitLenToDecimalDigits[bits.Len(uint(num))]; num >= decimalPows[digits] {
		return digits + 1
	} else {
		return digits
	}
}