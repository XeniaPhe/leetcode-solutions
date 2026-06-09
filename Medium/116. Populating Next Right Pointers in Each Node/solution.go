package populating_next_right_pointers_in_each_node

type Node struct {
	Val int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
    for head, prev, curr := root, root, root; head != nil && head.Left != nil; head = head.Left {
        for prev, curr, head.Left.Next = head, head.Next, head.Right; curr != nil; prev, curr = curr, curr.Next {
            prev.Right.Next, curr.Left.Next = curr.Left, curr.Right
        }
    }

    return root
}