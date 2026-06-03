package binary_tree_level_order_traversal

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    queue, traversal := make([]*TreeNode, 1, 4), make([][]int, 1, 4)
    queue[0], traversal[0] = root, make([]int, 1)

    for depth, curr := 0, root; len(queue) > 0; depth += 1 {
        for count, i := len(queue), 0; i < count; i += 1 {
            if queue, curr = queue[1:], queue[0]; curr.Left != nil {
                queue = append(queue, curr.Left)
            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }

            traversal[depth][i] = curr.Val
        }

        traversal = append(traversal, make([]int, len(queue)))
    }

    return traversal[:len(traversal) - 1]
}