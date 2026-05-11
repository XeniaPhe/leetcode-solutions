package rotate_list

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }

    i, nh, nt, ot := 0, head, head, head
    for ; ot.Next != nil && i < k; i, ot = i + 1, ot.Next { }

    if i != k {
        for k, i, ot = k % (i + 1), 0, head; ot.Next != nil && i < k; i, ot = i + 1, ot.Next { }
    }

    for ; ot.Next != nil; ot, nt = ot.Next, nt.Next { }
    ot.Next = head
    nh, nt.Next = nt.Next, nil
    return nh
}