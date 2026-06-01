package unique_binary_search_trees

func numTrees(n int) int {
    memo := make([]int, n)
    memo[0] = 1
    return numBst(memo, n - 1, 1, n)
}

func numBst(memo []int, m int, lb int, ub int) int {
    total := memo[min(max(ub - lb, 0), m)]
    if total != 0 {
        return total
    }

    for i := lb; i <= ub; i += 1 {
        total += numBst(memo, m, lb, i - 1) * numBst(memo, m, i + 1, ub)
    }

    memo[ub - lb] = total
    return total
}