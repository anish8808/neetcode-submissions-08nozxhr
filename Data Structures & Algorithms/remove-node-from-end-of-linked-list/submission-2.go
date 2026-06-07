/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyNode := &ListNode{}
	dummyNode.Next = head
	startNode := dummyNode
	for n>0{
		startNode = startNode.Next
		n--
	}
	prev := dummyNode
	for startNode.Next != nil && startNode != nil {
		prev = prev.Next
		startNode = startNode.Next
	}

	if prev.Next.Next == nil {
		prev.Next = nil 
	}else {
		prev.Next = prev.Next.Next
	}

	ans := dummyNode.Next

	return ans
}
