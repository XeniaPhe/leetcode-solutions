package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    grpSize, sortSize, temp := 1, 2, ListNode{0, head}
    for mergeCount, prev, curr := 0, &temp, temp.Next; ; {
        frst, scnd, frstCurr, scndCurr, scndPrev, merged := 0, 0, curr, curr, prev, prev
        for i := 0; i < grpSize && scndCurr != nil; i, scndPrev, scndCurr = i + 1, scndCurr, scndCurr.Next { }

        for frst != grpSize && scnd != grpSize && scndCurr != nil {
            if frstCurr.Val <= scndCurr.Val {
                frst, merged, frstCurr = frst + 1, frstCurr, frstCurr.Next
            } else {
                scnd, scndPrev.Next = scnd + 1, scndCurr.Next
                merged, merged.Next, scndCurr, scndCurr.Next = scndCurr, scndCurr, scndPrev.Next, merged.Next
            }
        }
        
        for i, rem := 0, sortSize - frst - scnd; i < rem && merged != nil; i, merged = i + 1, merged.Next { }
        if mergeCount += 1; merged == nil || merged.Next == nil {
            if mergeCount == 1 {
                break
            } else {
                mergeCount, prev, curr, grpSize, sortSize = 0, &temp, temp.Next, sortSize, sortSize * 2
            }
        } else {
            prev, curr = merged, merged.Next
        }
    }

    return temp.Next
}