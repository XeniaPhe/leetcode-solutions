package group_anagrams

import (
    "unsafe"
    "bytes"
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)
    var sorter lowercaseEnglishSorter

    for _, str := range strs {
        sorted := sorter.Sorted(str)
        anagrams[sorted] = append(anagrams[sorted], str)
    }

    groups := make([][]string, len(anagrams))
    idx := 0
    for _, group := range anagrams {
        groups[idx] = group
        idx += 1
    }

    return groups
}

type lowercaseEnglishSorter struct {
    letterCounts [26]int
    buffer       bytes.Buffer
}

func (s *lowercaseEnglishSorter) Sorted(str string) string {
    for i := range len(str) {
        s.letterCounts[str[i] - 'a'] += 1
    }

    for i := range 26 {
        count := s.letterCounts[i]
        s.letterCounts[i] = 0

        for range count {
            s.buffer.WriteByte('a' + byte(i))
        }
    }

    bytes := slices.Clone(s.buffer.Bytes())
    s.buffer.Reset()
    return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}