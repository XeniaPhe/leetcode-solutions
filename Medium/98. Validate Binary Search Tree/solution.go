package validate_binary_search_tree

import "math"

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

type treeDfsNode struct {
    node *TreeNode
    lb, ub int
}

func isValidBST(root *TreeNode) bool {
    stack := make([]treeDfsNode, 0, 1)
    stack = append(stack, treeDfsNode{root, math.MinInt, math.MaxInt})

    for len(stack) > 0 {
        curr := stack[len(stack) - 1]
        if stack = stack[:len(stack) - 1]; curr.node.Val < curr.lb || curr.node.Val > curr.ub {
            return false
        }

        if curr.node.Left != nil {
            stack = append(stack, treeDfsNode{curr.node.Left, curr.lb, curr.node.Val - 1})
        }

        if curr.node.Right != nil {
            stack = append(stack, treeDfsNode{curr.node.Right, curr.node.Val + 1, curr.ub})
        }
    }

    return true
}