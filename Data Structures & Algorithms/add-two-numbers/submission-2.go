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
        // we assing sum to carry this ensures that going firther 
		//we automatically add teh sum when adding thsi deletes the addioton
		//if else block handling
	    sum := carry

		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		if cur3 != nil {
			sum += cur3.Val
			cur3 = cur3.Next
		}
        // / 10 gives you the 10s digit meanin the actual carry 
		carry = sum/10
		// % 10 gives me the remainder meaning which gives me the digit to add to the node

		digit := sum%10
		dummy.Next = &ListNode{Val:digit}
		dummy = dummy.Next


	}
	return cur.Next
}
