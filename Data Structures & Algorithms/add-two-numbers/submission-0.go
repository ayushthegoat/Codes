/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
	cur2 := l1
	cur3 := l2
    carry := 0
	dummy := &ListNode{}
	cur := dummy
	for cur2 != nil || cur3 != nil || carry > 0 {
        sum := carry

		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		if cur3 != nil {
			sum += cur3.Val
			cur3 = cur3.Next
		}

		carry = sum/10
		digit := sum%10
		dummy.Next = &ListNode{Val:digit}
		dummy = dummy.Next


	}
	return cur.Next
}
