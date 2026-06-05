package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    indices := make(map[int]int, len(inorder))
    for idx, val := range inorder {
        indices[val] = idx
    }

    return tree(preorder, inorder, indices, 0)
}

func tree(preorder []int, inorder []int, indices map[int]int, offset int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    } else if idx := indices[preorder[0]] + offset; idx < 0 {
        return &TreeNode{preorder[0], nil, nil}
    } else {
        return &TreeNode {
            Val: preorder[0],
            Left: tree(preorder[1:idx + 1], inorder[:idx], indices, offset),
            Right: tree(preorder[idx + 1:], inorder[idx + 1:], indices, offset - idx - 1),
        }
    }
}