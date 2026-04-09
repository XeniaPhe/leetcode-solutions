package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    resultNodes := make([]ListNode, 102)
    currIdx := 0
    head := &resultNodes[0]
    current := head
    carry := 0

    for ; l1 != nil || l2 != nil; {
        numLeft, numRight := 0, 0

        if l1 != nil {
            numLeft = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            numRight = l2.Val
            l2 = l2.Next
        }

        sum := numLeft + numRight + carry
        carry = 0

        if sum > 9 {
            sum -= 10
            carry = 1
        }

        currIdx += 1
        current.Val = sum
        current.Next = &resultNodes[currIdx]
        current = current.Next
    }

    if carry != 0 {
        currIdx += 1
        current.Val = carry
    }

    resultNodes[currIdx - 1].Next = nil
    return head
}