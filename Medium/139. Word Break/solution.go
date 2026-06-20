package word_break

func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]struct{}, len(wordDict))
    lo, hi := len(wordDict[0]), len(wordDict[0])

    for i := 0; i < len(wordDict); i += 1 {
        dict[wordDict[i]] = struct{}{}
        lo, hi = min(lo, len(wordDict[i])), max(hi, len(wordDict[i]))
    }

    dp := make([]bool, len(s) + 1)
    dp[len(s)] = true

    for i := len(s) - lo; i >= 0; i -= 1 {
        for j, ij := lo, lo + i; ij <= len(s) && j <= hi; j, ij = j + 1, ij + 1 {
            if !dp[ij] {
                continue
            } else if _, exists := dict[s[i:ij]]; exists {
                dp[i] = true
                break
            }
        }
    }

    return dp[0]
}