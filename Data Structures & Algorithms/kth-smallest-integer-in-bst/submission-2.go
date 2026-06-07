/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func inorder(root *TreeNode , level *[]int){
    if root == nil {
        return
    }

    inorder(root.Left , level)
    *level = append(*level , root.Val)
    inorder(root.Right , level)
 }

func kthSmallest(root *TreeNode, k int) int {
    level := []int{}
    inorder(root , &level)

    return level[k-1]
}
