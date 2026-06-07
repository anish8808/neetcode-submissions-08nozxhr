/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

   queue := []*TreeNode{root}
   result := [][]int{}
   for len(queue)>0{
	size := len(queue)

	level := []int{}
	for i:=0; i<size; i++{
		frontNode := queue[0]
		queue = queue[1:]

		level = append(level , frontNode.Val)
		if frontNode.Left != nil {
			queue = append (queue , frontNode.Left)
		}
		if frontNode.Right != nil {
			queue = append(queue , frontNode.Right)
		}
	}

	result = append(result , level)
   }
	return result
}
