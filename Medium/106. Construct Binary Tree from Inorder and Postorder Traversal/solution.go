package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    indices := make(map[int]int, len(inorder))
    for idx, val := range inorder {
        indices[val] = idx
    }

    return build(postorder, indices, 0)
}

func build(postorder []int, indices map[int]int, offset int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    
    lastIdx, rootVal := len(postorder) - 1, postorder[len(postorder) - 1]
    if idx := indices[rootVal] + offset; idx < 0 {
        return &TreeNode{rootVal, nil, nil}
    } else {
        return &TreeNode {
            Val: rootVal,
            Left: build(postorder[:idx], indices, offset),
            Right: build(postorder[idx:lastIdx], indices, offset - idx - 1),
        }
    }
}