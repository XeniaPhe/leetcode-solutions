package clone_graph

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    nodes, stack := make([]Node, 100), []*Node{node}
    for curr := node; len(stack) > 0; {
        stack, curr = stack[:len(stack) - 1], stack[len(stack) - 1]
        nodes[curr.Val - 1].Val = curr.Val
        nodes[curr.Val - 1].Neighbors = make([]*Node, len(curr.Neighbors))

        for i, nbor := range curr.Neighbors {
            nodes[curr.Val - 1].Neighbors[i] = &nodes[nbor.Val - 1]
            if nodes[nbor.Val - 1].Val == 0 {
                stack = append(stack, nbor)
            }
        }
    }

    return &nodes[node.Val - 1]
}