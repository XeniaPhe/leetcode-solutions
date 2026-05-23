package remove_duplicates_from_sorted_list_II

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    for val, cnt, prev, curr, iter := 0, 1, (*ListNode)(nil), head, head; iter != nil; cnt = 1 {
        for val, iter = curr.Val, iter.Next; iter != nil && iter.Val == val; iter, cnt = iter.Next, cnt + 1 { }
        if cnt > 1 && prev != nil {
            prev.Next, curr = iter, iter
        } else if cnt > 1 {
            head, curr = iter, iter
        } else {
            prev, curr = curr, iter
        }
    }

    return head
}