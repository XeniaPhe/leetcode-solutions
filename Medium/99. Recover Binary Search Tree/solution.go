package recover_binary_search_tree

import "math"

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func recoverTree(root *TreeNode) {
    var first, second, prev *TreeNode
    prevVal, nrLinks, removal, found := math.MinInt, 0, false, false

    morris:
    for curr, pre := root, root.Left; curr != nil; {
        if pre = curr.Left; pre == nil || removal {
            if prevVal > curr.Val && first == nil {
                first, second = prev, curr
            } else if prevVal > curr.Val && !found {
				second, found = curr, true
            }

            if found && nrLinks == 0 {
				break
            }

            prev, prevVal, curr, removal = curr, curr.Val, curr.Right, false
            continue
        }

        for ; pre.Right != nil; pre = pre.Right {
            if pre.Right == curr {
                nrLinks, removal, pre.Right = nrLinks - 1, true, nil
                continue morris
            }
        }

        nrLinks, pre.Right, curr = nrLinks + 1, curr, curr.Left
    }

	first.Val, second.Val = second.Val, first.Val
}