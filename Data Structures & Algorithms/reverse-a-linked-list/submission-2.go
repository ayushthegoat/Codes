/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prevNode *ListNode
	curNode := head

	for curNode != nil {
		temp := curNode.Next //moving temp forward one jump
		curNode.Next = prevNode // reversing the currentNodes -> NEXT LINK -> to previous node
		prevNode = curNode //moving prev forward one jump
		curNode = temp      //moving curNode forward one jump
	}
	return prevNode
}
