/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    runner := head
    cur := head
	for cur != nil {
		if runner == nil || runner.Next == nil {
    return false  // No cycle possible
}
             cur = cur.Next
		
		
		    runner = runner.Next.Next
		


		if cur == runner {
			return true
		}
	}
	return false
}
