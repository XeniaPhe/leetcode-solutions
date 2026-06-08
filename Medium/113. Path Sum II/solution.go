package path_sum_II

import "slices"

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }

    return pathSumBacktrack(root, targetSum, &[]int{}, [][]int{})
}

func pathSumBacktrack(root *TreeNode, targetSum int, path *[]int, result [][]int) [][]int {
    if *path, targetSum = append(*path, root.Val), targetSum - root.Val; root.Left != nil {
        result = pathSumBacktrack(root.Left, targetSum, path, result)
    }

	if root.Right != nil {
		result = pathSumBacktrack(root.Right, targetSum, path, result)
	} else if root.Left == nil && targetSum == 0 {
		result = append(result, slices.Clone(*path))
	}

	*path = (*path)[:len(*path) - 1]
	return result
}