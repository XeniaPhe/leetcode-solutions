package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
    prev, nodes := &ListNode{0, nil}, []*ListNode{}
    for node := head; node != nil; node, nodes = node.Next, append(nodes, node) { }

    for ; len(nodes) > 1; nodes = nodes[1:len(nodes) - 1] {
        curr, next := nodes[0], nodes[len(nodes) - 1]
        prev, prev.Next, curr.Next, next.Next = next, curr, next, nil
    }

    if len(nodes) == 1 {
        prev.Next, nodes[0].Next = nodes[0], nil
    }
}