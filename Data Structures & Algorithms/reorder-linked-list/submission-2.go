/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reorderList(head *ListNode) {
//     slow := head
//     fast := head
//     curl := head
//     for fast != nil && fast.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//     secondHalf := slow.Next
//     slow.Next = nil                     
//     revl := reverser(secondHalf)
     
//     for curl != nil && revl != nil {
//         cn := curl.Next
//         rn := revl.Next 

//         curl.Next = revl
//         if cn != nil { 
//           revl.Next = cn 
//         }

//         curl = cn
//         revl = rn
//     }
    

// }

// func reverser(head *ListNode) *ListNode {
//     node := head
//     var prevNode *ListNode

//     for node != nil {
//        temp := node.Next
//        node.Next = prevNode
//        prevNode = node
//        node = temp
//     }
//     return prevNode
// }
//optimized method hai ye this one 
func reorderList(head *ListNode) {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    second := slow.Next
    slow.Next = nil
    var prev *ListNode
    for second != nil {
        temp := second.Next
        second.Next = prev
        prev = second
        second = temp
    }

    first := head
    second = prev

    //most crucial loop in place merge of a linked list 
    //logic is we MUST MUST have 2 pointers one for fist one and then for second one 
    // then what we must do is simple
    // store the next poitners of teh current Node
    //first.Next goes to the first node of teh second list and then 
    //the second.Next goes to the second pointer of the first List 
    //meaning simple  ki i am iterating and changing more like a simulation
    // then just move the pointers next 
    for second != nil {
        tmp1, tmp2 := first.Next, second.Next
        first.Next = second
        second.Next = tmp1
        first, second = tmp1, tmp2
    }
}
