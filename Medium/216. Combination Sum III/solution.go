package combination_sum_III

func combinationSum3(k int, n int) [][]int {
	type stackType struct {
		comb []int
		sum, startsAt int
	}

	if (k * (k - 1)) >> 1 > n || (k * (19 - k)) >> 1 < n {
		return nil
	}

    res, stck := make([][]int, 0), []stackType{{make([]int, 0, k), 0, 1}}
    for curr, n2, rem := (stackType{}), n << 1, 0; len(stck) > 0; {
        curr, stck = stck[len(stck) - 1], stck[:len(stck) - 1]
        if rem = k - len(curr.comb); rem == 1 {
            res = append(res, append(curr.comb, n - curr.sum))
        } else {
			minPick := max(curr.startsAt, n - curr.sum - (((rem - 1) * (20 - rem)) >> 1))
			maxPick := min(10 - rem, (n2 - (curr.sum << 1) - (rem * rem) + rem) / (rem << 1))
			stck = append(stck, stackType{append(curr.comb, minPick), curr.sum + minPick, minPick + 1})
			
			for i := minPick + 1; i <= maxPick; i += 1 {
				stck = append(stck, stackType{append(append(make([]int, 0, k), curr.comb...), i), curr.sum + i, i + 1})
			}
		}
    }

    return res
}