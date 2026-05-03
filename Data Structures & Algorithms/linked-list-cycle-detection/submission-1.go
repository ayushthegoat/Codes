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
	for runner != nil && runner.Next != nil {
        cur = cur.Next
	    runner = runner.Next.Next
		if cur == runner {
			return true
		}
	}
	return false
}
