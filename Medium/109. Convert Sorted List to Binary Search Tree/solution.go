package convert_sorted_list_to_binary_search_tree

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    var nodes []TreeNode
    for node := head; node != nil; node = node.Next {
        nodes = append(nodes, TreeNode{node.Val, nil, nil})
    }

    return buildBST(nodes)
}   

func buildBST(nodes []TreeNode) *TreeNode {
    if len(nodes) == 0 {
        return nil
    }

    mid := len(nodes) / 2
    if mid > 0 {
        nodes[mid].Left = buildBST(nodes[:mid])
        nodes[mid].Right = buildBST(nodes[mid + 1:])
    }
    
    return &nodes[mid]
}