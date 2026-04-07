/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func reverse(head *ListNode) *ListNode{
    if head ==nil{
         return nil
    }

    curr := head
    prev := &ListNode{}
    nexti := &ListNode{}

    for curr != nil {
        nexti  = curr.Next
        curr.Next = prev
        prev = curr
        curr = nexti
    }

    return prev
 }
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return nil
    }
    if left == right{
        return head
    }

    dummy := &ListNode{Val: 0 , Next: head}

    prev := dummy

    for i :=1; i<left; i++{
        prev = prev.Next
    }

    leftNode := prev.Next
    rightNode := leftNode

    for i := left; i<right; i++{
        rightNode = rightNode.Next
    }

    nexlist := rightNode.Next
    rightNode.Next = nil 
    startNode := reverse(leftNode)

    prev.Next = startNode
    leftNode.Next = nexlist

    return dummy.Next
}
