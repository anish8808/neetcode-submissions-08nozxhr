/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyNode := &ListNode{Val : 0 , Next: nil}
	
	if list1 == nil && list2 == nil {
		return nil 
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	ansNode := dummyNode

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val{
			dummyNode.Next = list2
			dummyNode = dummyNode.Next
			list2 = list2.Next
		} else {
			dummyNode.Next = list1
			dummyNode = dummyNode.Next
			list1 = list1.Next
		}
	}

	if list1!=nil {
		dummyNode.Next = list1
	}
	if list2 != nil {
		dummyNode.Next = list2
	}

	ansNode = ansNode.Next
	return ansNode

}
