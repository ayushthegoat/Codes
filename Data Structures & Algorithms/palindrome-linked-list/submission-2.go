/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
      fast, slow := head, head

	  for fast != nil && fast.Next != nil {
		slow  = slow.Next
		fast = fast.Next.Next
	  }

	  var prevNode *ListNode 
	  for slow != nil {
		tmp := slow.Next
		slow.Next = prevNode
		prevNode = slow
		slow = tmp
	  }

	  left, right := head, prevNode
	  for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	  }
	  

	  return true 
}
