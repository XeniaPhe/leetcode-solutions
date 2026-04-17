package remove_nth_node_from_end_of_list

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prev, node *ListNode = nil, head
    for i := 1; i <= n; node, i = node.Next, i + 1 { }

    if node == nil {
        return head.Next
    }

    for prev = head; node.Next != nil; prev, node = prev.Next, node.Next { }
    prev.Next = prev.Next.Next
    return head
}