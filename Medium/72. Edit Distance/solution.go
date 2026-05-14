package edit_distance

func minDistance(word1 string, word2 string) int {
    smaller, larger, m, n := word1, word2, len(word1), len(word2)
    if len(smaller) > len(larger) {
        smaller, larger, m, n = larger, smaller, n, m
    } else if m == 0 {
        return n
    }

    dpRead, dpWrite := make([]int, m + 1), make([]int, m + 1)
    for j := 0; j <= m; dpRead[j], j = j, j + 1 { }
    
    for i, j := 1, 1; i <= n; dpRead, dpWrite, i = dpWrite, dpRead, i + 1 {
        for dpWrite[0], j = i, 1; j <= m; j += 1 {
            if larger[i - 1] == smaller[j - 1] {
                dpWrite[j] = dpRead[j - 1]
            } else {
                dpWrite[j] = 1 + min(dpWrite[j - 1], dpRead[j - 1], dpRead[j])
            }
        }
    }

    return dpRead[m]
}