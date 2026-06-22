package linked_list_cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    for walker, stroller, runner := head, head, head; ; {
        if runner.Next == nil || runner.Next.Next == nil {
            return nil
        } else if walker, runner = walker.Next, runner.Next.Next; walker == runner {
            for ; walker != stroller; walker, stroller = walker.Next, stroller.Next { }
            return walker
        }
    }

    return nil
}