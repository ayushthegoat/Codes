/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    
	dummy := &ListNode{Next: head}
    
    // Step 1: Find node BEFORE left position
    prevNode := dummy
    cur := head
    for i := 1; i < left; i++ {
        prevNode = cur
        cur = cur.Next
    }
	//step 2:find node next to left 
	rightNode := cur
	for i := left; i < right; i++ {
       rightNode = rightNode.Next
	}
	nextNode := rightNode.Next
	prevNode.Next = reverseList(cur, right-left+1)
	cur.Next = nextNode
	return dummy.Next
}

func reverseList(head *ListNode, n int) *ListNode {
    var prev *ListNode
    cur := head
    
    for i := 0; i < n; i++ {
        temp := cur.Next
        cur.Next = prev
        prev = cur
        cur = temp
    }
    
    return prev
}