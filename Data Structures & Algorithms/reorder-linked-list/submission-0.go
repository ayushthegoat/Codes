/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow := head
    fast := head
    curl := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    secondHalf := slow.Next
    slow.Next = nil                     
    revl := reverser(secondHalf)
     
    for curl != nil && revl != nil {
        cn := curl.Next
        rn := revl.Next 

        curl.Next = revl
        if cn != nil { 
          revl.Next = cn 
        }

        curl = cn
        revl = rn
    }
    

}

func reverser(head *ListNode) *ListNode {
    node := head
    var prevNode *ListNode

    for node != nil {
       temp := node.Next
       node.Next = prevNode
       prevNode = node
       node = temp
    }
    return prevNode
}
