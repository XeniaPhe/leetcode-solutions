package binary_tree_right_side_view

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func rightSideView(root *TreeNode) (view []int) {
    type stackNode struct {
        node *TreeNode
        level int
    }

    stack := []stackNode {stackNode{root, 0}}
    for curr := (stackNode{nil, 0}); len(stack) > 0; {
        if curr, stack = stack[len(stack) - 1], stack[:len(stack) - 1]; curr.node != nil {
            if len(view) == curr.level {
                view = append(view, curr.node.Val)
            }

            stack = append(stack, stackNode{curr.node.Left, curr.level + 1})
            stack = append(stack, stackNode{curr.node.Right, curr.level + 1})
        }
    }

    return view
}