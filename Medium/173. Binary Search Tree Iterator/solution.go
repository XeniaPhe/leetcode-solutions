package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
    lineage []*TreeNode
    current *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    return *(&BSTIterator{nil, nil}).extendLineageFrom(root)
}

func (this *BSTIterator) Next() (next int) {
    if next = this.current.Val; this.current.Right != nil {
        this.extendLineageFrom(this.current.Right)
    } else if last := len(this.lineage) - 1; last >= 0 {
        this.current, this.lineage = this.lineage[last], this.lineage[:last]
    } else {
        this.current = nil
    }

    return next
}

func (this *BSTIterator) HasNext() bool {
    return this.current != nil || len(this.lineage) > 0
}

func (this *BSTIterator) extendLineageFrom(node *TreeNode) *BSTIterator {
    for ; node.Left != nil; this.lineage, node = append(this.lineage, node), node.Left { }
    this.current = node
    return this
}