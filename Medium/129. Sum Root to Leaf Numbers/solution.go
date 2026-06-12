package sum_root_to_leaf_numbers

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
    type DFSNode struct {
        node *TreeNode
        num int
    }

    result, stack := 0, []DFSNode{DFSNode{root, 0}}
    for len(stack) > 0 {
        size := len(stack) - 1
        curr := stack[size]
        stack = stack[:size]
        num := 10 * curr.num + curr.node.Val

        if curr.node.Left != nil {
            stack = append(stack, DFSNode{curr.node.Left, num})
        }
        
        if curr.node.Right != nil {
            stack = append(stack, DFSNode{curr.node.Right, num})
        }

        if len(stack) == size {
            result += num
        }
    }

    return result
}