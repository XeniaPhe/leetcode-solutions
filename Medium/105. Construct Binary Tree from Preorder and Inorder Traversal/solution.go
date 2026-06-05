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

    return build(preorder, indices, 0)
}

func build(preorder []int, indices map[int]int, offset int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    } else if idx := indices[preorder[0]] + offset + 1; idx <= 0 {
        return &TreeNode{preorder[0], nil, nil}
    } else {
        return &TreeNode {
            Val: preorder[0],
            Left: build(preorder[1:idx], indices, offset),
            Right: build(preorder[idx:], indices, offset - idx),
        }
    }
}