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
    solve(root.Right, ans)
    *ans = append(*ans , root.Val)
}

func postorderTraversal(root *TreeNode) []int {
    ans := []int{}

    solve(root , &ans)

    return ans
}
