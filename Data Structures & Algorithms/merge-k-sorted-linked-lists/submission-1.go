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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) ==0 {
		return nil
	}

	if len(lists) ==1 {
		return lists[0]
	}

	list1 := lists[0]
	finalList := &ListNode{}
	for i:=1; i<len(lists); i++{
		finalList = mergeTwoLists(list1 , lists[i])
		list1 = finalList
	}

	return finalList
}
