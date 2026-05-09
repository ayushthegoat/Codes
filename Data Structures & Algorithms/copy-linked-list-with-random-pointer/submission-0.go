/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
     mapper := make(map[*Node]*Node)
	 cur := head
	 for cur != nil {
		mapper[cur] = &Node{Val:cur.Val}
		cur = cur.Next
	 }

	 cur = head
	 for cur != nil {
		mapper[cur].Next = mapper[cur.Next]
		mapper[cur].Random = mapper[cur.Random]
		cur = cur.Next
	 }
	 return mapper[head]
}
