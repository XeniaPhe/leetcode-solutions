package copy_list_with_random_pointer

type Node struct {
	Val int
	Next, Random *Node
}

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    nodeMap := map[*Node]*Node{head: &Node{head.Val, nil, head.Random}}
    for prev, node := nodeMap[head], head.Next; node != nil; prev, node = prev.Next, node.Next {
        curr := &Node{node.Val, nil, node.Random}
        prev.Next, nodeMap[node] = curr, curr
    }

    for node := nodeMap[head]; node != nil; node = node.Next {
        node.Random = nodeMap[node.Random]
    }

    return nodeMap[head]
}