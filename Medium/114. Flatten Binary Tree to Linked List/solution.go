package flatten_binary_tree_to_linked_list

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func flatten(root *TreeNode) {
    for curr, pre := root, root; curr != nil; {
        if curr.Left == nil {
            curr = curr.Right
        } else {
            for pre = curr.Left; pre.Right != nil; pre = pre.Right { }
            curr, curr.Right, pre.Right, curr.Left = curr.Left, curr.Left, curr.Right, nil
        }
    }
}