/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	if l == n {
		return head.Next
	}
	cur = head
	for i:=1; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}
