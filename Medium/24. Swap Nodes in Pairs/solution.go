package swap_nodes_in_pairs

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := head.Next

    for prevPrev, prev, curr := (&ListNode{0, nil}), head, head.Next; curr != nil; {
        prevPrev.Next, prev.Next = curr, curr.Next
        curr.Next = prev

        if prev.Next == nil {
            break
        }

        prevPrev, prev = prev, prev.Next
        curr = prev.Next
    }

    return newHead
}