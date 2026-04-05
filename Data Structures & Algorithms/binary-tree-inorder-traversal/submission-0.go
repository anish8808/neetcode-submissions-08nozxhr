/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func solve(root *TreeNode , ans *[]int){
	if root == nil {
		return
	}

	solve(root.Left , ans)
	*ans = append(*ans ,root.Val)
	solve(root.Right , ans)
}
func inorderTraversal(root *TreeNode) []int {
	
	ans := []int{}

	solve(root , &ans)

	return ans

}
