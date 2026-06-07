/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func max(a , b int) int {
	if a>b{
		return a
	}

	return b
 }

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftHeight , rightDepth)
}
