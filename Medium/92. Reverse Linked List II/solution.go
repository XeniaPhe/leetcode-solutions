package reverse_linked_list_II

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var anchorLeft, anchorRight, last, iter *ListNode
    if left == right {
        return head
    } else if left > 1 {
        iter = head
        for pos := 2; pos < left; iter, pos = iter.Next, pos + 1 { }
    }

    if anchorLeft = iter; iter == nil {
        anchorRight, last, iter = head, head, head.Next
    } else {
        anchorRight, last, iter = iter.Next, iter.Next, iter.Next.Next
    }

    for pos := left + 1; pos < right; pos += 1 {
        anchorRight, iter, iter.Next = iter, iter.Next, anchorRight
    }

    if iter == nil {
        return head
    } else if last.Next, iter.Next = iter.Next, anchorRight; anchorLeft != nil {
        anchorLeft.Next = iter
    } else {
        head = iter
    }

    return head
}