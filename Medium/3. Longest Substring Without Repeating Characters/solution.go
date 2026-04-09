package longest_substring_without_repeating_characters

/*
Complexity:
    Time Complexity           : O(n)
    Total Space Complexity    : O(n)
    Auxilary Space Complexity : O(n)

    n: Size of the input
*/

func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int, len(s))
    longestRun := 0
    currRunStart := 0

    for i := 0; i < len(s); i++ {
        chr := s[i]
        lastSeenIdx, seen := lastSeen[chr]

        if seen && lastSeenIdx >= currRunStart {
            currentRun := i - currRunStart
            currRunStart = lastSeenIdx + 1

            if currentRun > longestRun {
                longestRun = currentRun
            }
        }

        lastSeen[chr] = i
    }

    currentRun := len(s) - currRunStart
    if currentRun > longestRun {
        longestRun = currentRun
    }

    return longestRun
}