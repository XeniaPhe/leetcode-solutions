package unique_binary_search_trees_II

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
    nums, bst, undo := make([]bool, n), make([]int, (1 << n) - 1), make([][]int, n)
    visited, leaves := make([]int, (1 << n) - 1), make([][3]int, 1, n + 1)
    leaves[0][2] = n
    
    for i := 0; i < n; i += 1 {
        undo[i] = make([]int, i + 1)
    } 

    return genTrees([]*TreeNode{}, nums, bst, visited, undo, leaves)
}

func genTrees(result []*TreeNode, nums []bool, bst []int, visited []int, undo [][]int, leaves [][3]int) []*TreeNode {
    if len(leaves) > len(nums) {
        stack, nodes := make([]int, 2, max(2, len(nums))), make([]TreeNode, len(nums))
        stack[0], stack[1] = 1, 2

        for i := 0; i < len(nums); i += 1 {
            nodes[i].Val = i + 1
        }

        for len(stack) > 0 {
            curr := stack[len(stack) - 1]
            if stack = stack[:len(stack) - 1]; curr >= len(bst) || bst[curr] == 0 {
                continue
            }

            stack = append(stack, (2 * curr) + 1, (2 * curr) + 2)
            if curr & 1 == 1 {
                nodes[bst[(curr - 1) / 2] - 1].Left = &nodes[bst[curr] - 1]
            } else {
                nodes[bst[(curr - 1) / 2] - 1].Right = &nodes[bst[curr] - 1]
            }
        }

        return append(result, &nodes[bst[0] - 1])
    }

    j := len(leaves)
    leaves = append(leaves, [3]int{0,0,0})
    for i := 0; i < j; i += 1 {
        idx, lb, ub := leaves[i][0], leaves[i][1], leaves[i][2]
        leaves[i][0], leaves[j][0] = (2 * idx) + 1, (2 * idx) + 2
        undo[j - 1][i] = visited[idx]

        for k := lb; k < ub; k += 1 {
            if !nums[k] && visited[idx] & (1 << k) == 0 {
                visited[idx] |= (1 << k)
                leaves[i][1], leaves[i][2] = lb, k
                leaves[j][1], leaves[j][2] = k + 1, ub
                bst[idx], nums[k] = k + 1, true
                result = genTrees(result, nums, bst, visited, undo, leaves)
                nums[k] = false
            }
        }

        leaves[i], bst[idx] = [3]int {idx, lb, ub}, 0
    }

    leaves = leaves[:len(leaves) - 1]
    for i := 0; i < j; i += 1 {
        visited[leaves[i][0]] = undo[j - 1][i]
    }

    return result
}