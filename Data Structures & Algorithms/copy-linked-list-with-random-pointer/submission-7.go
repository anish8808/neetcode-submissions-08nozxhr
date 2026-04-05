/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    //create copy
    if head == nil {
        return nil
    }
    curr := head

    for curr != nil{
        copyN := &Node{Val:curr.Val , Next:nil , Random:nil}
        copyN.Next = curr.Next
        curr.Next = copyN
        curr = copyN.Next
    }

    // adding random pointer

    curr = head

    for curr !=nil {
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        if curr.Next != nil{
            curr = curr.Next.Next
        }
    }

    // seperat the nodes 

    curr = head
    head2 := head.Next

    for curr !=nil {
        copy := curr.Next
        curr.Next = copy.Next
        curr = curr.Next

        if curr!= nil && copy.Next != nil {
            copy.Next = curr.Next
        }
    }

    return head2
}
