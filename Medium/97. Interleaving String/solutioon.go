package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    s, l, t, i, j := s1, s2, s3, 0, 0
    if len(s) > len(l) {
        s, l = l, s
    }

    dp := make([]bool, len(s) + 1)
    for dp[0], i = true, 1; i <= len(s); i += 1 {
        dp[i] = dp[i - 1] && s[i - 1] == t[i - 1]
    }

    for i = 1; i <= len(l); i += 1 {
        for dp[0], j = dp[0] && l[i - 1] == t[i - 1], 1; j <= len(s); j += 1 {
            dp[j] = (dp[j - 1] && s[j - 1] == t[i + j - 1]) || (dp[j] && l[i - 1] == t[i + j - 1])
        }
    }

    return dp[len(s)]
}