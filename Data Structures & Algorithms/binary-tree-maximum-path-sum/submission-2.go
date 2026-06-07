/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func max(a  , b int) int {
	if a>b {
		return a
	}

	return b
 }
func solve(root *TreeNode , maxSum *int)int{
	if root == nil {
		return 0
	}

	left := max(0 , solve(root.Left , maxSum))
	right := max(0 , solve(root.Right , maxSum))
	sum:= root.Val + left + right
	leftSum := root.Val + left
	rightSum := root.Val + right
	*maxSum = max(*maxSum , max(sum , max(leftSum, rightSum))) 

	return  max(leftSum , rightSum)
}

func maxPathSum(root *TreeNode) int {
	maxSum:= math.MinInt64
    solve(root , &maxSum)
	return maxSum
}
