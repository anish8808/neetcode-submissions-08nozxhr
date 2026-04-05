/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    maps := make(map[*Node]*Node)

    curr := head

    for curr != nil {
        node1 := &Node{Val:curr.Val , Next:nil , Random:nil}
        maps[curr] = node1
        curr = curr.Next 
    }

    curr = head

    for curr != nil {
        nodes , _ := maps[curr]
        nexts , _ := maps[curr.Next]
        randoms, _ := maps[curr.Random]
        nodes.Next =  nexts
        nodes.Random = randoms
        curr = curr.Next
    }   

    ans , _ := maps[head]
    return ans
}
