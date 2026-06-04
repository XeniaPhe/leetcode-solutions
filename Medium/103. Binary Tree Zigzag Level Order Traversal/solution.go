package binary_tree_zigzag_level_order_traversal

import "slices"

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    queue, traversal := make([]*TreeNode, 1, 4), make([][]int, 1, 4)
    queue[0], traversal[0] = root, make([]int, 1)

    for depth, count, asc := 0, 1, true; count > 0; depth, count, asc = depth + 1, len(queue), !asc {
        for i := 0; i < count; i += 1 {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }

            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }

            traversal[depth][i] = queue[i].Val
        }

        queue, traversal = queue[count:], append(traversal, make([]int, len(queue[count:])))
        if !asc {
            slices.Reverse(traversal[depth])
        }
    }

    return traversal[:len(traversal) - 1]
}