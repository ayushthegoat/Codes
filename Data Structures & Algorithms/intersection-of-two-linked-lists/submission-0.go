/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    first := headA
	second := headB

	 for first != second {
		if first == nil {
			first = headB
		} else {
			first = first.Next
		}
		if second == nil {
			second = headA
		} else {
			second = second.Next
		}
	 }
	 return first
}
