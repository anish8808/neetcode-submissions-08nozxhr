/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func solve(root *TreeNode , inorder *[]int){
    if root==nil{
        return
    }

    solve(root.Left , inorder)
    *inorder = append(*inorder , root.Val)
    solve(root.Right , inorder)
 }

func kthSmallest(root *TreeNode, k int) int {
    var inorder []int
    solve(root , &inorder)

    return inorder[k-1]
}
