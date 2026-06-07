/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func checkBST(root *TreeNode , minVal , maxVal int)bool{
	if root == nil {
		return true
	}

	if root.Val <=minVal || root.Val >= maxVal{
		return false
	}

	return checkBST(root.Left , minVal , root.Val) && checkBST(root.Right , root.Val , maxVal)
}
func isValidBST(root *TreeNode) bool {
    return checkBST(root , math.MinInt64 , math.MaxInt64)
}
