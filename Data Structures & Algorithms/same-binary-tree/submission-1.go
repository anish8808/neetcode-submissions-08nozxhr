/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q ==nil {
		return true
	}

	if p != nil && q == nil {
		return false 
	}
	if p == nil && q != nil {
		return false 
	}

	if p.Val != q.Val {
		return false
	}
	leftans := false
	rightans := false
	if p.Val == q.Val {
		leftans = isSameTree(p.Left , q.Left)
		rightans = isSameTree(p.Right , q.Right)
	}

	return leftans && rightans

}
