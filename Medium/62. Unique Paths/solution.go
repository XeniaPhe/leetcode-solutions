package unique_paths

func uniquePaths(m int, n int) int {
    m, n = m - 1, n - 1
    sFact, manhFact, i, manhDst, s, l := 1, 1, 2, m + n, min(m, n), max(m, n)
    for ; i <= s; sFact, i = sFact * i, i + 1 { }
    for i = l + 1; i <= manhDst && manhFact % sFact != 0; manhFact, i = manhFact * i, i + 1 { }
    for manhFact /= sFact; i <= manhDst; manhFact, i = manhFact * i, i + 1 { }
    return manhFact
}