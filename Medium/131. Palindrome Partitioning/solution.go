package palindrome_partitioning

import "slices"

func partition(s string) [][]string {
    pals := make([][]bool, len(s))
    pals[0] = make([]bool, len(s))

    for i := 0; i < len(s); pals[0][i], i = true, i + 1 { }
    if i := 0; len(s) > 1 {
        for pals[1] = make([]bool, len(s) - 1); i < len(s) - 1; i += 1 {
            if s[i] == s[i + 1] {
                pals[1][i] = true
            }
        }
    }

    for i, j := 2, 0; i < len(s); i += 1 {
        for pals[i], j = make([]bool, len(s) - i), 0; j < len(s) - i; j += 1 {
            if pals[i - 2][j + 1] && s[j] == s[i + j] {
                pals[i][j] = true
            }
        }
    }

	path, res := make([]string, 0, len(s)), make([][]string, 0, 4)
    parts(s, 0, pals, &path, &res)
	return res
}

func parts(s string, idx int, pals [][]bool, path *[]string, res *[][]string) {
    for i := 0; i < len(s) - idx; i += 1 {
        if end := idx + i + 1; pals[i][idx] {
            if *path = append(*path, s[idx:end]); end == len(s) {
                *res = append(*res, slices.Clone(*path))
            } else {
                parts(s, end, pals, path, res)
            }

            *path = (*path)[:len(*path) - 1]
        }
    }
}