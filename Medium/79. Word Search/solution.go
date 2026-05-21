package word_search

import "math"

type dfsNode struct {
    row, col, depth, wordIdx int
    processed, bounded bool
}

func newDfsNode(row, col, depth, wordIdx int, bounded bool) dfsNode {
    return dfsNode{row: row, col: col, depth: depth, wordIdx: wordIdx, bounded: bounded, processed: false}
}

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

var directions [][2]int = [][2]int {[2]int {-1, 0}, [2]int {1, 0}, [2]int {0, -1}, [2]int {0, 1}}

func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    minScore, maxSearchDist, searchLetterIdx, wordDir := math.MaxInt, 0, 0, 1
    var wordLetters, boardLetters ['z' - 'A' + 1]int

    for i := 0; i < len(word); i += 1 {
        wordLetters[word[i] - 'A'] += 1
    }

    for i := 0; i < m; i += 1 {
        for j := 0; j < n; j += 1 {
            boardLetters[board[i][j] - 'A'] += 1
        }
    }

    for i, wl := range wordLetters {
        if wl == 0 {
            continue
        }

        minDist, dir, letterIdx := math.MaxInt, 1, 0
        for dist, letter, j := 0, byte(i) + 'A', 0; wl != 0; j += 1 {
            if word[j] != letter {
                continue
            }

            if wl, dist = wl - 1, min(j, len(word) - j - 1); dist <= minDist {
                if minDist, letterIdx, dir = dist, j, -1; dist == j {
                    dir = 1
                }
            }
        }

        if boardLetters[i] < wl {
            return false
        } else if score := boardLetters[i] * (minDist + 1); score <= minScore {
            minScore, maxSearchDist, wordDir, searchLetterIdx = score, minDist, dir, letterIdx
        }
    }
            
    exists, stack := false, make([]dfsNode, 0, 8)
    startIdx, endIdx, nextLetterIdx := 0, len(word) - 1, searchLetterIdx + wordDir

    if wordDir == -1 {
        startIdx, endIdx = endIdx, startIdx
    }

    dfs:
    for mm, nn, mn, k := m - 1, n - 1, m * n, 0; k < mn; k += 1 {
        i, j := k / n, k % n
        if board[i][j] != word[searchLetterIdx] {
            continue
        }

        if continueSearching := true; searchLetterIdx != endIdx {
            for _, dir := range directions {
                if ii, jj := i + dir[0], j + dir[1]; ii < 0 || ii > mm || jj < 0 || jj > nn {
                    continue
                } else if board[ii][jj] == word[nextLetterIdx] {
                    continueSearching = false
                    break
                }
            }

            if continueSearching {
                continue
            }
        }

        ilb, iub := max(0, i - maxSearchDist), min(mm, i + maxSearchDist)
        jlb, jub := j - maxSearchDist, j + maxSearchDist

        for ii := ilb; ii <= iub; ii += 1 {
            rowDist := abs(i - ii)
            jjlb, jjub := min(nn, max(0, jlb + rowDist)), min(nn, max(0, jub - rowDist))

            for jj := jjlb; jj <= jjub; jj += 1 {
                if board[ii][jj] != word[startIdx] {
                    continue
                }

                stack = append(stack, newDfsNode(ii, jj, 0, startIdx, true))
                for k := 0; len(stack) > 0; k = len(stack) - 1 {
                    curr := stack[k]
                    if curr.processed {
                        board[curr.row][curr.col], stack = word[curr.wordIdx], stack[:k]
                        continue
                    } else if curr.wordIdx == endIdx {
                        exists = true
                        break dfs
                    } else if curr.depth == maxSearchDist {
                        curr.bounded = false
                    }

                    stack[k].processed, board[curr.row][curr.col] = true, byte(0)
                    nextDepth, nextWordIdx := curr.depth + 1, curr.wordIdx + wordDir

                    for _, dir := range directions {
                        iii, jjj := curr.row + dir[0], curr.col + dir[1]
                        if curr.bounded && maxSearchDist - curr.depth < abs(i - iii) + abs(j - jjj) {
                            continue
                        } else if iii < 0 || iii > mm || jjj < 0 || jjj > nn {
                            continue
                        } else if board[iii][jjj] == word[nextWordIdx] {
                            stack = append(stack, newDfsNode(iii, jjj, nextDepth, nextWordIdx, curr.bounded))
                        }
                    }
                }
            }
        }
    }

    for _, dfsNode := range stack {
        if dfsNode.processed {
            board[dfsNode.row][dfsNode.col] = word[dfsNode.wordIdx]
        }
    }

    return exists
}