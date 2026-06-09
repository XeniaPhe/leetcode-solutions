package populating_next_right_pointers_in_each_node

type Node struct {
	Val int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
    for head, dummy := root, &(Node{}); head != nil; head, dummy.Next = dummy.Next, nil {
        for prev, curr := dummy, head; curr != nil; curr = curr.Next {
            if curr.Left != nil {
                prev, prev.Next = curr.Left, curr.Left
            }

            if curr.Right != nil {
                prev, prev.Next = curr.Right, curr.Right
            }
        }
    }

    return root
}