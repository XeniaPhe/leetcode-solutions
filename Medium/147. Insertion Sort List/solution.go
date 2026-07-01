package insertion_sort_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
    temp := ListNode{math.MinInt, head}
    for pi, pj, ci, cj := &temp, &temp, head, head; ci != nil; {
		if pi.Val > ci.Val {
			if cj.Val > ci.Val {
				pj, cj = &temp, temp.Next
			}
			
			for ; cj.Val < ci.Val; pj, cj = cj, cj.Next { }
			if ci != cj {
				ci, pi.Next, pj.Next, ci.Next = ci.Next, ci.Next, ci, cj
				continue
			}
		}

		pi, ci = ci, ci.Next
    }

    return temp.Next
}