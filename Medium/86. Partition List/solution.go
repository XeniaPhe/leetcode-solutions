package partition_list

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    left, curr := (*ListNode)(nil), head
    for ; curr != nil && curr.Val < x; left, curr = curr, curr.Next { }

    for prev := left; curr != nil; prev, curr = curr, curr.Next {
        if curr.Val >= x {
            continue
        } else if prev == nil {
            left = head
            continue
        }

        prev.Next = curr.Next
        if left != nil {
            left.Next, curr.Next = curr, left.Next
        } else {
            head, curr.Next = curr, head
        }

        left, curr = curr, prev
    }

    return head
}